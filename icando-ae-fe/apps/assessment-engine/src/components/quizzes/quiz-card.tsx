import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@ui/components/ui/card.tsx";
import { QuizDetail } from "../../interfaces/quiz.ts";
import { Button } from "@ui/components/ui/button.tsx";
import { Link } from "react-router-dom";
import { Badge } from "@ui/components/ui/badge.tsx";
import { useState } from "react";
import { QuizHistory } from "./quiz-history.tsx";
import { formatDate, formatHour } from "../../utils/format-date.ts";

export function QuizCard({ quiz }: { quiz: QuizDetail }) {
  const [isHistoryOpen, setHistoryOpen] = useState(false);
  const handleOpenHistory = () => {
    setHistoryOpen(true);
  };

  const handleCloseHistory = () => {
    setHistoryOpen(false);
  };

  return (
    <Card className="space-x-2">
      <CardHeader className="flex flex-row justify-between">
        <CardTitle>{quiz.name ? quiz.name : "Untitled Quiz"}</CardTitle>
        {quiz.lastPublishedAt ? (
          <Badge key={"Published"} variant={"green"}>Published</Badge>
        ) : (
          <Badge key={"Draft"} variant={"destructive"}>Draft</Badge>
        )}
      </CardHeader>
      <CardContent>
        <CardDescription>
          {quiz.subject && quiz.subject.length > 0 && (
            <div className="flex flex-wrap gap-2">
              {quiz.subject.map((subject) => (
                <Badge key={subject}>{subject}</Badge>
              ))}
            </div>
          )}
        </CardDescription>

      </CardContent>
      <CardContent>
        <div className="flex flex-row justify-between">
          <div>
            <p className="mb-2">
              Updated By:{" "}
              {quiz.updatedBy ? <Badge key={quiz.updatedBy} variant={"secondary"}>{quiz.updatedBy}</Badge>
                : "-"}
            </p>
            <p>
              Last Published at: {quiz.lastPublishedAt ?
                <>
                  <Badge key={formatDate(new Date(quiz.lastPublishedAt))} className="mr-2" variant={"outline"}>{formatDate(new Date(quiz.lastPublishedAt))}</Badge>
                  <Badge key={formatHour(new Date(quiz.lastPublishedAt))} variant={"outline"}>{formatHour(new Date(quiz.lastPublishedAt))}</Badge>
                </>
                : "-"}
            </p>
          </div>
        </div>
      </CardContent>
      {quiz.lastPublishedAt ? (
        <CardFooter className="flex justify-between">
          <div onClick={handleOpenHistory} className="cursor-pointer underline text-gray-500 hover:text-gray-400">
            Check Version History
          </div>

          {isHistoryOpen && (
            <QuizHistory quizId={quiz.id} quizName={quiz.name} onClose={handleCloseHistory} />
          )}
          <div className="flex flex-row justify-between space-x-2">
            <Button variant="outline">
              <Link to={`/quiz/${quiz.id}/edit`}>Edit</Link>
            </Button>
            <Button>
              <Link to={`/quiz/${quiz.id}/publish`}>Publish</Link>
            </Button>
          </div>
        </CardFooter>
      ) : (
        <CardFooter className="flex justify-end">
          <div className="flex flex-row justify-between space-x-2">
            <Button variant="outline">
              <Link to={`/quiz/${quiz.id}/edit`}>Edit</Link>
            </Button>
            <Button>
              <Link to={`/quiz/${quiz.id}/publish`}>Publish</Link>
            </Button>
          </div>
        </CardFooter>
      )}
    </Card>
  );
}
