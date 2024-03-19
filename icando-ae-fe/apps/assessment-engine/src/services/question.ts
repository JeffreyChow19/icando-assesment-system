import { api } from "../utils/api.ts";
import { Meta } from "../interfaces/meta.ts";
import { Competency } from "../interfaces/competency.ts";
import { Question } from "../interfaces/question.ts";

const path = "/designer/competency";

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
  correctAnswer: number;
  competencies: string[];
}

export interface CreateQuestionResponse {
  data: Question;
}

export interface UpdateQuestionResponse {
  data: Question;
}

export type UpdateQuestionPayload = CreateQuestionPayload;

export const getAllCompetency = async (filter?: GetAllCompetencyFilter) => {
  return (await api.get(path, { params: filter }))
    .data as GetAllCompetencyResponse;
};

export const createQuestion = async (payload: CreateQuestionPayload) => {
  return (await api.post(path, payload)) as CreateQuestionResponse;
};
export const updateQuestion = async (payload: UpdateQuestionPayload) => {
  return (await api.put(path, payload)) as UpdateQuestionResponse;
};
