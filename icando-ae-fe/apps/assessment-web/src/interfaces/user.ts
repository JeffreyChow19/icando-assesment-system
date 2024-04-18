export interface User {
  id: number;
  username: string;
  name: string;
  role: "ADMIN" | "USER";
}
export interface Teacher extends User {
  institutionId: string;
}
