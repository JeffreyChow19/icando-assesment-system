import { Quiz, StudentQuiz } from "../interfaces/quiz";
import { api } from "../utils/api";
import { StudentCompetency } from "../interfaces/competency";
import { Meta } from "../interfaces/meta";

const path = "/teacher/quiz";

export interface StudentQuizReviewResponseData {
  quiz: StudentQuiz;
  competency: StudentCompetency[];
}

// todo: change endpoint for getallquiz
export interface GetAllQuizFilter {
  name?: string;
  subject?: string;
  page: number;
  limit: number;
}

interface GetAllQuiz {
  meta: Meta;
  data: Quiz[];
}

export const getAllQuiz = async (filter: GetAllQuizFilter) => {
  return (await api.get("/designer/quiz", { params: filter }))
    .data as GetAllQuiz;
};
// todo: endpoint for quiz history

export const getStudentQuizReview = async (
  quizId: string,
  studentQuizId: string,
) => {
  return (await api.get(`${path}/${quizId}/students/${studentQuizId}`)).data
    .data as StudentQuizReviewResponseData;
};
