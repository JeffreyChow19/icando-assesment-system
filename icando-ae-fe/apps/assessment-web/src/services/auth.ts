import { setToken } from "../utils/local-storage.ts";

export const saveQuizToken = async (token: string) => {
  setToken(token);
};
