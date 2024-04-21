import { Class } from "./classes";
import { Question } from "./question";
import { Teacher } from "./user";

export interface QuizDetail {
  id: string;
  name: string | null;
  subject: string[] | null;
  passingGrade: number;
  lastPublishedAt: Date | null;
  startAt: Date | null;
  endAt: Date | null;
  createdBy: string;
  updatedBy: string;
  duration: number | null;
  creator: Teacher;
  updater: Teacher;
  questions: Question[];
  classes: Class[];
}
