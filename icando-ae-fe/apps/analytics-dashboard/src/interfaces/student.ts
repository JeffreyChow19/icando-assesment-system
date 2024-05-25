export interface Student {
  id: string;
  firstName: string;
  lastName: string;
  nisn: string;
  email: string;
  classId?: string;
}
export interface StudentData {
  student: {
    nisn: string;
    firstName: string;
    lastName: string;
    email: string;
  };
  class: {
    name: string;
    grade: string;
  };
}

export interface StudentPerformance {
  quizzesPassed: number;
  quizzesFailed: number;
}

export interface StudentQuiz {
  id: string;
  quizId: string;
  totalScore: number;
  correctCount: number;
  completedAt: string;
  name: string;
  passingGrade: number;
}

export interface StudentSubmissions{
  className: string;
  grade: string;
  quizName: string;
  firstName: string;
  lastName: string;
  publishedAt: Date;
  completedAt: Date;
}