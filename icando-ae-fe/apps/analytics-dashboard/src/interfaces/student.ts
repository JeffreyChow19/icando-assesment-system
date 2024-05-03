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
  quiz_id: string;
  total_score: number;
  correct_count: number;
  completed_at: string;
  name: string;
  passing_grade: number;
}
