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
        <CardContent className="flex flex-col gap-2 w-full mt-6">
          <div className="grid grid-cols-4 ">
            <p className="col-span-1 font-semibold">Subjects</p>
            {!isLoading && quiz && (
              <p className="col-span-3">{quiz.subject.join(", ")}</p>
            )}
            {isLoading && <Skeleton className="w-4/5 h-4" />}
          </div>
          <div className="grid grid-cols-4 ">
            <p className="col-span-1 font-semibold">Published at</p>
            {!isLoading && quiz && (
              <p className="col-span-3">
                {formatDateHour(new Date(quiz.publishedAt))}
              </p>
            )}
            {isLoading && <Skeleton className="w-4/5 h-4" />}
          </div>
          <div className="grid grid-cols-4 ">
            <p className="col-span-1 font-semibold">Start at</p>
            {!isLoading && quiz && (
              <p className="col-span-3">
                {formatDateHour(new Date(quiz.startAt))}
              </p>
            )}
            {isLoading && <Skeleton className="w-4/5 h-4" />}
          </div>
          <div className="grid grid-cols-4 ">
            <p className="col-span-1 font-semibold">End at</p>
            {!isLoading && quiz && (
              <p className="col-span-3">
                {formatDateHour(new Date(quiz.endAt))}
              </p>
            )}
            {isLoading && <Skeleton className="w-4/5 h-4" />}
          </div>
          <div className="grid grid-cols-4 ">
            <p className="col-span-1 font-semibold">Passing Grade</p>
            {!isLoading && quiz && (
              <p className="col-span-3">{quiz.passingGrade}</p>
            )}
            {isLoading && <Skeleton className="w-4/5 h-4" />}
          </div>
        </CardContent>
      </Card>
    </Fragment>
  );
};
