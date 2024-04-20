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
// import { useQuery } from "@tanstack/react-query";

export function QuizCard({ quiz }: { quiz: QuizDetail }) {
  const [isHistoryOpen, setHistoryOpen] = useState(false);

  function formatDate(date: Date): string {
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0'); // getMonth() is zero-indexed
    const year = date.getFullYear().toString().slice(-2);

    return `${day}-${month}-${year}`;
  }
  function formatHour(date: Date): string {
    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');

    return `${hours}:${minutes}`;
  }

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
