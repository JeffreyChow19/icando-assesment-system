import {
  Dialog,
  DialogContent,
  DialogTrigger,
  DialogHeader,
  DialogTitle,
} from "@ui/components/ui/dialog.tsx";
import { Button } from "@ui/components/ui/button.tsx";
import { z } from "zod";
import { Question } from "../../../interfaces/question.ts";
import { EditIcon } from "lucide-react";
import { useForm } from "react-hook-form";
import { Form } from "@ui/components/ui/form.tsx";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { createQuestion, updateQuestion } from "../../../services/quiz.ts";
import { useMutation } from "@tanstack/react-query";
import { toast } from "@ui/components/ui/use-toast.ts";
import { onErrorToast } from "../../ui/error-toast.tsx";
import { questionFormSchema } from "./question-schema.ts";
import { QuestionStep1 } from "./question-step-1.tsx";
import { QuestionStep2 } from "./question-step-2.tsx";

interface QuestionFormProps {
  type: "edit" | "new";
  question?: Question;
  quizId?: string;
  onSuccess: () => void;
  nextOrder: () => number;
}

export const QuestionForm = ({
  type,
  question,
  quizId,
  onSuccess,
  nextOrder,
}: QuestionFormProps) => {
  const [step, setStep] = useState<number>(0);

  const form = useForm<z.infer<typeof questionFormSchema>>({
    resolver: zodResolver(questionFormSchema),
    defaultValues: {
      text: question?.text || "",
      choices: question?.choices || [
        {
          text: "Option 1",
          id: 0,
        },
      ],
      answerId:
        question?.answerId && question?.choices
          ? question.choices.findIndex(
              (choice) => choice.id == question.answerId,
            )
          : 0,
      competencies: question?.competencies || [],
      order: question?.order || 0,
    },
  });

  const [open, setOpen] = useState<boolean>(false);

  const mutation = useMutation({
    mutationFn: async (payload: z.infer<typeof questionFormSchema>) => {
      const { competencies, ...others } = payload;
      if (type == "new" && quizId) {
        await createQuestion(quizId, {
          ...others,
          competencies: competencies.map((competency) => competency.id),
          order: nextOrder(),
        });
      } else if (type === "edit" && quizId && question) {
        await updateQuestion(quizId, question.id, {
          ...others,
          competencies: competencies.map((competency) => competency.id),
        });
      }
    },
    onSuccess: () => {
      onSuccess();
      toast({
        description: `Successfully ${type === "new" ? "added" : "updated"} question`,
      });
      setOpen(false);
      setStep(0);
      form.reset();
    },
    onError: (e: Error) => {
      onErrorToast(e);
    },
  });

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        {type == "new" ? (
          <Button>New Question</Button>
        ) : (
          <Button className="h-8 w-8 p-1" variant="secondary">
            <EditIcon size="12" />
          </Button>
        )}
      </DialogTrigger>
      <DialogContent className="min-w-[70vw]">
        <Form {...form}>
          <form
            id="question"
            onSubmit={(e) => {
              e.preventDefault();
              e.stopPropagation();
              form.handleSubmit((payload) => {
                mutation.mutate(payload);
              })(e);
            }}
            className="space-y-8"
          >
            <DialogHeader>
              <DialogTitle>
                {type == "new" ? "New Question" : "Edit Question"}
              </DialogTitle>
            </DialogHeader>
            {step == 0 && <QuestionStep1 next={() => setStep(1)} />}
            {step == 1 && <QuestionStep2 prev={() => setStep(0)} />}
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
};
