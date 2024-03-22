import { Class } from "../interfaces/classes.ts";
import { api } from "../utils/api.ts";

export interface GetClassFilter {
  withTeacher: boolean;
  withStudents: boolean;
}

export interface CreateClassPayload {
  name: string;
  grade: string;
  teacherIds?: string[];
}

export interface UpdateClassPayload {
  name?: string;
  grade?: string;
  teacherIds?: string[];
}

export interface AssignStudentsRequest {
  studentIds: string[];
}

export interface GetAllClassesResponse {
  data: Class[];
}

export interface GetClassResponse {
  data: Class;
}

const path = "/designer/class";

export const createClass = async (payload: CreateClassPayload) => {
  await api.post(path, payload);
};

export const getAllClasses = async () => {
  return (await api.get(path)).data as GetAllClassesResponse;
};

export const getClass = async (param: GetClassFilter, id: string) => {
  return (await api.get(`${path}/${id}`, { params: param }))
    .data as GetClassResponse;
};

export const updateClass = async (payload: UpdateClassPayload, id: string) => {
  await api.patch(`${path}/${id}`, payload);
};

export const deleteClass = async (id: string) => {
  await api.delete(`${path}/${id}`);
};

export const assignStudents = async (
  id: string,
  payload: AssignStudentsRequest,
) => {
  await api.post(`${path}/${id}/students`, payload);
};

export const unAssignStudents = async (
  id: string,
  payload: AssignStudentsRequest,
) => {
  await api.patch(`${path}/${id}/students`, payload);
};
