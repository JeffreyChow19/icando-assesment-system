import { SearchIcon } from "lucide-react";
import { Pagination } from "../pagination";
import { CustomCard } from "../ui/custom-card";
import { useQuery } from "@tanstack/react-query";
import { getAllQuiz } from "../../services/quiz";
import { useState } from "react";
import {
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@ui/components/ui/card";
import { Badge } from "@ui/components/ui/badge";

export const QuizzesTable = () => {
  const [page, setPage] = useState(1);

  // todo: add query param to query and queryKey
  const { data, isLoading } = useQuery({
    queryFn: () => getAllQuiz({ page: page, limit: 10 }),
    queryKey: ["quizzes", page],
  });

  return (
    // todo: search bar and filters
    <div className="w-full mb-2">
      <div className="w-full flex flex-row font-normal gap-x-3 justify-between">
        <div className="w-full flex flex-row font-normal gap-x-3"></div>
        <div>
          <p>search</p>
        </div>
      </div>

      <div className="grid grid-cols-1 gap-2">
        {data &&
          data.data.length > 0 &&
          data.data.map((quiz) => {
            return (
              // todo: quiz detail card and history dropdown
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
                <CardContent>Check Version History</CardContent>
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
          <div className="flex w-full justify-end my-10">
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
