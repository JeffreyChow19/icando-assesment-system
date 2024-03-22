import { z } from "zod";

export const competencySchema = z.object({
  id: z.string(),
  name: z.string(),
  numbering: z.string(),
  description: z.string(),
});

export const questionSchema = z.object({
  id: z.string(),
  text: z.string(),
  answerId: z.number(),
  choices: z
    .object({
      id: z.number(),
      text: z.string(),
    })
    .array(),
  competencies: competencySchema.array(),
});

export const quizFormSchema = z.object({
  name: z.string({ required_error: "Name should not be empty" }).min(2),
  subject: z.string({ required_error: "Subject should not be empty" }).min(2),
  passingGrade: z.coerce
    .number({ required_error: "Passing grade should not be empty" })
    .min(0)
    .max(100),
  questions: questionSchema.array(),
});
