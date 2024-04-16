import { api } from "../utils/api.ts";
import { Meta } from "../interfaces/meta.ts";
import { Competency } from "../interfaces/competency.ts";
import { Question } from "../interfaces/question.ts";
import { QuizDetail } from "src/interfaces/quiz.ts";

const path = "/designer/quiz";

export interface GetAllCompetencyFilter {
  page?: number;
  limit?: number;
  query?: string;
}

export interface GetAllCompetencyResponse {
  meta: Meta;
  competencies: Competency[];
}

export interface CreateQuestionPayload {
  text: string;
  choices: {
    id: number;
    text: string;
  }[];
  answerId: number;
  order: number;
  competencies: string[];
}

export interface CreateQuestionResponse {
  data: Question;
}

export interface UpdateQuestionResponse {
  data: Question;
}

export interface CreateQuizResponseData {
  id: string;
  name: string | null;
  subject: string[] | null;
  passingGrade: number;
  publishedAt: string | null;
  startAt: string | null;
  endAt: string | null;
  duration: number | null;
}

export interface PublishQuizResponseData extends CreateQuizResponseData {}

export interface UpdateQuizPayload {
  id: string;
  name: string;
  subject: string[];
  passingGrade: number;
  startAt?: string | null;
  endAt?: string | null;
}

export interface PublishQuizPayload {
  quizId: string;
  quizDuration: number;
  startDate: string;
  endDate: string;
  assignedClasses: string[];
}

export interface GetAllQuizFilter {
  name?: string;
  subject?: string;
  page: number;
  limit: number;
}
interface GetAllQuiz {
  meta: Meta;
  data: QuizDetail[];
}

export type UpdateQuestionPayload = CreateQuestionPayload;

export const createQuestion = async (
  quizId: string,
  payload: CreateQuestionPayload,
) => {
  return (await api.post(
    `${path}/${quizId}/question`,
    payload,
  )) as CreateQuestionResponse;
};

export const updateQuestion = async (
  quizId: string,
  questionId: string,
  payload: UpdateQuestionPayload,
) => {
  return (await api.patch(`${path}/${quizId}/question/${questionId}`, payload))
    .data as UpdateQuestionResponse;
};

export const deleteQuestion = async (quizId: string, questionId: string) => {
  return await api.delete(`${path}/${quizId}/question/${questionId}`);
};

export const createQuiz = async () => {
  return (await api.post(path)).data.data as CreateQuizResponseData;
};

export const updateQuiz = async (payload: UpdateQuizPayload) => {
  await api.patch(path, payload);
};

export const publishQuiz = async (payload: PublishQuizPayload) => {
  return (await api.post(path, payload)).data.data as PublishQuizResponseData;
};

export const getQuiz = async (id: string) => {
  return (await api.get(`${path}/${id}`)).data.data as QuizDetail;
};

export const getAllQuiz = async (filter: GetAllQuizFilter) => {
  return (await api.get(path, { params: filter })).data as GetAllQuiz;
};
