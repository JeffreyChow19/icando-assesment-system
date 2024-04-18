import { StudentQuiz } from "../interfaces/quiz";
import { api } from "../utils/api";

const path = "/student/quiz";

export const getQuizAvailability = async () => {
  return (await api.get(`${path}/overview`)).data.data as StudentQuiz;
};
