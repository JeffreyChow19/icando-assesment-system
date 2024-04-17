package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"icando/internal/domain/repository"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/internal/worker/client"
	"icando/internal/worker/task"
	"icando/lib"
	"icando/utils/httperror"
	"net/http"
)

type QuizService interface {
	CreateQuiz(id uuid.UUID) (*dao.QuizDao, *httperror.HttpError)
	GetQuiz(id uuid.UUID) (*dao.QuizDao, *httperror.HttpError)
	UpdateQuiz(userID uuid.UUID, quizDto dto.UpdateQuizDto) (*dao.QuizDao, *httperror.HttpError)
	GetAllQuizzes(filter dto.GetAllQuizzesFilter) ([]dao.ParentQuizDao, *dao.MetaDao, *httperror.HttpError)
	PublishQuiz(teacherID uuid.UUID, quizDto dto.PublishQuizDto) (*dao.QuizDao, *httperror.HttpError)
}

type QuizServiceImpl struct {
	quizRepository    repository.QuizRepository
	teacherRepository repository.TeacherRepository
	authService       AuthService
	db                *gorm.DB
	config            *lib.Config
	workerClient      *client.WorkerClient
}

func NewQuizServiceImpl(
	quizRepository repository.QuizRepository,
	teacherRepository repository.TeacherRepository,
	db *lib.Database,
	config *lib.Config,
	workerClient *client.WorkerClient,
	authService AuthService) *QuizServiceImpl {
	return &QuizServiceImpl{
		quizRepository:    quizRepository,
		db:                db.DB,
		authService:       authService,
		config:            config,
		teacherRepository: teacherRepository,
		workerClient:      workerClient,
	}
}

var ErrCreateQuiz = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when creating quiz"),
}

func (s *QuizServiceImpl) CreateQuiz(id uuid.UUID) (*dao.QuizDao, *httperror.HttpError) {
	quiz := model.Quiz{
		CreatedBy: id,
		UpdatedBy: &id,
	}
	quiz, err := s.quizRepository.CreateQuiz(quiz)

	if err != nil {
		return nil, ErrCreateQuiz
	}

	quizDao := quiz.ToDao(true)

	return &quizDao, nil
}

var ErrQuizNotFound = &httperror.HttpError{
	StatusCode: http.StatusNotFound,
	Err:        errors.New("Quiz Not Found"),
}

var ErrGetQuiz = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when retrieving quiz"),
}

var ErrUpdateQuiz = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when updating quiz"),
}

func (s *QuizServiceImpl) GetQuiz(id uuid.UUID) (*dao.QuizDao, *httperror.HttpError) {
	quiz, err := s.quizRepository.GetQuiz(dto.GetQuizFilter{ID: id, WithCreator: true, WithUpdater: true, WithQuestions: true})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrQuizNotFound
		}
		return nil, ErrGetQuiz
	}

	quizDao := quiz.ToDao(true)
	return &quizDao, nil
}

func (s *QuizServiceImpl) UpdateQuiz(userID uuid.UUID, quizDto dto.UpdateQuizDto) (*dao.QuizDao, *httperror.HttpError) {
	quiz, err := s.quizRepository.GetQuiz(dto.GetQuizFilter{ID: quizDto.ID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrQuizNotFound
		}
		return nil, ErrUpdateQuiz
	}

	quiz.UpdatedBy = &userID

	if quizDto.Name != nil {
		quiz.Name = quizDto.Name
	}
	if quizDto.Subject != nil {
		quiz.Subject = quizDto.Subject
	}
	if quizDto.PassingGrade != 0 {
		quiz.PassingGrade = quizDto.PassingGrade
	}

	err = s.quizRepository.UpdateQuiz(*quiz)
	if err != nil {
		return nil, ErrUpdateQuiz
	}

	quizDao := quiz.ToDao(true)

	return &quizDao, nil
}

