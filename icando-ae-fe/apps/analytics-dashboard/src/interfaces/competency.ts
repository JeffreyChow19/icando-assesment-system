export interface Competency {
  id: string;
  name: string;
  numbering: string;
  description: string;
}

export interface StudentCompetency {
  competencyId: string;
  competencyName: string;
  correctCount: number;
  totalCount: number;
}
