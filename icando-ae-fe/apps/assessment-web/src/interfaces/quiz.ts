import { Question } from "./question";
import { Teacher, Student } from "./user";

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

export interface QuizAttempt {
  id: string;
  createdAt: string;
  updatedAt: string;
  totalScore: number | null;
  correctCount: number | null;
  startedAt: string;
  completedAt: string | null;
  status: 'STARTED' | 'SUBMITTED' | 'NOT_STARTED';
  quiz_id: string;
  studentId: string;
  student: Student | null;
  studentAnswers: StudentAnswer[] | null; 
}

// todo fix student answer interface
export interface StudentAnswer {
  questionId: string;
  answerId: string;
  isCorrect: boolean;
}