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

export const quizPublishFormSchema = z
  .object({
    quizId: z.string(),
    quizDuration: z.coerce
      .number({ required_error: "Quiz duration should not be empty" })
      .int(),
    startAt: z.coerce.date({
      required_error: "Quiz start time should not be empty",
    }),
    endAt: z.coerce.date({
      required_error: "Quiz end time should not be empty",
    }),
    assignedClasses: z.array(z.string()).min(1, "Select at least one class"),
  })
  .refine(
    (data) => {
      return data.endAt > data.startAt;
    },
    {
      message: "Quiz end time should be greater than start time",
      path: ["endAt"],
    },
  )
  .refine(
    (data) => {
      const epochStart = data.startAt.getTime();
      const epochEnd = data.endAt.getTime();

      return data.quizDuration * 60 * 1000 <= epochEnd - epochStart;
    },
    {
      message:
        "Quiz duration should be less than or equal to quiz availability schedule",
      path: ["quizDuration"],
    },
  );
