import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@ui/components/ui/dialog.tsx";
import { Button } from "@ui/components/ui/button.tsx";
import { z } from "zod";
import { Input } from "@ui/components/ui/input.tsx";
import { Question } from "../../../interfaces/question.ts";
import { EditIcon, SearchIcon, XIcon } from "lucide-react";
import { useForm } from "react-hook-form";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@ui/components/ui/form.tsx";
import { zodResolver } from "@hookform/resolvers/zod";
import { Textarea } from "@ui/components/ui/textarea.tsx";
import { useMemo, useState } from "react";
import { RadioGroup, RadioGroupItem } from "@ui/components/ui/radio-group.tsx";
import { cn } from "@ui/lib/utils.ts";
import { ScrollArea } from "@ui/components/ui/scroll-area.tsx";
import { useQuery } from "@tanstack/react-query";
import { getAllCompetency } from "../../../services/competency.ts";
import { createQuestion } from "../../../services/quiz.ts";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@ui/components/ui/table.tsx";
import { Pagination } from "../../pagination.tsx";
import { LoadingComponent } from "../../loading.tsx";
import { Checkbox } from "@ui/components/ui/checkbox.tsx";
import { Competency } from "../../../interfaces/competency.ts";
import { CompetencyBadge } from "./competency-badge.tsx";
import { useMutation } from "@tanstack/react-query";
import { toast } from "@ui/components/ui/use-toast.ts";
import { onErrorToast } from "../../ui/error-toast.tsx";
interface QuestionFormProps {
  type: "edit" | "new";
  question?: Question;
}
export const QuestionForm = ({ type, question }: QuestionFormProps) => {
  const [step, setStep] = useState<number>(0);
  // const numSteps = 2;
  const [answerLength, setAnswerLength] = useState(
    question?.choices?.length || 1,
  );
  const answerFormSchema = z.object({
    text: z.string({ required_error: "Choice is required" }).min(1, {
      message: "Choice must be at least 1 character",
    }),
    id: z.number(),
  });

  const competencyFormSchema = z.object({
    name: z.string(),
    id: z.string(),
    numbering: z.string(),
  });

  const secondFormSchema = z.object({
    competencies: z.array(competencyFormSchema).min(1),
  });

  const formSchema = z.object({
    text: z.string({ required_error: "Question text is required" }).min(2, {
      message: "Question text must be at least 2 characters.",
    }),
    choices: z.array(answerFormSchema).min(1),
    correctAnswer: z.number(),
  });

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      text: question?.text || "",
      choices: question?.choices || [
        {
          text: "Option 1",
          id: 0,
        },
      ],
      correctAnswer:
        question?.answerId && question?.choices
          ? question.choices.findIndex(
              (choice) => choice.id == question.answerId,
            )
          : 0,
    },
  });

  const secondForm = useForm<z.infer<typeof secondFormSchema>>({
    resolver: zodResolver(secondFormSchema),
    defaultValues: {
      competencies: question?.competencies || [],
    },
  });

  const choicesState = form.watch("choices");

  // choices is always
  const largestId = useMemo(() => {
    const choices = [...form.getValues("choices")];
    const sortedChoices = choices.sort((a, b) => b.id - a.id);
    return sortedChoices[0].id;
  }, [choicesState]);

  const removeAnswer = (index: number) => {
    const currAnswer = form.getValues("choices");
    const correctAnswer = form.getValues("correctAnswer");
    currAnswer.splice(index, 1);
    form.setValue("choices", currAnswer);
    if (index == correctAnswer) form.setValue("correctAnswer", 0);
    setAnswerLength(answerLength - 1);
  };

  const addAnswer = () => {
    const currAnswer = form.getValues("choices");
    currAnswer.push({
      text: `Option ${answerLength + 1}`,
      id: largestId + 1,
    });
    form.setValue("choices", currAnswer);
    setAnswerLength(answerLength + 1);
  };

  const [page, setPage] = useState<number>(1);
  const [open, setOpen] = useState<boolean>(false);

  const CompetencyTable = () => {
    const { data } = useQuery({
      queryKey: ["competency", page],
      queryFn: () => getAllCompetency({ page, limit: 8 }),
      retry: false,
    });
    const competencies = data?.competencies;
    const meta = data?.meta;
    const checkCompetency = (id: string) => {
      return secondForm.watch("competencies").findIndex((competency) => {
        return competency.id == id;
      });
    };

    const onSelectCompetency = (
      { id, name, numbering }: Competency,
      checked: boolean | "indeterminate",
    ) => {
      if (checked === "indeterminate") return;
      const newCompetencies = [...secondForm.getValues("competencies")];

      if (checked) {
        newCompetencies.push({
          id,
          name,
          numbering,
        });
      } else {
        const index = checkCompetency(id);
        if (index > -1) {
          newCompetencies.splice(index, 1);
        }
      }
      secondForm.setValue("competencies", newCompetencies);
    };

    return (
      <div className="flex flex-col gap-2 px-4 h-[70vh]">
        <FormLabel>Competencies</FormLabel>
        <Table>
          <TableCaption>
            {competencies ? (
              competencies.length <= 0 ? (
                <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
                  <SearchIcon className="w-10 h-10" />
                  No questions yet for this quiz
                </div>
              ) : (
                meta &&
                meta.totalPage > 1 && (
                  <div className="flex w-full justify-end">
                    <Pagination
                      page={page}
                      totalPage={meta.totalPage}
                      setPage={setPage}
                    />
                  </div>
                )
              )
            ) : (
              <LoadingComponent />
            )}
          </TableCaption>
          <TableHeader>
            <TableRow>
              <TableHead className="w-[4vw]" />
              <TableHead className="w-[4vw]">#</TableHead>
              <TableHead>Competency</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {competencies?.map((competency) => (
              <TableRow key={competency.id}>
                <TableCell>
                  <Checkbox
                    checked={checkCompetency(competency.id) > -1}
                    onCheckedChange={(checked) => {
                      onSelectCompetency(competency, checked);
                    }}
                  />
                </TableCell>
                <TableCell>{competency.numbering}</TableCell>
                <TableCell>{competency.name}</TableCell>
                <TableCell>{competency.description}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        <div className="flex w-full gap-2 flex-wrap mt-2">
          {secondForm.watch("competencies").map((competency) => (
            <CompetencyBadge competency={competency} key={competency.id} />
          ))}
        </div>
      </div>
    );
  };
  const answerElements = Array.from({ length: answerLength }, (_, index) => (
    <div key={index} className="grid grid-cols-1">
      <FormField
        control={form.control}
        name={`choices.${index}.text`}
        render={({ field }) => (
          <FormItem>
            <FormLabel>{`Option ${index + 1}`} </FormLabel>
            <div
              className={cn(
                form.watch("correctAnswer") == index && "bg-green-100",
                "flex flex-row justify-between gap-2 items-center py-2 px-3 rounded-md",
              )}
            >
              <RadioGroupItem
                value={index.toString()}
                checked={index == form.watch("correctAnswer")}
                className="mr-2"
              />
              <FormControl>
                <Input placeholder={`Option ${index + 1}`} {...field} />
              </FormControl>
              <Button
                className="h-8 w-8 p-1"
                variant="ghost"
                onClick={(e) => {
                  e.preventDefault();
                  removeAnswer(index);
                }}
                disabled={answerLength === 1}
              >
                <XIcon size="12" />
              </Button>
            </div>
            <FormMessage />
          </FormItem>
        )}
      />
    </div>
  ));

  const mutation = useMutation({
    mutationFn: async (
      payload: z.infer<typeof formSchema> & z.infer<typeof secondFormSchema>,
    ) => {
      const { competencies, ...others } = payload;
      if (type == "new") {
        await createQuestion({
          ...others,
          competencies: competencies.map((competency) => competency.id),
        });
      }
    },

    onSuccess: () => {
      toast({
        description: "Successfully created question",
      });
    },
    onError: (e: Error) => {
      onErrorToast(e);
    },
  });

  const handleSubmit = () => {
    const firstPayload = form.getValues();
    const secondPayload = secondForm.getValues();
    mutation.mutate({ ...firstPayload, ...secondPayload });
  };

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
        {step == 0 && (
          <Form {...form}>
            <form
              onSubmit={form.handleSubmit(() => setStep(1))}
              className="space-y-8"
            >
              <DialogHeader>
                <DialogTitle>
                  {type == "new" ? "New Question" : "Edit Question"}
                </DialogTitle>
              </DialogHeader>
              <ScrollArea className="h-[70vh] px-4">
                {/*<Progress />*/}
                {step == 0 && (
                  <>
                    <FormField
                      control={form.control}
                      name="text"
                      render={({ field }) => (
                        <FormItem>
                          <FormLabel>Question Text</FormLabel>
                          <FormControl>
                            <Textarea placeholder="Question Text" {...field} />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                    <div className="flex flex-col gap-2 mt-4">
                      <RadioGroup
                        onValueChange={(val) =>
                          form.setValue("correctAnswer", parseInt(val))
                        }
                      >
                        {answerElements}
                      </RadioGroup>
                      <Button
                        variant="ghost"
                        className="text-primary p-2 w-fit"
                        onClick={(e) => {
                          e.preventDefault();
                          addAnswer();
                        }}
                        disabled={answerLength === 6}
                      >
                        + Add choice
                      </Button>
                    </div>
                  </>
                )}
              </ScrollArea>
              <DialogFooter>
                <Button variant="outline" type="submit">
                  Next
                </Button>
              </DialogFooter>
            </form>
          </Form>
        )}

        {step == 1 && (
          <Form {...form}>
            <form
              onSubmit={form.handleSubmit(handleSubmit)}
              className="space-y-8"
            >
              <DialogHeader>
                <DialogTitle>
                  {type == "new" ? "New Question" : "Edit Question"}
                </DialogTitle>
              </DialogHeader>
              <CompetencyTable />
              <DialogFooter>
                <div className="w-full flex gap-4 justify-end">
                  <Button variant="outline" onClick={() => setStep(0)}>
                    Back
                  </Button>
                  <Button type="submit">Save</Button>
                </div>
              </DialogFooter>
            </form>
          </Form>
        )}
      </DialogContent>
    </Dialog>
  );
};
