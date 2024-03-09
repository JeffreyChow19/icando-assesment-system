import { setToken } from "../utils/local-storage.ts";
import { User } from "../interfaces/user.ts";
import { api } from "../utils/api.ts";

export interface ChangePasswordPayload {
  oldPassword: string;
  newPassword: string;
}

export interface LoginPayload {
  email: string;
  password: string;
}

export interface RegisterPayload {
  name: string;
  password: string;
  email: string;
}

export interface UpdateUserRolePayload {
  userId: number;
  role: "USER" | "ADMIN";
}

interface AuthResponse {
  token: string;
}

interface CheckAuthResponse {
  data: User;
  token?: string;
}

const path = "/designer/auth";

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

export const changePassword = async (payload: ChangePasswordPayload) => {
  await api.put(`${path}/password`, payload);
};

export const register = async (payload: RegisterPayload): Promise<void> => {
  await api.post(`${path}/register`, payload);
};

export const getAllUsers = async () => {
  return (await api.get(`${path}`)).data.data as User[];
};

export const updateUserRole = async (payload: UpdateUserRolePayload) => {
  await api.put(`${path}/role`, payload);
};

export const deleteUser = async (id: number) => {
  await api.delete(`${path}/${id}`);
};
