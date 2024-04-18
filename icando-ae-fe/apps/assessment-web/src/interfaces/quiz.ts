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
