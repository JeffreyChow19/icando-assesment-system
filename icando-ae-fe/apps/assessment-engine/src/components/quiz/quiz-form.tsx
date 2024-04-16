import { QuestionList } from "./questions/question-list.tsx";
import { QuestionForm } from "./questions/question-form.tsx";
import { QuizDetail } from "../../interfaces/quiz.ts";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form } from "@ui/components/ui/form.tsx";
import { quizFormSchema } from "./quiz-schema.ts";
import { QuizInfo } from "./quiz-info.tsx";
import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from "@ui/components/ui/tabs.tsx";
import { getQuiz, updateQuiz } from "../../services/quiz.ts";
import { toast } from "@ui/components/ui/use-toast.ts";
import { onErrorToast } from "../ui/error-toast.tsx";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useCallback } from "react";

export const QuizForm = ({ quiz }: { quiz: QuizDetail }) => {
  const queryClient = useQueryClient();
  const form = useForm<z.infer<typeof quizFormSchema>>({
    resolver: zodResolver(quizFormSchema),
    defaultValues: {
      name: quiz.name || "",
      subject: quiz.subject || [],
      passingGrade: quiz.passingGrade,
      questions: quiz.questions || [],
    },
  });

  const onQuestionUpdate = useCallback(async () => {
    const result = await getQuiz(quiz.id);
    form.setValue("questions", result.questions);
  }, [form, quiz.id]);

  const nextOrder = () => {
    const questions = form.getValues("questions");
    if (questions.length === 0) {
      return 0;
    }

    return (
      questions.reduce((max, curr) => {
        if (max < curr.order) {
          return curr.order;
        }

        return max;
      }, 0) + 1
    );
  };

  const mutation = useMutation({
    mutationFn: async (payload: z.infer<typeof quizFormSchema>) => {
      await updateQuiz({
        id: quiz.id,
        name: payload.name,
        subject: payload.subject,
        passingGrade: payload.passingGrade,
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["quiz", quiz.id] });
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
      <form
        onSubmit={form.handleSubmit((val) => {
          mutation.mutate(val);
        })}
      >
        <div className="mb-4">
          <h1 className="text-center text-xl font-bold">
            {form.watch("name") || "Untitled Quiz"}
          </h1>
        </div>
        <Tabs defaultValue="information" className="w-flul">
          <TabsList>
            <TabsTrigger value="information">Information</TabsTrigger>
            <TabsTrigger value="questions">Questions</TabsTrigger>
          </TabsList>
          <TabsContent value="information">
            <QuizInfo isPending={mutation.isPending} />
          </TabsContent>
          <TabsContent value="questions">
            <div className="flex w-full justify-end">
              <QuestionForm
                type="new"
                quizId={quiz.id}
                onSuccess={() => onQuestionUpdate()}
                nextOrder={nextOrder}
              />
            </div>
            <QuestionList
              onSuccess={() => onQuestionUpdate()}
              quizId={quiz.id}
            />
          </TabsContent>
        </Tabs>
      </form>
    </Form>
  );
};
