import { setToken } from "../utils/local-storage.ts";
import { User } from "../interfaces/user.ts";
import { api } from "../utils/api.ts";

interface CheckAuthResponse {
  data: User;
  token?: string;
}

const path = "/auth/student";

export const checkAuth = async (): Promise<User> => {
  const response = (await api.get(`${path}/profile`)).data as CheckAuthResponse;

  if (response.token) {
    setToken(response.token);
  }

  return response.data;
};

export const saveQuizToken = async (token: string) => {
  setToken(token);
};
