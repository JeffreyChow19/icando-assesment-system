import { StudentQuiz } from "src/interfaces/quiz";
import { api } from "../utils/api";
import { StudentCompetency } from "src/interfaces/competency";
import { Meta } from "../interfaces/meta.ts";

const path = "/teacher/student-quiz";

export interface StudentQuizReviewResponseData {
  quiz: StudentQuiz;
  competency: StudentCompetency[];
}

export interface GetStudentQuizzesResponse {
  meta: Meta;
  data: {
    id: string;
    name: string;
    className: string;
    classGrade: string;
    completedAt?: string;
    totalScore: number;
    correctCount: string;
    status: string;
  }[];
}

export interface GetStudentQuizzesFilter {
  page: number;
  limit: number;
  classId?: string;
  studentName?: string;
  quizStatus?: string;
  quizId?: string;
}

export const getStudentQuizReview = async (studentQuizId: string) => {
  return (await api.get(`${path}/${studentQuizId}`)).data
    .data as StudentQuizReviewResponseData;
};

export const getStudentQuizzes = async (params: GetStudentQuizzesFilter) => {
  return (await api.get(path, { params })).data as GetStudentQuizzesResponse;
};
