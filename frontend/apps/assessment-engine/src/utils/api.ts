import axios from "axios";
import { VITE_API_URL } from "./config.ts";
import { getToken } from "./local-storage.ts";

export const api = axios.create({
  baseURL: VITE_API_URL,
});

api.defaults.headers.common["Content-Type"] = "application/json";

api.interceptors.request.use((config) => {
  const token = getToken();
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});
