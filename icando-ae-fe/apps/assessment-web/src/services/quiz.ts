import { QuizDetail } from "../interfaces/quiz";
import { api } from "../utils/api";

const path = "/student/quiz";

export const getQuizAvailability = async () => {
  return (await api.get(path)).data.data as QuizDetail;
};
