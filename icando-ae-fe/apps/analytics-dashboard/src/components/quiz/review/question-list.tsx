import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@ui/components/ui/table.tsx";
import { Choice } from "../../../interfaces/question.ts";
import { useState } from "react";
import {
  Check,
  ChevronsUpDown,
  GripVertical,
  SearchIcon,
  X,
} from "lucide-react";
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@ui/components/ui/collapsible.tsx";
import { Button } from "@ui/components/ui/button.tsx";
import { cn } from "@ui/lib/utils.ts";
import { Badge } from "@ui/components/ui/badge.tsx";
import { QuestionWithAnswer } from "../../../interfaces/quiz.ts";

const ChoiceList = ({
  choices,
  correctAnswer,
  studentAnswer,
}: {
  choices: Choice[];
  correctAnswer: number;
  studentAnswer: number | null;
}) => {
  const isCorrect = correctAnswer === studentAnswer;

  return (
    <div className="w-full flex flex-col">
      {choices.map((choice) => {
        const isShouldAnswer = correctAnswer === choice.id;
        const isStudentSelectedAnswer = studentAnswer === choice.id;

        return (
          <div
            key={choice.id}
            className={cn(
              isStudentSelectedAnswer &&
                `${isCorrect ? "bg-green-100" : studentAnswer !== null ? "bg-red-100" : ""}`,
              "flex w-full p-2 items-center gap-6 rounded-md",
            )}
          >
            <GripVertical className="size-3.5 text-muted-foreground opacity-50" />
            <p className="flex gap-x-2">
              {choice.text}{" "}
              {isStudentSelectedAnswer && (isCorrect ? <Check /> : <X />)}
              {isShouldAnswer && !isStudentSelectedAnswer && !isCorrect && (
                <Check className="text-green-700" />
              )}
            </p>
          </div>
        );
      })}
    </div>
  );
};

interface QuestionListProps {
  questions: QuestionWithAnswer[];
}

export const QuestionList = ({ questions }: QuestionListProps) => {
  const QuestionDetails = ({
    question,
    index,
  }: {
    question: QuestionWithAnswer;
    index: number;
  }) => {
    const [isOpen, setIsOpen] = useState<boolean>(true);
    return (
      <Collapsible
        open={isOpen}
        onOpenChange={setIsOpen}
        className="w-full space-y-2"
        asChild
      >
        <>
          <TableRow key={question.id}>
            <TableCell>{index + 1}</TableCell>
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
                    studentAnswer={
                      question.studentAnswer !== null
                        ? question.studentAnswer.answerId
                        : null
                    }
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
        {questions.length <= 0 && (
          <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
            <SearchIcon className="w-10 h-10" />
            No questions yet for this quiz
          </div>
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
        {questions.map((question, index) => (
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
