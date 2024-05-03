export interface User {
  id: number;
  firstName: string;
  lastName?: string;
  institutionId: string;
}
export interface Teacher extends User {
  institutionId: string;
}

export interface Student {
  id: string;
  firstName: string;
  lastName: string;
  nisn: string;
  email: string;
  classId: string;
}
