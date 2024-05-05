import { Question } from "../interfaces/question.ts";
import { cn } from "@ui/lib/utils.ts";
import { CheckIcon, XIcon } from "lucide-react";

export const QuestionAnswerCard = ({
  question,
  answerId,
  questionNumber,
}: {
  question: Question;
  answerId: number | null;
  questionNumber: number;
}) => {
  return (
    <div className="flex flex-col gap-3 mt-2 mb-2">
      <div>
        {questionNumber}. {question.text}{" "}
        {!answerId && (
          <span className="italic text-muted-foreground">(Tidak dijawab)</span>
        )}
      </div>
      {question.choices.map((choice) => {
        const isCorrectAnswer =
          choice.id == answerId && question.answerId == choice.id;
        const isWrongAnswer =
          choice.id == answerId && question.answerId != choice.id;

        return (
          <div
            key={choice.id}
            className={cn(
              "flex py-2 px-4 items-center rounded-md justify-between",
              isCorrectAnswer
                ? "bg-green-200"
                : isWrongAnswer
                  ? "bg-rose-200"
                  : "bg-white",
            )}
          >
            {choice.text}
            {isCorrectAnswer && <CheckIcon className="text-green-600 size-5" />}
            {isWrongAnswer && <XIcon className="text-rose-600 size-5" />}
          </div>
        );
      })}
      <div className="flex flex-col gap-2">
        <p className="italic text-muted-foreground">Jawaban benar:</p>
        <p>
          {
            question.choices.find((choice) => choice.id == question.answerId)
              ?.text
          }
        </p>
      </div>
    </div>
  );
};
