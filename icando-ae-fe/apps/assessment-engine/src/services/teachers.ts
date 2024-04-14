import { api } from "../utils/api.ts";
import { Teacher } from "../interfaces/teacher.ts";

export interface GetAllTeachersResponse {
  data: Teacher[];
}

const path = "/designer/teacher";

export const getAllTeachers = async () => {
  return (await api.get(path)).data as GetAllTeachersResponse;
};
