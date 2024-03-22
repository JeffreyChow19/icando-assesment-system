import { Link } from "react-router-dom";
import { useState } from "react";
import { useQuery } from "@tanstack/react-query";

import { useConfirm } from "../../context/alert-dialog.tsx";
import { Button } from "@ui/components/ui/button.tsx";
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
import { Pagination } from "../pagination.tsx";
import { deleteClass, getAllClasses } from "../../services/classes.ts";

export function ClassesTable() {
  const [page, setPage] = useState(1);
  const pageSize = 10;

  const { data, isLoading, refetch } = useQuery({
    queryKey: ["classes"],
    queryFn: () => getAllClasses(),
  });

  const totalPage = data ? Math.ceil(data.data.length / 10) : 0;

  const confirm = useConfirm();

  return (
    <div className="w-full mb-2">
      <div className="w-full flex flex-row font-normal gap-x-3 justify-between">
        <div className="w-full flex flex-row font-normal gap-x-3"></div>
        <div>
          <Button size={"sm"}>
            <Link to={"/classes/new"}>New Class</Link>
          </Button>
        </div>
      </div>
      <Table>
        <TableCaption>
          {data && data.data.length === 0 ? (
            <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
              <SearchIcon className="w-10 h-10" />
              No classes available.
            </div>
          ) : (
            !isLoading &&
            data &&
            data.data.length / 10 >= 1 && (
              <div className="flex w-full justify-end">
                <Pagination
                  page={page}
                  totalPage={totalPage}
                  setPage={setPage}
                />
              </div>
            )
          )}
        </TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead>Name</TableHead>
            <TableHead>Grade</TableHead>
            <TableHead>Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {data &&
            data.data.length > 0 &&
            data.data
              .slice((page - 1) * pageSize, page * pageSize)
              .map((classes) => {
                return (
                  <TableRow key={classes.id}>
                    <TableCell>{classes.name}</TableCell>
                    <TableCell>{classes.grade}</TableCell>
                    {/* todo: display teacher correctly (not by id?) */}
                    {/* todo: display participant (student count?) */}
                    <TableCell>
                      <div className="flex space-x-2">
                        <>
                          <Button size={"sm"}>
                            <Link to={`/classes/edit/${classes.id}`}>
                              Edit Class
                            </Link>
                          </Button>
                          <Button size={"sm"}>
                            <Link to={`#`}>Edit Participants</Link>
                          </Button>
                          <Button
                            size={"sm"}
                            variant={"destructive"}
                            onClick={() => {
                              confirm({
                                title: "Are you sure?",
                                body: "Are you sure want to delete this student?",
                              }).then((result) => {
                                if (result) {
                                  deleteClass(classes.id).then(() => {
                                    refetch();
                                  });
                                }
                              });
                            }}
                          >
                            Delete
                          </Button>
                        </>
                      </div>
                    </TableCell>
                  </TableRow>
                );
              })}
        </TableBody>
      </Table>
    </div>
  );
}
