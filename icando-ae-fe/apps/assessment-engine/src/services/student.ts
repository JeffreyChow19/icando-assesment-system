import { api } from '../utils/api.ts';
import { Meta } from '../interfaces/meta.ts';
import { Student } from '../interfaces/student.ts';

export interface CreateStudentPayload {
  firstName: string;
  lastName: string;
  nisn: string;
  email: string;
}

export interface UpdateStudentPayload {
  firstName?: string;
  lastName?: string;
  classId?: string;
}

export interface GetAllStudentsFilter {
  classId?: string;
  name?: string;
  page: number;
  limit: number;
}

interface GetAllStudentsResponse {
  meta: Meta;
  data: Student[];
}

interface GetStudentResponse {
  data: Student;
}


const path = '/designer/student';

export const createStudent = async (payload: CreateStudentPayload) => {
  await api.post(path, payload);
};

export const updateStudent = async (payload: UpdateStudentPayload, id: string) => {
  await api.patch(`${path}/${id}`, payload);
};

export const getAllStudent = async (filter: GetAllStudentsFilter) => {
  return (await api.get(path, { params: filter })).data as GetAllStudentsResponse;
};

export const getStudent = async (id: string) => {
  const { data } = (await api.get(`${path}/${id}`)).data as GetStudentResponse;
  return data;
};

export const deleteStudent = async (id: string) => {
  await api.delete(`${path}/${id}`);
};
