package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"icando/internal/domain/repository"
	"icando/internal/model"
	"icando/internal/model/dto"
)

type StudentQuizService interface {
	CalculateScore(id uuid.UUID) error
}

type StudentQuizImpl struct {
	studentQuizRepository repository.StudentQuizRepository
	db                    *gorm.DB
}

func (s *StudentQuizImpl) CalculateScore(id uuid.UUID) error {
	studentQuiz, err := s.studentQuizRepository.GetStudentQuiz(dto.GetStudentQuizFilter{
		WithQuizQuestions: true,
		WithAnswers:       true,
		ID:                id,
	})

	if err != nil {
		return err
	}

	// todo check status

	answerMap := make(map[string]model.StudentAnswer)

	for _, answer := range studentQuiz.StudentAnswers {
		answerMap[answer.QuestionID.String()] = answer
	}

	//questionCount := len(studentQuiz.Quiz.Questions)
	correctCount := 0

	for _, question := range studentQuiz.Quiz.Questions {
		answer, ok := answerMap[question.ID.String()]

		if ok {
			isCorrect := answer.AnswerID == question.AnswerID
			answer.IsCorrect = &isCorrect

			answerCompetencies := make([]model.StudentAnswerCompetency, 0)

			for _, competency := range question.Competencies {
				answerCompetencies = append(answerCompetencies, model.StudentAnswerCompetency{
					CompetencyID: competency.ID,
					IsPassed:     isCorrect,
				})
			}

			if err := answer.SetCompetencies(answerCompetencies); err != nil {
				return err
			}

			answerMap[question.ID.String()] = answer
			correctCount++
		} // not ok result mean that the question is not answered
	}

	// todo set correct count
	// todo update score

	return nil
}
