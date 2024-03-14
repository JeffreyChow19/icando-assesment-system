import {
  Table,
  TableBody,
  TableCaption, TableCell,
  TableHead,
  TableHeader,
  TableRow
} from "@ui/components/ui/table.tsx";
import {Question} from "../../../interfaces/question.ts";
import {useState} from "react";
import {SearchIcon} from "lucide-react";
import {Pagination} from "../../pagination.tsx";
import {QuestionForm} from "./question-form.tsx";

interface QuestionListProps {
  questions: Question[]
}
export const QuestionList = ({questions} : QuestionListProps) => {
  const [page, setPage] = useState<number>(1);
  const questionsPerPage = 10;
  const startIndex = (page - 1) * questionsPerPage;
  const endIndex = startIndex + questionsPerPage;
  const paginatedQuestions = questions.slice(startIndex, endIndex);
  const totalPage = Math.ceil(questions.length / questionsPerPage);
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
          <TableRow key={question.id}>
            <TableCell>{(page - 1)*questionsPerPage + index + 1}</TableCell>
            <TableCell>{question.text}</TableCell>
            <TableCell className="text-right">
              <QuestionForm type="edit" />
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  )
}
