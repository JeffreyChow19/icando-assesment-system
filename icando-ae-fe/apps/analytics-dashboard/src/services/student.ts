import { api } from "../utils/api.ts";
import { Meta } from "../interfaces/meta.ts";
import { Student } from "../interfaces/student.ts";
import { Class } from '../interfaces/class.ts';

export interface GetAllStudentsFilter {
  classId?: string;
  name?: string;
  page: number;
  limit: number;
  orderBy?: string;
  asc?: boolean;
}

interface GetAllStudentsResponse {
  meta: Meta;
  data: Student[];
}

const path = "/teacher/analytics";

export const getAllStudent = async (filter: GetAllStudentsFilter) => {
  return (await api.get(`${path}/student`, { params: filter }))
    .data as GetAllStudentsResponse;
};

export interface GetAllClassesResponse {
  data: Class[];
}

export const getAllClasses = async () => {
  const { data } = (await api.get(`${path}/class`)).data as GetAllClassesResponse;
  return data;
};





