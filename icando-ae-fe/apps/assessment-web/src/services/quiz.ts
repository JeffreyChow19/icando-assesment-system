import { QuizAttempt, StudentQuiz } from "../interfaces/quiz";
import { api } from "../utils/api";

const path = "/student/quiz";

export const getQuizAvailability = async () => {
  return (await api.get(`${path}/overview`)).data.data as StudentQuiz;
};

export const startQuiz = async () => { 
  return (await api.patch(`${path}/start`)).data.data as QuizAttempt;
}