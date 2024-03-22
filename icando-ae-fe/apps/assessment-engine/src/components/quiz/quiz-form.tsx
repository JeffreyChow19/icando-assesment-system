import { QuestionList } from "./questions/question-list.tsx";
import { QuestionForm } from "./questions/question-form.tsx";
import { QuizDetail } from "../../interfaces/quiz.ts";
import { z } from "zod";
import { useMutation } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { updateQuiz } from "../../services/quiz.ts";
import { toast } from "@ui/components/ui/use-toast.ts";
import { onErrorToast } from "../ui/error-toast.tsx";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@ui/components/ui/form.tsx";
import { Input } from "@ui/components/ui/input.tsx";

const competencySchema = z.object({
  id: z.string(),
  name: z.string(),
  numbering: z.string(),
  description: z.string(),
});

const questionSchema = z.object({
  id: z.string(),
  text: z.string(),
  answerId: z.number(),
  choices: z
    .object({
      id: z.number(),
      text: z.string(),
    })
    .array(),
  quizId: z.string(),
  competencies: competencySchema.array(),
});

export const quizFormSchema = z.object({
  name: z.string({ required_error: "Name should not be empty" }),
  subject: z.string({ required_error: "Subject should not be empty" }),
  passingGrade: z.coerce
    .number({ required_error: "Passing grade should not be empty" })
    .min(0)
    .max(100),
  questions: questionSchema.array(),
});

export const QuizForm = ({ quiz }: { quiz: QuizDetail }) => {
  const form = useForm<z.infer<typeof quizFormSchema>>({
    resolver: zodResolver(quizFormSchema),
    defaultValues: {
      name: quiz.name || "",
      subject: quiz.subject || "",
      passingGrade: quiz.passingGrade,
      questions: quiz.questions,
    },
  });

  const mutation = useMutation({
    mutationFn: async (payload: z.infer<typeof quizFormSchema>) => {
      await updateQuiz({
        id: quiz.id,
        name: payload.name,
        subject: payload.subject,
        passingGrade: payload.passingGrade,
        deadline: quiz.deadline,
      });
    },
    onSuccess: () => {
      toast({
        description: `Quiz updated`,
      });
    },
    onError: (err: Error) => {
      onErrorToast(err);
    },
  });

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit((val) => mutation.mutate(val))}>
        <div className={"flex flex-col gap-4 max-w-[500px]"}>
          <FormField
            control={form.control}
            name={"name"}
            render={({ field }) => (
              <FormItem>
                <FormLabel>Quiz Name</FormLabel>
                <FormControl>
                  <Input {...field} placeholder={"Enter quiz name"} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name={"subject"}
            render={({ field }) => (
              <FormItem>
                <FormLabel>Quiz Subject</FormLabel>
                <FormControl>
                  <Input {...field} placeholder={"Enter quiz subject"} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name={"passingGrade"}
            render={({ field }) => (
              <FormItem>
                <FormLabel>Passing Grade</FormLabel>
                <FormDescription>
                  Passing grade value should be between 0 and 100
                </FormDescription>
                <FormControl>
                  <Input
                    type="number"
                    min="0"
                    max="100"
                    {...field}
                    placeholder={"Enter quiz passing grade"}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <div className="flex w-full justify-end">
          <QuestionForm type="new" />
        </div>
        <QuestionList />
      </form>
    </Form>
  );
};
