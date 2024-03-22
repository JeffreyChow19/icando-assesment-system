import { z } from "zod";

export const answerFormSchema = z.object({
  text: z.string({ required_error: "Choice is required" }).min(1, {
    message: "Choice must be at least 1 character",
  }),
  id: z.number(),
});

export const competencyFormSchema = z.object({
  name: z.string(),
  id: z.string(),
  numbering: z.string(),
});

export const questionFormSchema = z.object({
  text: z.string({ required_error: "Question text is required" }).min(2, {
    message: "Question text must be at least 2 characters.",
  }),
  choices: z.array(answerFormSchema).min(1),
  answerId: z.number(),
  competencies: z.array(competencyFormSchema),
  order: z.number(),
});
