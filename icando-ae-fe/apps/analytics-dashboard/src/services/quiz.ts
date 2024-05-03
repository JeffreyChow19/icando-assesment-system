import { Quiz } from "../interfaces/quiz.ts";
import { api } from "../utils/api.ts";

const path = "/teacher/quiz";

export interface GetQuizResponse {
  quiz: Quiz;
}
export const getQuiz = async (id: string) => {
  console.log((await api.get(`${path}/${id}`)).data.data);
  return (await api.get(`${path}/${id}`)).data.data as GetQuizResponse;
};
