import { Fragment } from "react";
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@ui/components/ui/card.tsx";
import {
  formatDate,
  formatDateHour,
  formatHour,
} from "../../utils/format-date.ts";
import { Quiz } from "../../interfaces/quiz.ts";

export const QuizDetailCard = ({ quiz }: { quiz: Quiz }) => {
  return (
    <Fragment>
      {quiz && (
        <Card className="w-1/2">
          <CardHeader>
            <CardTitle>{quiz.name}</CardTitle>
          </CardHeader>
          <CardContent className="grid grid-cols-3 w-full">
            <div className="col-span-1 font-bold text-md">
              <p>Subjects</p>
              <p>Published at</p>
              <p>Start at</p>
              <p>End at</p>
              <p>Passing grade</p>
            </div>
            <div className="col-span-2 text-md">
              <p>{quiz.subject.join(", ")}</p>
              <p>{formatDateHour(new Date(quiz.publishedAt))}</p>
              <p>{formatDateHour(new Date(quiz.startAt))}</p>
              <p>{formatDateHour(new Date(quiz.endAt))}</p>
              <p>{quiz.passingGrade}</p>
            </div>
          </CardContent>
        </Card>
      )}
    </Fragment>
  );
};
