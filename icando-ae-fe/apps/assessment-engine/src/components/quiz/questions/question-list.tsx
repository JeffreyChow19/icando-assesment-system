import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@ui/components/ui/table.tsx";
import { Choice, Question } from "../../../interfaces/question.ts";
import { useState } from "react";
import { ChevronsUpDown, GripVertical, SearchIcon, Trash2 } from "lucide-react";
import { Pagination } from "../../pagination.tsx";
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@ui/components/ui/collapsible.tsx";
import { Button } from "@ui/components/ui/button.tsx";
import { cn } from "@ui/lib/utils.ts";
import { QuestionForm } from "./question-form.tsx";
import { Badge } from "@ui/components/ui/badge.tsx";
import { useFormContext } from "react-hook-form";
import { z } from "zod";
import { quizFormSchema } from "../quiz-schema.ts";
import { useConfirm } from "../../../context/alert-dialog.tsx";
import { deleteQuestion } from "../../../services/quiz.ts";
import { toast } from "@ui/components/ui/use-toast.ts";

const ChoiceList = ({
  choices,
  correctAnswer,
}: {
  choices: Choice[];
  correctAnswer: number;
}) => {
  return (
    <div className="w-full flex flex-col">
      {choices.map((choice) => {
        return (
          <div
            key={choice.id}
            className={cn(
              correctAnswer == choice.id && "bg-green-100",
              "flex w-full p-2 items-center gap-6 rounded-md",
            )}
          >
            <GripVertical className="size-3.5 text-muted-foreground opacity-50" />
            <p>{choice.text}</p>
          </div>
        );
      })}
    </div>
  );
};

interface QuestionListProps {
  onSuccess: () => void;
  quizId: string;
}

export const QuestionList = ({ onSuccess, quizId }: QuestionListProps) => {
  const form = useFormContext<z.infer<typeof quizFormSchema>>();
  const questions = form.watch("questions");

  const [page, setPage] = useState<number>(1);
  const questionsPerPage = 10;
  const startIndex = (page - 1) * questionsPerPage;
  const endIndex = startIndex + questionsPerPage;
  const paginatedQuestions = questions.slice(startIndex, endIndex);
  const totalPage = Math.ceil(questions.length / questionsPerPage);

  const confirm = useConfirm();

  const QuestionDetails = ({
    question,
    index,
  }: {
    question: Question;
    index: number;
  }) => {
    const [isOpen, setIsOpen] = useState<boolean>(false);
    return (
      <Collapsible
        open={isOpen}
        onOpenChange={setIsOpen}
        className="w-full space-y-2"
        asChild
      >
        <>
          <TableRow key={question.id}>
            <TableCell>{(page - 1) * questionsPerPage + index + 1}</TableCell>
            <TableCell>{question.text}</TableCell>
            <TableCell className="text-right">
              <CollapsibleTrigger asChild>
                <Button variant="ghost" size="sm" className="w-9 p-0">
                  <ChevronsUpDown className="h-4 w-4" />
                  <span className="sr-only">Toggle</span>
                </Button>
              </CollapsibleTrigger>
            </TableCell>
          </TableRow>
          <CollapsibleContent asChild>
            <TableRow className="bg-muted">
              <TableCell colSpan={4}>
                <div className="w-full flex flex-col gap-4 px-8 py-4">
                  <h2 className="font-bold">Choices</h2>
                  <ChoiceList
                    choices={question.choices}
                    correctAnswer={question.answerId}
                  />
                  <h2 className="font-bold">Competencies</h2>
                  <div className="flex flex-wrap gap-1">
                    {question.competencies.map((competency) => {
                      return (
                        <Badge key={competency.id} variant="outline">
                          {competency.numbering} - {competency.name}
                        </Badge>
                      );
                    })}
                  </div>
                  <div className="flex w-full justify-end">
                    <div className="flex gap-2">
                      <QuestionForm
                        type="edit"
                        quizId={quizId}
                        question={question}
                        onSuccess={onSuccess}
                      />
                      <Button
                        type="button"
                        className="h-8 w-8 p-1"
                        variant="destructive"
                        onClick={() => {
                          confirm({
                            title: "Are you sure?",
                            body: "Are you sure want to delete this question?",
                          }).then((res) => {
                            if (res) {
                              deleteQuestion(quizId, question.id).then(() => {
                                onSuccess();
                                toast({
                                  description: "Question deleted successfully",
                                });
                              });
                            }
                          });
                        }}
                      >
                        <Trash2 size="12" />
                      </Button>
                    </div>
                  </div>
                </div>
              </TableCell>
            </TableRow>
          </CollapsibleContent>
        </>
      </Collapsible>
    );
  };

  return (
    <Table>
      <TableCaption>
        {questions.length <= 0 ? (
          <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
            <SearchIcon className="w-10 h-10" />
            No questions yet for this quiz
          </div>
        ) : (
          totalPage > 1 && (
            <div className="flex w-full justify-end">
              <Pagination page={page} totalPage={totalPage} setPage={setPage} />
            </div>
          )
        )}
      </TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="w-[4vw]">#</TableHead>
          <TableHead>Question</TableHead>
          <TableHead className="text-right">Action</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {paginatedQuestions.map((question, index) => (
          <QuestionDetails
            question={question}
            index={index}
            key={question.id}
          />
        ))}
      </TableBody>
    </Table>
  );
};
