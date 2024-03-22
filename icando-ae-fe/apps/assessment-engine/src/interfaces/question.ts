import { Competency } from "./competency.ts";

export interface Choice {
  id: number;
  text: string;
}

export interface Question {
  id: string;
  choices: Choice[];
  text: string;
  answerId: number;
  quizId: string;
  competencies: Competency[];
}
