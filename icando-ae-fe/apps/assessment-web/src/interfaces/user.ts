export interface User {
  id: number;
  username: string;
  name: string;
  role: "ADMIN" | "USER";
}
export interface Teacher extends User {
  institutionId: string;
}

export interface Student{
  id: string;
  firstName: string;
  lastName: string;
  nisn: string;
  email: string;
  classId: string;
}