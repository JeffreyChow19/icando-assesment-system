import { Student, Teacher } from "./user.ts";

export interface Class {
  id: string;
  name: string;
  grade: string;
  institutionId: string;
  students?: Student[];
  teachers?: Teacher[];
}
