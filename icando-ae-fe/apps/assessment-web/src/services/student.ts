import { Student } from "../interfaces/user";
import { api } from "../utils/api";

const path = "/auth/student";

export const getStudentProfile = async () => {
  return (await api.get(`${path}/profile`)).data.data as Student;
};

