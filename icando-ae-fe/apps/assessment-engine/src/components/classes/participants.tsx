import { useState } from "react";
import { Button } from "@ui/components/ui/button";
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
import { Pagination } from "../pagination";
import { Student } from "../../interfaces/student";
import { unAssignStudents } from "../../services/classes";
import { Class } from "../../interfaces/classes";
import { Checkbox } from "@ui/components/ui/checkbox";
import { useConfirm } from "../../context/alert-dialog";
import { AddParticipantsDialog } from "./participants-add";

export const ParticipantsTable = ({
  classData,
  refresh,
}: {
  classData: Class;
  refresh: () => void;
}) => {
  const classId = classData.id;
  const [isEditing, setIsEditing] = useState(false);

  const onCheckedStudents = (
    checked: boolean | "indeterminate",
    student: Student,
  ) => {
    if (checked === "indeterminate") return;
    if (checked) {
      if (unassignList.length == 0) setIsEditing(true);
      setUnassignList([...unassignList, student]);
    } else {
      if (unassignList.length == 1) setIsEditing(false);
      setUnassignList(unassignList.filter((item) => item.id !== student.id));
    }
  };

  const deleteStudents = () => {
    if (classData) {
      if (unassignList.length > 0) {
        unAssignStudents(classId, {
          studentIds: unassignList.map((item) => item.id),
        }).then(() => {
          setPage(1);
          refresh();
        });
      }
    }
    setUnassignList([]);

    setIsEditing(false);
  };
  const confirm = useConfirm();

  const [page, setPage] = useState(1);
  const pageSize = 10;

  const classStudents = classData?.students;

  const totalPage = classStudents
    ? Math.ceil(classStudents.length / pageSize)
    : 0;

  const [unassignList, setUnassignList] = useState<Student[]>([]);

  // useEffect(() => {
  //   console.log(unassignList);
  //   console.log(isEditing);
  // });

  return (
    <>
      <div className="w-full mb-2">
        <div className="w-full flex flex-row font-normal gap-x-3 justify-between">
          {isEditing == true && (
            <div className="flex flex-row gap-x-3 justify-between"></div>
          )}

          <div className="w-full flex flex-row font-normal gap-x-3"></div>
          <div className="flex flex-row gap-x-3 justify-between">
            {isEditing == true && (
              <Button
                variant={"destructive"}
                onClick={() => {
                  confirm({
                    title: "Remove Student(s) from Class",
                    body: "Are you sure to remove these students from this class?",
                  }).then((confirm) => {
                    if (confirm) {
                      deleteStudents();
                    }
                  });
                }}
              >
                Remove Students
              </Button>
            )}
            <AddParticipantsDialog classId={classData.id} onSuccess={refresh} />
          </div>
        </div>

        <Table>
          <TableCaption>
            {(classStudents && classStudents.length == 0) || !classStudents ? (
              <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
                <SearchIcon className="w-10 h-10" />
                No students.
              </div>
            ) : (
              classStudents &&
              classStudents.length / pageSize >= 1 && (
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
              <TableHead className="w-[4vw]">Select</TableHead>
              <TableHead>Name</TableHead>
              <TableHead>NISN</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {classStudents &&
              classStudents.length > 0 &&
              classStudents
                .slice((page - 1) * pageSize, page * pageSize)
                .map((student: Student) => {
                  return (
                    <TableRow key={student.id}>
                      <TableCell>
                        <Checkbox
                          onCheckedChange={(checked) =>
                            onCheckedStudents(checked, student)
                          }
                          checked={
                            unassignList.findIndex(
                              (item) => item.id === student.id,
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
      </div>
    </>
  );
};
