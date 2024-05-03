import { Student } from "./student";

export interface Class {
  id: string;
  name: string;
  grade: string;
  institutionId: string;
  students?: Student[];
}
