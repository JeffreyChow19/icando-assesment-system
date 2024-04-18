import { setToken } from "../utils/local-storage.ts";
import { User } from "../interfaces/user.ts";
import { api } from "../utils/api.ts";

export interface LoginPayload {
  email: string;
  password: string;
}

interface AuthResponse {
  token: string;
}

interface CheckAuthResponse {
  data: User;
  token?: string;
}

const path = "/auth/designer";

export const login = async (payload: LoginPayload): Promise<void> => {
  const { token } = (await api.post(`${path}/login`, payload))
    .data as AuthResponse;
  setToken(token);
};

export const checkAuth = async (): Promise<User> => {
  const response = (await api.get(`${path}/profile`)).data as CheckAuthResponse;

  if (response.token) {
    setToken(response.token);
  }

  return response.data;
};
