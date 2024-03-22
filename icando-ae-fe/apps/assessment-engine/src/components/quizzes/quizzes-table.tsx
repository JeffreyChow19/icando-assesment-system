import { useNavigate, useSearchParams } from "react-router-dom";
import { useState } from "react";
// import { useQuery } from '@tanstack/react-query';
// import { deleteStudent, getAllStudent } from '../../services/student.ts';
import { useConfirm } from "../../context/alert-dialog.tsx";
import { Button } from "@ui/components/ui/button.tsx";
import { Table, TableBody, TableCaption } from "@ui/components/ui/table.tsx";
import { SearchIcon } from "lucide-react";
import { Pagination } from "../pagination.tsx";
import { QuizCard } from "./quiz-card.tsx";
import { createQuiz } from "../../services/quiz.ts";
import { toast } from "@ui/components/ui/use-toast.ts";

export function QuizzesTable() {
  const [searchParams] = useSearchParams();
  const pageParams = searchParams.get("page");
  const navigate = useNavigate();

  const [page, setPage] = useState(pageParams ? parseInt(pageParams) : 1);

  // const { data, isLoading, refetch } = useQuery({
  //   queryKey: ['students', page],
  //   queryFn: () => getAllStudent({ page: page, limit: 10 }),
  // });

  // useEffect(() => {
  //   if (data) {
  //     if (data.meta.page != page) {
  //       setPage(data.meta.page);
  //     }
  //   }
  // }, [data, page]);

  const confirm = useConfirm();

  const data = [
    { id: 1, firstName: "John", lastName: "Doe", email: "" },
    { id: 2, firstName: "Jane", lastName: "Doe", email: "" },
  ];

  return (
    <div className="w-full mb-2">
      <div className="w-full flex flex-row font-normal gap-x-3 justify-between">
        <div className="w-full flex flex-row font-normal gap-x-3"></div>
        <div>
          <Button
            size={"sm"}
            onClick={() => {
              createQuiz()
                .then((data) => {
                  navigate(`/quizzes/${data.id}/edit`);
                })
                .catch(() => {
                  toast({
                    variant: "destructive",
                    description: "Failed to create quiz",
                  });
                });
            }}
          >
            New Quiz
          </Button>
        </div>
      </div>
      <Table className="my-2">
        <TableCaption>
          {data ? (
            <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
              <SearchIcon className="w-10 h-10" />
              No quiz yet.
            </div>
          ) : (
            data && (
              <div className="flex w-full justify-end">
                <Pagination
                  page={page}
                  totalPage={5 || 1}
                  setPage={setPage}
                  withSearchParams={true}
                />
              </div>
            )
          )}
        </TableCaption>
        <TableBody className="space-y-2">
          {data &&
            data.length > 0 &&
            data.map((quiz) => {
              return <QuizCard />;
            })}
        </TableBody>
      </Table>
    </div>
  );
}
