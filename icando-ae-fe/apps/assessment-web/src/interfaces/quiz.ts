import { Question } from './question';
import { Teacher, Student } from './user';

export interface Quiz {
  id: string;
  name: string;
  subject: string[];
  passingGrade: number;
  publishedAt: Date;
  duration: number;
  startAt: string;
  endAt: string;
  createdBy?: string;
  updatedBy?: string;
  creator?: Teacher;
  updater?: Teacher;
  questions?: Question[];
}

export interface StudentQuiz {
  id: string;
  createdAt: string;
  updatedAt: string;
  totalScore: number | null;
  correctCount: number | null;
  startedAt: string;
  completedAt: string | null;
  status: 'STARTED' | 'SUBMITTED' | 'NOT_STARTED';
  quiz_id: string;
  quiz?: Quiz;
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
