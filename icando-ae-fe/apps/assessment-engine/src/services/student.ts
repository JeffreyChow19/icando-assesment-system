import { api } from '../utils/api.ts';

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


const path = '/designer/student';

export const createStudent = async (payload: CreateStudentPayload) => {
  await api.post(path, payload);
};

export const updateStudent = async (payload: UpdateStudentPayload, id: string) => {
  await api.patch(`${path}/${id}`, payload);
};

