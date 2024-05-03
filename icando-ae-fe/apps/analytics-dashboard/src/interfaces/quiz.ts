import { Question } from "./question";
import { Student, Teacher } from "./user";

export interface Quiz {
  id: string;
  name: string;
  subject: string[];
  passingGrade: number;
  publishedAt: Date;
  lastPublishedAt?: Date;
  duration: number;
  startAt: string;
  endAt: string;
  createdBy?: string;
  updatedBy?: string;
  creator?: Teacher;
  updater?: Teacher;
  questions?: Question[];
  hasNewerVersion?: boolean;
}

export interface StudentQuiz {
  id: string;
  createdAt: string;
  updatedAt: string;
  totalScore: number | null;
  correctCount: number | null;
  startedAt: string;
  completedAt: string | null;
  status: "STARTED" | "SUBMITTED" | "NOT_STARTED";
  quiz_id: string;
  quiz?: Quiz;
  studentId: string;
  student: Student | null;
  studentAnswers: StudentAnswer[] | null;
}

export interface StudentAnswer {
  questionId: string;
  answerId: number;
}

export interface QuestionWithAnswer extends Question {
  studentAnswer: StudentAnswer | null;
}
