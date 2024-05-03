import { Quiz, StudentQuiz } from "../interfaces/quiz";
import { api } from "../utils/api";
import { StudentCompetency } from "../interfaces/competency";
import { Meta } from "../interfaces/meta";

const path = "/teacher/quiz";

export interface GetQuizResponse {
  quiz: Quiz;
}

export interface StudentQuizReviewResponseData {
  quiz: StudentQuiz;
  competency: StudentCompetency[];
}

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
export interface GetQuizHistoryFilter {
  id: string;
  page: number;
  limit: number;
}

export const getAllQuiz = async (filter: GetAllQuizFilter) => {
  return (await api.get(path, { params: filter })).data as GetAllQuiz;
};
export const getQuizHistory = async (filter: GetQuizHistoryFilter) => {
  return (await api.get(`${path}/history/${filter.id}`, { params: filter }))
    .data as GetAllQuiz;
};

export const getQuiz = async (id: string) => {
  console.log((await api.get(`${path}/${id}`)).data.data);
  return (await api.get(`${path}/${id}`)).data.data as GetQuizResponse;
};
