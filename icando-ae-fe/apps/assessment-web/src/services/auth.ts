import { setToken } from "../utils/local-storage.ts";
import { api } from "../utils/api.ts";
import { StudentQuiz } from "src/interfaces/quiz.ts";

interface CheckQuizAvailabilityResponse {
  data: StudentQuiz;
}

export const checkQuizAvailability = async (): Promise<StudentQuiz> => {
  const response = (await api.get(`student/quiz`)).data as CheckQuizAvailabilityResponse;

  return response.data;
};

export const saveQuizToken = async (token: string) => {
  setToken(token);
};
