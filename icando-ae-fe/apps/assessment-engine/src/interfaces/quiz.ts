import { Question } from "./question";
import { Teacher } from "./user";

export interface QuizDetail {
  id: string;
  name: string | null;
  subject: string | null;
  passingGrade: number;
  publishedAt: string | null;
  deadline: string | null;
  creator: Teacher;
  updater: Teacher;
  questions: Question[];
}
