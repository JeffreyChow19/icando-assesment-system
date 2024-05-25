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
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"runtime"
	"strings"
	"syreclabs.com/go/faker"
	"time"
)

// Structs to match the JSON data
type QuestionData struct {
	Text    string `json:"text"`
	Choices []struct {
		ID   int    `json:"id"`
		Text string `json:"text"`
	} `json:"choices"`
	AnswerID int    `json:"answer_id"`
	Subject  string `json:"subject"`
}

type CompetencyData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type InstitutionData struct {
	Name string `json:"name"`
	Nis  string `json:"nis"`
	Slug string `json:"slug"`
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Ensure randomness

	// Get the directory of the current file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		logger.Log.Error("Failed to get the current file directory")
		return
	}
	dir := filepath.Dir(filename)

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

	// Load questions from JSON file once
	var questionsData []QuestionData
	questionFile, err := ioutil.ReadFile(filepath.Join(dir, "data/data_question.json"))
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	if err := json.Unmarshal(questionFile, &questionsData); err != nil {
		logger.Log.Error(err.Error())
		return
	}

	// Load competencies from JSON file
	var competenciesData []CompetencyData
	competencyFile, err := ioutil.ReadFile(filepath.Join(dir, "data/data_competency.json"))
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	if err := json.Unmarshal(competencyFile, &competenciesData); err != nil {
		logger.Log.Error(err.Error())
		return
	}

	// Create competencies
	for _, comp := range competenciesData {
		competency := model.Competency{
			Model: model.Model{
				ID: uuid.New(),
			},
			Numbering:   comp.ID,
			Name:        comp.Name,
			Description: comp.Description,
		}

		if err := tx.Create(&competency).Error; err != nil {
			tx.Rollback()
			logger.Log.Error(err.Error())
			return
		}
	}

	// Load institutions from JSON file
	var institutionsData []InstitutionData
	institutionFile, err := ioutil.ReadFile(filepath.Join(dir, "data/data_institution.json"))
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	if err := json.Unmarshal(institutionFile, &institutionsData); err != nil {
		logger.Log.Error(err.Error())
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)

	for _, inst := range institutionsData {
		institution := model.Institution{
			Name: inst.Name,
			Nis:  inst.Nis,
			Slug: inst.Slug,
		}

		if err := tx.Create(&institution).Error; err != nil {
			tx.Rollback()
			logger.Log.Error(err.Error())
			return
		}

		// Create 5 learning designers for the institution
		var learningDesigners []model.Teacher
		for i := 0; i < 5; i++ {
			firstName, lastName, email := generateAccount()

			learningDesigner := model.Teacher{
				Model: model.Model{
					ID: uuid.New(),
				},
				FirstName:     firstName,
				LastName:      lastName,
				Email:         email,
				Password:      string(hashedPassword),
				InstitutionID: institution.ID,
				Role:          enum.TEACHER_ROLE_LEARNING_DESIGNER,
			}

			if err := tx.Create(&learningDesigner).Error; err != nil {
				tx.Rollback()
				logger.Log.Error(err.Error())
				return
			}

			learningDesigners = append(learningDesigners, learningDesigner)
		}

		// Create 50 regular teachers for the institution
		var teachers []model.Teacher
		for i := 0; i < 50; i++ {
			firstName, lastName, email := generateAccount()

			teacher := model.Teacher{
				Model: model.Model{
					ID: uuid.New(),
				},
				FirstName:     firstName,
				LastName:      lastName,
				Email:         email,
				Password:      string(hashedPassword),
				InstitutionID: institution.ID,
				Role:          enum.TEACHER_ROLE_REGULAR,
			}

			if err := tx.Create(&teacher).Error; err != nil {
				tx.Rollback()
				logger.Log.Error(err.Error())
				return
			}

			teachers = append(teachers, teacher)
		}

		grades := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
		classes := []string{"A", "B", "C", "D"}

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

				// Associate two teachers with class
				rand.Shuffle(len(teachers), func(i, j int) {
					teachers[i], teachers[j] = teachers[j], teachers[i]
				})

				selectedTeachers := teachers[:2]

				for _, teacher := range selectedTeachers {
					class.Teachers = append(class.Teachers, teacher)
				}

				if err := tx.Save(&class).Error; err != nil {
					tx.Rollback()
					logger.Log.Error(err.Error())
					return
				}

				// Create 7 students for each class
				for i := 0; i < 7; i++ {
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

				// Shuffle the questions data once
				rand.Shuffle(len(questionsData), func(i, j int) {
					questionsData[i], questionsData[j] = questionsData[j], questionsData[i]
				})

				// Create 2 quizzes for each class
				for i := 0; i < 2; i++ {
					quizName := fmt.Sprintf("Quiz %d", i+1)

					// Randomize passing grade between 50 and 70
					passingGrade := float64(rand.Intn(21) + 50)

					// Select 7 random questions from the shuffled questions data
					selectedQuestions := make([]QuestionData, 7)
					copy(selectedQuestions, questionsData)

					// Shuffle the selected questions again for randomness
					rand.Shuffle(len(selectedQuestions), func(i, j int) {
						selectedQuestions[i], selectedQuestions[j] = selectedQuestions[j], selectedQuestions[i]
					})

					// Collect distinct subjects from the selected questions
					subjectSet := make(map[string]struct{})
					for _, q := range selectedQuestions {
						subjectSet[q.Subject] = struct{}{}
					}

					var subjects []string
					for subject := range subjectSet {
						subjects = append(subjects, subject)
					}

					quiz := model.Quiz{
						Model: model.Model{
							ID: uuid.New(),
						},
						Name:         &quizName,
						Subject:      base.StringArray(subjects),
						PassingGrade: passingGrade,
						CreatedBy:    learningDesigners[i].ID,
						UpdatedBy:    &learningDesigners[i].ID,
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

					// Create questions for each quiz
					for j, q := range selectedQuestions {
						var choices []model.QuestionChoice
						for _, choice := range q.Choices {
							choices = append(choices, model.QuestionChoice{
								ID:   choice.ID,
								Text: choice.Text,
							})
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
							Text:     q.Text,
							Choices:  &postgres.Jsonb{RawMessage: jsonChoices},
							AnswerID: q.AnswerID,
							QuizID:   quiz.ID,
							Order:    j + 1,
						}

						// Assign 2 random competencies to each question
						var compList []model.Competency
						tx.Order("random()").Limit(2).Find(&compList)
						question.Competencies = compList

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
