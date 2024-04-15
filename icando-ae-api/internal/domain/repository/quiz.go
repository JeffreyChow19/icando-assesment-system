package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"icando/internal/model"
	"icando/internal/model/dao"
	"icando/internal/model/dto"
	"icando/lib"
	"math"
	"sort"
	"strings"
	"time"
)

type QuizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *lib.Database) QuizRepository {
	return QuizRepository{
		db: db.DB,
	}
}

func (r *QuizRepository) GetQuiz(filter dto.GetQuizFilter) (*model.Quiz, error) {
	query := r.db.Session(&gorm.Session{})

	if filter.WithCreator {
		query = query.Preload("Creator")
	}

	if filter.WithUpdater {
		query = query.Preload("Updater")
	}

	if filter.WithClasses {
		query = query.Preload("Classes")
	}

	if filter.WithQuestions {
		query = query.Preload("Questions.Competencies")
	}

	if filter.ID != uuid.Nil {
		query = query.Where("id = ?", filter.ID)
	}

	var quiz model.Quiz
	err := query.First(&quiz).Error
	if err != nil {
		return nil, err
	}

	if filter.WithQuestions {
		sort.Slice(quiz.Questions, func(i, j int) bool {
			return quiz.Questions[i].Order < quiz.Questions[j].Order
		})
	}

	return &quiz, nil
}

func (r *QuizRepository) CreateQuiz(quiz model.Quiz) (model.Quiz, error) {
	err := r.db.Create(&quiz).Error
	return quiz, err
}

func (r *QuizRepository) UpdateQuiz(quiz model.Quiz) error {
	return r.db.Save(&quiz).Error
}

func (r *QuizRepository) GetAllQuiz(filter dto.GetAllQuizzesFilter) ([]dao.ParentQuizDao, *dao.MetaDao, error) {
	query := r.db.Table("quizzes").Select(
		`quizzes.id, quizzes.name, quizzes.subject, quizzes.passing_grade, MAX(c.published_at) as last_published_at, t1.first_name || ' ' || t1.last_name as created_by, t2.first_name || ' ' || t2.last_name as updated_by`).
		Joins("INNER JOIN teachers t1 ON quizzes.created_by=t1.id").
		Joins("INNER JOIN teachers t2 ON quizzes.updated_by=t2.id").
		Joins("LEFT JOIN quizzes c ON quizzes.id=c.parent_quiz").
		Where("quizzes.parent_quiz IS NULL")

	if filter.Query != nil {
		query.Where("LOWER(name) LIKE ?", strings.ToLower(fmt.Sprintf("%%%s%%", *filter.Query)))
	}
	if filter.Subject != nil {
		query.Where("LOWER(subject) LIKE ?", strings.ToLower(fmt.Sprintf("%%%s%%", *filter.Subject)))
	}

	query.Group("quizzes.id, t1.first_name, t1.last_name, t2.first_name, t2.last_name")

	var totalItem int64
	err := query.Session(&gorm.Session{}).Count(&totalItem).Error
	if err != nil {
		return nil, nil, err
	}

	meta := dao.MetaDao{
		Page:      filter.Page,
		Limit:     filter.Limit,
		TotalItem: totalItem,
		TotalPage: int(math.Ceil(float64(totalItem) / float64(filter.Limit))),
	}
	Paginate(query, filter.Page, filter.Limit)

	quizzes := []dao.ParentQuizDao{}
	err = query.Session(&gorm.Session{}).Scan(&quizzes).Error

	return quizzes, &meta, err
}

func (r *QuizRepository) CloneQuiz(db *gorm.DB, quizDto dto.PublishQuizDto) (*model.Quiz, error) {
	var oldQuiz model.Quiz

	if err := db.Preload("Questions.Competencies").Where("id = ?", quizDto.QuizID.String()).First(&oldQuiz).Error; err != nil {
		return nil, err
	}

	// get all classes in ids
	var classes []model.Class

	classIds := make([]string, 0)

	for _, classID := range quizDto.AssignedClasses {
		classIds = append(classIds, classID.String())
	}

	if err := db.Where("id in ?", classIds).Find(&classes).Error; err != nil {
		return nil, err
	}

	if len(classIds) != len(classes) {
		return nil, errors.New("Some assigned class not found")
	}

	now := time.Now()

	newQuiz := model.Quiz{
		Name:         oldQuiz.Name,
		Subject:      oldQuiz.Subject,
		PassingGrade: oldQuiz.PassingGrade,
		ParentQuiz:   &oldQuiz.ID,
		CreatedBy:    oldQuiz.CreatedBy,
		UpdatedBy:    oldQuiz.UpdatedBy,
		PublishedAt:  &now,
		Questions:    make([]model.Question, 0),
		Classes:      classes,
		// todo assign startdate, duration, and enddate here. wait for other commiter
	}

	for _, question := range oldQuiz.Questions {
		newQuiz.Questions = append(newQuiz.Questions, model.Question{
			Text:         question.Text,
			AnswerID:     question.AnswerID,
			Competencies: question.Competencies,
			Order:        question.Order,
			Choices:      &postgres.Jsonb{RawMessage: question.Choices.RawMessage},
		})
	}

	if err := db.Create(&newQuiz).Error; err != nil {
		return nil, err
	}

	return &newQuiz, nil
}
