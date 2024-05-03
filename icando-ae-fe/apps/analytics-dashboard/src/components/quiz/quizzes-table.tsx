import { SearchIcon } from "lucide-react";
import { Pagination } from "../pagination";
import { CustomCard } from "../ui/custom-card";
import { useQuery } from "@tanstack/react-query";
import { getAllQuiz } from "../../services/quiz";
import { useEffect, useState } from "react";
import {
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@ui/components/ui/card";
import { Badge } from "@ui/components/ui/badge";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@ui/components/ui/select";
import { Input } from "@ui/components/ui/input";
import { useSearchParams } from "react-router-dom";
import { SUBJECTS } from "../../utils/constants";
import { useDebounce } from "use-debounce";
import { HistoryCollapsible } from "./history-collapsible";
import { formatDate, formatHour } from "../../utils/format-date";

export const QuizzesTable = () => {
  const [page, setPage] = useState(1);

  // search filter params
  const [searchParams, setSearchParams] = useSearchParams();
  const quizName = searchParams.get("quiz");
  const [quizNameFilter, setQuizNameFilter] = useState<string | undefined>(
    quizName ?? undefined,
  );

  const subjectOptions = SUBJECTS;
  const subjectQuery = searchParams.get("subject");
  const [subjectFilter, setSubjectFilter] = useState<string>(
    subjectQuery ?? "all",
  );
  const [query] = useDebounce([quizNameFilter, subjectFilter], 300);

  useEffect(() => {
    if (quizNameFilter && quizNameFilter !== "") {
      searchParams.set("name", quizNameFilter);
    } else {
      searchParams.delete("name");
    }
    if (subjectFilter === "all") {
      searchParams.delete("subject");
    } else {
      searchParams.set("subject", subjectFilter);
    }
    setSearchParams(searchParams);
  }, [quizNameFilter, subjectFilter, searchParams, setSearchParams]);

  // quiz data query
  const { data, isLoading } = useQuery({
    queryFn: () =>
      getAllQuiz({
        page: page,
        limit: 10,
        name: quizNameFilter,
        subject: subjectFilter === "all" ? undefined : subjectFilter,
      }),
    queryKey: ["quizzes", page, ...query],
  });

  return (
    <div className="w-full mb-2">
      <div className="w-full flex flex-row font-normal gap-x-3 justify-between">
        <div className="w-full flex flex-row font-normal gap-x-3">
          <Input
            className={"w-[360px]"}
            placeholder={"Search name ..."}
            value={quizNameFilter}
            onChange={(e) => setQuizNameFilter(e.target.value)}
          />
          <Select value={subjectFilter} onValueChange={setSubjectFilter}>
            <SelectTrigger className="w-[240px]">
              <SelectValue placeholder="Select Subject" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="all" key="all">
                Select Subjects
              </SelectItem>
              {subjectOptions &&
                subjectOptions.map((subject) => (
                  <SelectItem value={subject} key={subject}>
                    {subject}
                  </SelectItem>
                ))}
            </SelectContent>
          </Select>
        </div>
      </div>

      <div className="grid grid-cols-1 gap-2 my-5">
        {data &&
          data.data.length > 0 &&
          data.data.map((quiz) => {
            return (
              <CustomCard key={quiz.id}>
                <CardHeader>
                  <CardTitle>{quiz.name}</CardTitle>
                  <CardDescription>
                    <div className="flex flex-wrap gap-2">
                      {quiz.subject.map((subject) => (
                        <Badge key={subject}>{subject}</Badge>
                      ))}
                    </div>
                  </CardDescription>
                </CardHeader>
                <CardContent>
                  <p>
                    <strong>Passing Grade:</strong> {quiz.passingGrade}
                    <br />
                    <strong>Created By:</strong> {quiz.createdBy}
                    <br />
                    <strong>Latest Release:</strong>{" "}
                    {formatDate(new Date(quiz.lastPublishedAt!))} -{" "}
                    {formatHour(new Date(quiz.lastPublishedAt!))}
                    <br />
                  </p>
                  <br />
                  <HistoryCollapsible quizId={quiz.id} />
                </CardContent>
              </CustomCard>
            );
          })}
      </div>

      {data && data.meta.totalItem === 0 ? (
        <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
          <SearchIcon className="w-10 h-10" />
          No quiz available.
        </div>
      ) : (
        !isLoading &&
        data &&
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
    </div>
  );
};
