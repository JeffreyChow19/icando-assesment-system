package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
	"icando/internal/model"
	"icando/internal/model/base"
	"icando/internal/model/enum"
	"icando/lib"
	"icando/utils/logger"
	"math/rand"
	"strings"
	"syreclabs.com/go/faker"
)

func main() {
	config, err := lib.NewConfig()
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}

	db, err := lib.NewDatabase(config)
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}

	tx := db.DB.Begin()

	// Create 50 competencies
	for i := 1; i <= 50; i++ {
		competency := model.Competency{
			Model: model.Model{
				ID: uuid.New(),
			},
			Numbering:   fmt.Sprintf("C%02d", i),
			Name:        fmt.Sprintf("Competency %d", i),
			Description: fmt.Sprintf("Description for competency %d", i),
		}

		if err := tx.Create(&competency).Error; err != nil {
			tx.Rollback()
			logger.Log.Error(err.Error())
			return
		}
	}

	// Create 5 institutions
	institutions := []model.Institution{
		{Name: "SMAN 1 Kwangya", Nis: "23456781", Slug: "sman-1-kwangya"},
		{Name: "SMAN 2 Kwangya", Nis: "23456782", Slug: "sman-2-kwangya"},
		{Name: "SMAN 3 Kwangya", Nis: "23456783", Slug: "sman-3-kwangya"},
		{Name: "SMAN 4 Kwangya", Nis: "23456784", Slug: "sman-4-kwangya"},
		{Name: "SMAN 5 Kwangya", Nis: "23456785", Slug: "sman-5-kwangya"},
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)

	for _, institution := range institutions {
		if err := tx.Create(&institution).Error; err != nil {
			tx.Rollback()
			logger.Log.Error(err.Error())
			return
		}

		grades := []string{"Grade 1", "Grade 2", "Grade 3"}
		classes := []string{"Class A", "Class B", "Class C"}

		for _, grade := range grades {
			for _, className := range classes {
				class := model.Class{
					Model: model.Model{
						ID: uuid.New(),
					},
					Name:          className,
					Grade:         grade,
					InstitutionID: institution.ID,
				}

				if err := tx.Create(&class).Error; err != nil {
					tx.Rollback()
					logger.Log.Error(err.Error())
					return
				}

				// Create 2 teachers for each class
				for i := 0; i < 2; i++ {
					firstName, lastName, email := generateAccount()

					var role enum.TeacherRole
					if i == 0 {
						role = enum.TEACHER_ROLE_LEARNING_DESIGNER
					} else {
						role = enum.TEACHER_ROLE_REGULAR
					}

					teacher := model.Teacher{
						Model: model.Model{
							ID: uuid.New(),
						},
						FirstName:     firstName,
						LastName:      lastName,
						Email:         email,
						Password:      string(hashedPassword),
						InstitutionID: institution.ID,
						Role:          role,
					}

					if err := tx.Create(&teacher).Error; err != nil {
						tx.Rollback()
						logger.Log.Error(err.Error())
						return
					}

					// Associate teacher with class
					class.Teachers = append(class.Teachers, teacher)
					if err := tx.Save(&class).Error; err != nil {
						tx.Rollback()
						logger.Log.Error(err.Error())
						return
					}
				}

				// Create 8 students for each class
				for i := 0; i < 8; i++ {
					firstName, lastName, email := generateAccount()
					nisn := fmt.Sprintf("%010d", rand.Intn(10000000000))

					student := model.Student{
						Model: model.Model{
							ID: uuid.New(),
						},
						FirstName:     firstName,
						LastName:      lastName,
						Nisn:          nisn,
						Email:         email,
						InstitutionID: institution.ID,
						ClassID:       &class.ID,
					}

					if err := tx.Create(&student).Error; err != nil {
						tx.Rollback()
						logger.Log.Error(err.Error())
						return
					}
				}

				// Create 2 quizzes for each class
				for i := 0; i < 2; i++ {
					quizName := fmt.Sprintf("Quiz %d", i+1)
					quiz := model.Quiz{
						Model: model.Model{
							ID: uuid.New(),
						},
						Name:         &quizName,
						Subject:      base.StringArray{"MATEMATIKA", "ILMU PENGETAHUAN ALAM (IPA)"},
						PassingGrade: 70.0,
						CreatedBy:    class.Teachers[0].ID,
						UpdatedBy:    &class.Teachers[0].ID,
					}

					if err := tx.Create(&quiz).Error; err != nil {
						tx.Rollback()
						logger.Log.Error(err.Error())
						return
					}

					// Associate quiz with class
					class.Quizzes = append(class.Quizzes, quiz)
					if err := tx.Save(&class).Error; err != nil {
						tx.Rollback()
						logger.Log.Error(err.Error())
						return
					}

					// Create 5 questions for each quiz
					for j := 0; j < 5; j++ {
						choices := []model.QuestionChoice{
							{ID: 0, Text: "Choice A"},
							{ID: 1, Text: "Choice B"},
							{ID: 2, Text: "Choice C"},
							{ID: 3, Text: "Choice D"},
						}

						jsonChoices, err := json.Marshal(choices)
						if err != nil {
							logger.Log.Error(err.Error())
							tx.Rollback()
							return
						}

						question := model.Question{
							Model: model.Model{
								ID: uuid.New(),
							},
							Text:     fmt.Sprintf("Question %d", j+1),
							Choices:  &postgres.Jsonb{RawMessage: jsonChoices},
							AnswerID: rand.Intn(4), // AnswerID is now between 0 and 3
							QuizID:   quiz.ID,
							Order:    j + 1,
						}

						// Assign 2 random competencies to each question
						var competencies []model.Competency
						tx.Order("random()").Limit(2).Find(&competencies)
						question.Competencies = competencies

						if err := tx.Create(&question).Error; err != nil {
							tx.Rollback()
							logger.Log.Error(err.Error())
							return
						}
					}
				}
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		logger.Log.Error(err.Error())
		return
	}
}

func generateAccount() (firstName string, lastName string, email string) {
	firstName = faker.Name().FirstName()
	lastName = faker.Name().LastName()
	email = fmt.Sprintf("%s.%s@email.com", strings.ToLower(firstName), strings.ToLower(lastName))

	return firstName, lastName, email
}
