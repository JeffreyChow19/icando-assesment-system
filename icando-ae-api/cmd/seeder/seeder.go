package main

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"icando/internal/model"
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

	institutions := []model.Institution{
		{ID: uuid.New(), Name: "SMAN 1 Kwangya", Nis: "23456781", Slug: "sman-1-kwangya"},
		{ID: uuid.New(), Name: "SMAN 2 Kwangya", Nis: "23456782", Slug: "sman-2-kwangya"},
		{ID: uuid.New(), Name: "SMAN 3 Kwangya", Nis: "23456783", Slug: "sman-3-kwangya"},
		{ID: uuid.New(), Name: "SMAN 4 Kwangya", Nis: "23456784", Slug: "sman-4-kwangya"},
		{ID: uuid.New(), Name: "SMAN 5 Kwangya", Nis: "23456785", Slug: "sman-5-kwangya"},
	}

	for _, institution := range institutions {
		if err := tx.Create(&institution).Error; err != nil {
			tx.Rollback()
			logger.Log.Error(err.Error())
			return
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)

		for i := 0; i < 3; i++ {
			firstName, lastName, email := generateAccount()

			ld := model.Teacher{
				ID:            uuid.New(),
				FirstName:     firstName,
				LastName:      lastName,
				Email:         email,
				Password:      string(hashedPassword),
				InstitutionID: institution.ID,
				Role:          enum.TEACHER_ROLE_LEARNING_DESIGNER,
			}

			if err := tx.Create(&ld).Error; err != nil {
				tx.Rollback()
				logger.Log.Error(err.Error())
				return
			}
		}

		for i := 0; i < 5; i++ {
			firstName, lastName, email := generateAccount()

			teacher := model.Teacher{
				ID:            uuid.New(),
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

			grades := []string{"Grade 1", "Grade 2", "Grade 3"}

			for _, grade := range grades {
				for j := 1; j <= 3; j++ {
					class := model.Class{
						ID:            uuid.New(),
						Name:          fmt.Sprintf("Class %d", j),
						Grade:         grade,
						InstitutionID: institution.ID,
						TeacherID:     teacher.ID,
					}

					if err := tx.Create(&class).Error; err != nil {
						tx.Rollback()
						logger.Log.Error(err.Error())
						return
					}

					for k := 0; k < 8; k++ {
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
							ClassID:       class.ID,
						}

						if err := tx.Create(&student).Error; err != nil {
							tx.Rollback()
							logger.Log.Error(err.Error())
							return
						}
					}
				}
			}
		}
	}

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

	if err := tx.Commit().Error; err != nil {
		logger.Log.Error(err.Error())
		return
	}
}

func generateAccount() (firstName string, lastName string, email string) {
	firstName = faker.Name().FirstName()
	lastName = faker.Name().FirstName()
	email = fmt.Sprintf("%s%s@email.com", strings.ToLower(firstName), strings.ToLower(lastName))

	return firstName, lastName, email
}
