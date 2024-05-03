import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@ui/components/ui/card.tsx";
import { useQuery } from "@tanstack/react-query";
import { useNavigate, useSearchParams } from "react-router-dom";
import { useDebounce } from "use-debounce";
import { useEffect, useState } from "react";
import { getStudentQuizzes } from "../../services/student-quiz.ts";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@ui/components/ui/select.tsx";
import { Quiz } from "../../interfaces/quiz.ts";
import {
  ToggleGroup,
  ToggleGroupItem,
} from "@ui/components/ui/toggle-group.tsx";
import { Input } from "@ui/components/ui/input.tsx";
import {
  CheckIcon,
  ListCollapse,
  SearchIcon,
  ViewIcon,
  XIcon,
} from "lucide-react";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@ui/components/ui/table.tsx";
import { Button } from "@ui/components/ui/button.tsx";
import { Pagination } from "../pagination.tsx";
import { Badge } from "@ui/components/ui/badge.tsx";
import { formatDateHour } from "../../utils/format-date.ts";
export const StudentQuizTable = ({ quiz }: { quiz: Quiz }) => {
  const [searchParams, setSearchParams] = useSearchParams();
  const [debouncedQuery] = useDebounce(searchParams, 300);
  const [studentName, setStudentNameQuery] = useState<string>(
    searchParams.get("studentName") || "",
  );
  const [classId, setClassId] = useState<string>(
    searchParams.get("classId") || "ALL",
  );
  const [page, setPage] = useState<number>(1);
  const [quizStatus, setQuizStatus] = useState<string>(
    searchParams.get("quizStatus") || "ALL",
  );
  const navigate = useNavigate();

  useEffect(() => {
    const newUrlSearchParams: URLSearchParams = new URLSearchParams();
    if (classId != "" && classId != "ALL") {
      newUrlSearchParams.set("classId", classId);
    }
    if (studentName != "") {
      newUrlSearchParams.set("studentName", studentName);
    }
    if (quizStatus != "" && quizStatus != "ALL") {
      newUrlSearchParams.set("quizStatus", quizStatus);
    }
    newUrlSearchParams.set("page", page.toString());
    setSearchParams(newUrlSearchParams);
  }, [classId, page, studentName, quizStatus]);

  const { data, isLoading } = useQuery({
    queryKey: ["student-quiz", ...debouncedQuery],
    queryFn: () =>
      getStudentQuizzes({
        page,
        limit: 20,
        ...(studentName != "" ? { studentName } : {}),
        ...(classId != "ALL" ? { classId } : {}),
        ...(quizStatus != "ALL" ? { quizStatus } : {}),
        quizId: quiz.id,
      }),
  });

  useEffect(() => {
    console.log(data);
  }, [data]);

  return (
    <Card>
      <CardHeader>
        <CardTitle>Assigned Students</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="flex flex-col gap-y-2">
          <div className="flex justify-between">
            <ToggleGroup
              type="single"
              onValueChange={(value) => setQuizStatus(value)}
              defaultValue={searchParams.get("quizStatus") || "ALL"}
            >
              <ToggleGroupItem value="ALL" aria-label="Toggle All">
                All
              </ToggleGroupItem>
              <ToggleGroupItem value="SUBMITTED" aria-label="Toggle Submitted">
                Submitted
              </ToggleGroupItem>
              <ToggleGroupItem value="STARTED" aria-label="Toggle Started">
                Started
              </ToggleGroupItem>
              <ToggleGroupItem
                value="NOT_STARTED"
                aria-label="Toggle Not Started"
              >
                Not Started
              </ToggleGroupItem>
            </ToggleGroup>
            <div className="flex gap-x-2 items-center">
              <Select onValueChange={(value) => setClassId(value)}>
                <SelectTrigger className="w-[200px]">
                  <SelectValue placeholder="Select class..." />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    <SelectLabel>Class</SelectLabel>
                    {quiz.classes?.map((c) => (
                      <SelectItem key={c.id} value={c.id}>
                        {c.name} - {c.grade}
                      </SelectItem>
                    ))}
                  </SelectGroup>
                </SelectContent>
              </Select>
              <div className="relative h-10 w-[300px]">
                <SearchIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-500 z-10" />
                <Input
                  type="text"
                  placeholder="Search student..."
                  className="pl-10 pr-3 py-2 text-md"
                  defaultValue={searchParams.get("studentName") || ""}
                  onChange={(e) => setStudentNameQuery(e.target.value)}
                />
              </div>
            </div>
          </div>
          {data && (
            <Table>
              <TableCaption>
                {data.meta && data.meta.totalItem === 0 ? (
                  <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
                    <SearchIcon className="w-10 h-10" />
                    No assigned students yet.
                  </div>
                ) : (
                  !isLoading &&
                  data.meta.totalPage > 1 && (
                    <div className="flex w-full justify-end">
                      <Pagination
                        page={page}
                        totalPage={data?.meta.totalPage || 1}
                        setPage={setPage}
                        withSearchParams={true}
                      />
                    </div>
                  )
                )}
              </TableCaption>
              <TableHeader>
                <TableRow>
                  <TableHead>Name</TableHead>
                  <TableHead className="text-center">Status</TableHead>
                  <TableHead className="text-right">Correct Count</TableHead>
                  <TableHead className="text-right">Total Score</TableHead>
                  <TableHead>Submitted At</TableHead>
                  <TableHead className="text-center">Passed</TableHead>
                  <TableHead className="text-right">Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {data &&
                  data.data.length > 0 &&
                  data.data.map((studentQuiz) => {
                    return (
                      <TableRow key={studentQuiz.id}>
                        <TableCell>{studentQuiz.name}</TableCell>
                        <TableCell className="text-center">
                          <Badge
                            variant={
                              studentQuiz.status == "SUBMITTED"
                                ? "default"
                                : studentQuiz.status == "NOT_STARTED"
                                  ? "outline"
                                  : "secondary"
                            }
                          >
                            {studentQuiz.status.replace("_", " ")}
                          </Badge>
                        </TableCell>
                        <TableCell className="text-right">
                          {studentQuiz.correctCount}/{quiz.questions?.length}
                        </TableCell>
                        <TableCell className="text-right">
                          {studentQuiz.totalScore}
                        </TableCell>
                        <TableCell>
                          {studentQuiz.completedAt
                            ? formatDateHour(new Date(studentQuiz.completedAt))
                            : "-"}
                        </TableCell>
                        <TableCell className="text-center flex justify-center">
                          {!!studentQuiz.completedAt &&
                          studentQuiz.totalScore >= quiz.passingGrade ? (
                            <CheckIcon className="rounded-full size-7 text-green-500 bg-green-200 p-1" />
                          ) : studentQuiz.completedAt ? (
                            <XIcon className="rounded-full size-7 text-rose-500 bg-rose-200 p-1" />
                          ) : (
                            "-"
                          )}
                        </TableCell>
                        <TableCell className="text-right">
                          <Button
                            className="h-fit py-1"
                            disabled={studentQuiz.status !== "SUBMITTED"}
                            onClick={() =>
                              navigate(
                                `/quiz/${quiz.id}/review/${studentQuiz.id}`,
                              )
                            }
                          >
                            Review
                          </Button>
                        </TableCell>
                      </TableRow>
                    );
                  })}
              </TableBody>
            </Table>
          )}
        </div>
      </CardContent>
    </Card>
  );
};
