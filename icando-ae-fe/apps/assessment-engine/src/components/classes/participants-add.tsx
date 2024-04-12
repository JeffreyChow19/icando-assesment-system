import {
  Dialog,
  DialogContent,
  DialogTrigger,
  DialogHeader,
  DialogTitle,
} from "@ui/components/ui/dialog.tsx";
import { Button } from "@ui/components/ui/button.tsx";
import { useEffect, useState } from "react";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@ui/components/ui/table";
import { SearchIcon } from "lucide-react";
import { Checkbox } from "@ui/components/ui/checkbox";
import { Student } from "../../interfaces/student";
import { Pagination } from "../pagination";
import { getAllStudent } from "../../services/student";
import { useQuery } from "@tanstack/react-query";
import { assignStudents } from "../../services/classes";

interface AddParticipantsProps {
  classId: string;
  onSuccess: () => void;
}

export const AddParticipantsDialog = ({
  classId,
  onSuccess,
}: AddParticipantsProps) => {
  const [open, setOpen] = useState<boolean>(false);
  const [page, setPage] = useState(1);
  const [assignList, setAssignList] = useState<string[]>([]);

  const addStudents = () => {
    if (assignList.length > 0) {
      assignStudents(classId, {
        studentIds: assignList,
      }).then(() => {
        onSuccess();
      });
    }
    setAssignList([]);
  };

  const onCheckedStudents = (
    checked: boolean | "indeterminate",
    student: Student,
  ) => {
    if (checked === "indeterminate") return;
    if (checked) {
      setAssignList([...assignList, student.id]);
    } else {
      setAssignList(assignList.filter((item) => item !== student.id));
    }
  };
  const pageSize = 8;

  // used for adding new students
  const { data: studentData } = useQuery({
    queryKey: ["students", page],
    queryFn: () => getAllStudent({ page: page, limit: pageSize }),
  });
  useEffect(() => {
    if (studentData) {
      if (studentData.meta.page != page) {
        setPage(studentData.meta.page);
      }
    }
  }, [studentData, page]);

  // useEffect(() => {
  //   console.log(assignList);
  // });

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button>Add Participants</Button>
      </DialogTrigger>
      <DialogContent className="min-w-[70vw] min-h-[40vh]">
        <DialogHeader>
          <DialogTitle>Add Participants</DialogTitle>
        </DialogHeader>
        <Table>
          <TableCaption>
            {studentData && studentData.data.length === 0 ? (
              <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
                <SearchIcon className="w-10 h-10" />
                No participants.
              </div>
            ) : (
              studentData &&
              studentData.meta.totalPage > 1 && (
                <div className="flex w-full justify-end">
                  <Button
                    size={"sm"}
                    disabled={assignList.length == 0}
                    onClick={() => {
                      addStudents();
                      setOpen(false);
                    }}
                  >
                    Add Students
                  </Button>
                  <Pagination
                    page={page}
                    totalPage={studentData.meta.totalPage || 1}
                    setPage={setPage}
                    withSearchParams={true}
                  />
                </div>
              )
            )}
          </TableCaption>
          <TableHeader>
            <TableRow>
              <TableHead>Select</TableHead>
              <TableHead>Name</TableHead>
              <TableHead>NISN</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {studentData &&
              studentData.data.length > 0 &&
              studentData.data
                .filter((student) => student.classId !== classId)
                .map((student: Student) => {
                  return (
                    <TableRow key={student.id}>
                      <TableCell>
                        <Checkbox
                          onCheckedChange={(checked) =>
                            onCheckedStudents(checked, student)
                          }
                          checked={
                            assignList.findIndex(
                              (item) => item === student.id,
                            ) != -1
                          }
                        />
                      </TableCell>
                      <TableCell>
                        {student.firstName} {student.lastName}
                      </TableCell>
                      <TableCell>{student.nisn}</TableCell>
                    </TableRow>
                  );
                })}
          </TableBody>
        </Table>
      </DialogContent>
    </Dialog>
  );
};