func (s *QuizServiceImpl) GetAllQuizzes(filter dto.GetAllQuizzesFilter) ([]dao.ParentQuizDao, *dao.MetaDao, *httperror.HttpError) {
	quizzes, meta, err := s.quizRepository.GetAllQuiz(filter)
	if err != nil {
		log.Print(err)
		return nil, nil, httperror.InternalServerError
	}

	return quizzes, meta, nil
}

func (s *QuizServiceImpl) PublishQuiz(teacherID uuid.UUID, quizDto dto.PublishQuizDto) (*dao.QuizDao, *httperror.HttpError) {
	teacher, err := s.teacherRepository.GetTeacher(dto.GetTeacherFilter{
		ID: &teacherID,
	})

	if err != nil {
		return nil, &httperror.HttpError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	tx := s.db.Begin()

	quiz, err := s.quizRepository.CloneQuiz(tx, quizDto)

	if err != nil {
		tx.Rollback()
		return nil, &httperror.HttpError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	var students []model.Student
	classIds := make([]string, 0)

	for _, classId := range quizDto.AssignedClasses {
		classIds = append(classIds, classId.String())
	}

	if err := tx.Where("class_id IN ?", classIds).Find(&students).Error; err != nil {
		tx.Rollback()
		return nil, &httperror.HttpError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	studentQuizzes := make([]model.StudentQuiz, 0)

	for _, student := range students {
		studentQuizzes = append(studentQuizzes, model.StudentQuiz{
			Status:    enum.NOT_STARTED,
			QuizID:    quiz.ID,
			StudentID: student.ID,
		})
	}

	if len(studentQuizzes) == 0 {
		tx.Rollback()
		return nil, &httperror.HttpError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Cannot publish because no student to assign to"),
		}
	}

	if err := tx.Create(&studentQuizzes).Error; err != nil {
		tx.Rollback()
		return nil, &httperror.HttpError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	for i, studentQuiz := range studentQuizzes {
		// token here
		token, err := s.authService.GenerateQuizToken(dto.GenerateQuizTokenDto{
			StudentQuizId: studentQuiz.ID,
			ExpiredAt:     *quiz.EndAt,
		})

		if err != nil {
			tx.Rollback()
			return nil, &httperror.HttpError{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}

		student := students[i]

		payload := task.SendQuizEmailPayload{
			QuizName:     *quiz.Name,
			QuizSubjects: quiz.Subject,
			QuizDuration: *quiz.Duration,
			QuizEndAt:    *quiz.EndAt,
			QuizStartAt:  *quiz.StartAt,
			QuizUrl:      s.BuildUrl(token.Token),
			TeacherName:  fmt.Sprintf("%s %s", teacher.FirstName, teacher.LastName),
			TeacherEmail: teacher.Email,
			StudentName:  fmt.Sprintf("%s %s", student.FirstName, student.LastName),
			StudentEmail: student.Email,
		}

		emailTask, err := task.NewSendQuizEmailTask(payload)

		if err != nil {
			tx.Rollback()
			return nil, &httperror.HttpError{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}

		_, err = s.workerClient.Enqueue(emailTask)

		if err != nil {
			tx.Rollback()
			return nil, &httperror.HttpError{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}

		calculateTask, err := task.NewCalcualteStudentQuizTask(task.CalculateStudentQuizPayload{
			StudentQuizID: studentQuiz.ID,
		})

		if err != nil {
			tx.Rollback()
			return nil, &httperror.HttpError{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}

		_, err = s.workerClient.Enqueue(calculateTask, asynq.ProcessAt(*quiz.EndAt))

		if err != nil {
			tx.Rollback()
			return nil, &httperror.HttpError{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, &httperror.HttpError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	quizDao := quiz.ToDao(true)

	return &quizDao, nil
}

func (s *QuizServiceImpl) BuildUrl(token string) string {
	return fmt.Sprintf("%s/quiz?token=%s", s.config.AssessmentWebHost, token)
}
