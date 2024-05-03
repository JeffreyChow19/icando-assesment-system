import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@ui/components/ui/table.tsx";

import { SearchIcon } from "lucide-react";

import { StudentQuiz } from "../../interfaces/student.ts";
import dayjs from "dayjs";
import { useNavigate } from "react-router-dom";

interface QuizListProps {
  quizzes: StudentQuiz[];
}

export const QuizList = ({ quizzes }: QuizListProps) => {
  const navigate = useNavigate();

  const handleOnRowClicked = (quiz: StudentQuiz) => {
    navigate(`/quiz/${quiz.quizId}/review/${quiz.id}`);
  };

  return (
    <Table>
      <TableCaption>
        {quizzes.length <= 0 && (
          <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
            <SearchIcon className="w-10 h-10" />
            No quizzes submitted
          </div>
        )}
      </TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="w-[4vw]">#</TableHead>
          <TableHead className="text-center">Quiz</TableHead>
          <TableHead className="text-center">Submitted At</TableHead>
          <TableHead className="text-center">Score</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {quizzes.map((quiz, index) => (
          <TableRow
            key={quiz.id}
            onClick={() => handleOnRowClicked(quiz)}
            className="hover:bg-gray-200 hover:cursor-pointer"
          >
            <TableCell>{index + 1}</TableCell>
            <TableCell>{quiz.name}</TableCell>
            <TableCell className="text-center">
              {dayjs(quiz.completedAt).format("L LT")}
            </TableCell>
            <TableCell className="text-center">{quiz.totalScore}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};
