import { Fragment } from "react";
import { Card, CardContent } from "@ui/components/ui/card.tsx";
import { formatDateHour } from "../../utils/format-date.ts";
import { Quiz } from "../../interfaces/quiz.ts";
import { Skeleton } from "@ui/components/ui/skeleton.tsx";

export const QuizDetailCard = ({
  quiz,
  isLoading,
}: {
  quiz?: Quiz;
  isLoading?: boolean;
}) => {
  return (
    <Fragment>
      <Card className="w-full">
        <CardContent className="grid grid-cols-4 w-full mt-6">
          <div className="flex flex-col gap-2 col-span-1 font-semibold">
            <p>Subjects</p>
            <p>Published at</p>
            <p>Start at</p>
            <p>End at</p>
            <p>Passing grade</p>
          </div>
          <div className="flex flex-col gap-2 col-span-3">
            {isLoading && (
              <>
                <Skeleton className="w-4/5 h-4" />
                <Skeleton className="w-4/5 h-4" />
                <Skeleton className="w-4/5 h-4" />
                <Skeleton className="w-4/5 h-4" />
                <Skeleton className="w-4/5 h-4" />
              </>
            )}
            {quiz && (
              <>
                <p>{quiz.subject.join(", ")}</p>
                <p>{formatDateHour(new Date(quiz.publishedAt))}</p>
                <p>{formatDateHour(new Date(quiz.startAt))}</p>
                <p>{formatDateHour(new Date(quiz.endAt))}</p>
                <p>{quiz.passingGrade}</p>
              </>
            )}
          </div>
        </CardContent>
      </Card>
    </Fragment>
  );
};
