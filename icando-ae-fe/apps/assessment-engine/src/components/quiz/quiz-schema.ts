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
  order: z.number(),
  competencies: competencySchema.array(),
});

export const quizFormSchema = z.object({
  name: z.string({ required_error: "Name should not be empty" }).min(2),
  subject: z
    .array(z.string({ required_error: "Subject should not be empty" }).min(2))
    .min(1, "Subject should not be empty"),
  passingGrade: z.coerce
    .number({ required_error: "Passing grade should not be empty" })
    .min(0)
    .max(100),
  questions: questionSchema.array(),
});
