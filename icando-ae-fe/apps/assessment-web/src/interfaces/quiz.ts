import { Question } from "./question";
import { Teacher } from "./user";

export interface QuizDetail {
  id: string;
  name: string | null;
  subject: string[] | null;
  passingGrade: number;
  lastPublishedAt: Date | null;
  startAt: string | null;
  endAt: string | null;
  createdBy: string;
  updatedBy: string;
  creator: Teacher;
  updater: Teacher;
  questions: Question[];
}
export interface StudentQuiz {
  id: string;
  name: string | null;
  subject: string[] | null;
  passingGrade: number;
  publishedAt: Date;
  duration: number;
  startAt: Date;
  endAt: Date;
}
