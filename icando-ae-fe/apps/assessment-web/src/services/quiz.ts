import { StudentQuiz } from "../interfaces/quiz";
import { api } from "../utils/api";

const path = "/student/quiz";

export const getQuizAvailability = async () => {
  return (await api.get(`${path}/overview`)).data.data as StudentQuiz;
};

export const getQuizDetail = async () => {
  return (await api.get(`${path}/detail`)).data.data as StudentQuiz;
};

export const updateAnswer = async (questionId: string, choiceId: number) => {
  await api.post(`${path}/question/${questionId}`, { answer_id: choiceId });
};

export const startQuiz = async () => {
  return (await api.patch(`${path}/start`)).data.data as StudentQuiz;
};

export const submitQuiz = async () => {
  await api.patch(`${path}/submit`);
};
