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

export function QuizCard({ quiz }: { quiz: QuizDetail }) {
  return (
    <Card className="space-x-2">
      <CardHeader className="flex flex-row justify-between">
        <CardTitle>{quiz.name ? quiz.name : "Untitled Quiz"}</CardTitle>
        {quiz.publishedAt ? (
          <p className="text-green-500">Published</p>
        ) : (
          <p className="text-red-500">Draft</p>
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
      <CardFooter className="flex justify-between">
        <div className="flex flex-column justify-between">
          <div>
            <p>
              Created By:{" "}
              {quiz.creator
                ? `${quiz.creator.firstName} ${quiz.creator.lastName}`
                : "-"}
            </p>
            <p>
              Last Published at: {quiz.publishedAt ? quiz.publishedAt : "-"}
            </p>
          </div>
        </div>
        <div className="flex flex-row justify-between space-x-2">
          <Button variant="outline">
            <Link to={`/quiz/${quiz.id}/edit`}>Edit</Link>
          </Button>
          <Button>
            <Link to={`/quiz/${quiz.id}/publish`}>Publish</Link>
          </Button>
        </div>
      </CardFooter>
    </Card>
  );
}
