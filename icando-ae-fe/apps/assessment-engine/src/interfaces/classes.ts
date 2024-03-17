import { Student } from "./student";
import { Teacher } from "./teacher";

export interface Class {
  id: string;
  name: string;
  grade: string;
  teacherId: string;
  students?: Student[];
  teacher?: Teacher;
}
