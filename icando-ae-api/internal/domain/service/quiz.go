package service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"icando/internal/domain/repository"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/internal/model/enum"
	"icando/lib"
	"icando/utils/httperror"
	"net/http"
	"time"
)

type QuizService interface {
	CreateQuiz(id uuid.UUID) (*dao.QuizDao, *httperror.HttpError)
	GetQuiz(id uuid.UUID) (*dao.QuizDao, *httperror.HttpError)
	UpdateQuiz(userID uuid.UUID, quizDto dto.UpdateQuizDto) (*dao.QuizDao, *httperror.HttpError)
	GetAllQuizzes(filter dto.GetAllQuizzesFilter) ([]dao.ParentQuizDao, *dao.MetaDao, *httperror.HttpError)
	PublishQuiz(quizDto dto.PublishQuizDto) (*dao.QuizDao, *httperror.HttpError)
}

type QuizServiceImpl struct {
	quizRepository repository.QuizRepository
	authService    AuthService
	db             *gorm.DB
}

func NewQuizServiceImpl(quizRepository repository.QuizRepository, db *lib.Database, authService AuthService) *QuizServiceImpl {
	return &QuizServiceImpl{
		quizRepository: quizRepository,
		db:             db.DB,
		authService:    authService,
	}
}

var ErrCreateQuiz = &httperror.HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected error happened when creating quiz"),
}

func (s *QuizServiceImpl) CreateQuiz(id uuid.UUID) (*dao.QuizDao, *httperror.HttpError) {
	quiz := model.Quiz{
		CreatedBy: id,
	}
	quiz, err := s.quizRepository.CreateQuiz(quiz)

	if err != nil {
		return nil, ErrCreateQuiz
	}

	quizDao := quiz.ToDao()

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

	quizDao := quiz.ToDao()
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
	if quizDto.Deadline != nil {
		quiz.Deadline = quizDto.Deadline
	}

	err = s.quizRepository.UpdateQuiz(*quiz)
	if err != nil {
		return nil, ErrUpdateQuiz
	}

	quizDao := quiz.ToDao()

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

func (s *QuizServiceImpl) PublishQuiz(quizDto dto.PublishQuizDto) (*dao.QuizDao, *httperror.HttpError) {
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

	if err := tx.Where("class_id IN ?", classIds).Error; err != nil {
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

	if err := tx.Create(&studentQuizzes).Error; err != nil {
		tx.Rollback()
		return nil, &httperror.HttpError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	// todo for each student quiz, enqueue email request
	for _, studentQuiz := range studentQuizzes {
		// token here
		_, err := s.authService.GenerateQuizToken(dto.GenerateQuizTokenDto{
			StudentQuizId: studentQuiz.ID,
			ExpiredAt:     time.Now().Add(time.Hour * 12), // placeholder. todo change when expiredAt column already implemented
		})

		// todo access variable quiz and students to get quiz and students information

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

	quizDao := quiz.ToDao()

	return &quizDao, nil
}
