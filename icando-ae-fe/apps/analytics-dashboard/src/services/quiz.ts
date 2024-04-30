import { StudentQuiz } from "src/interfaces/quiz";
import { api } from "../utils/api";
import { StudentCompetency } from "src/interfaces/competency";

const path = "/teacher/quiz";

interface StudentQuizReviewResponseData {
  quiz: StudentQuiz;
  competency: StudentCompetency[];
}

export const getStudentQuizReview = async (
  quizId: string,
  studentQuizId: string,
) => {
  return (await api.get(`${path}/${quizId}/students/${studentQuizId}`)).data
    .data as StudentQuizReviewResponseData;
};
