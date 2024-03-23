import { useEffect, useState } from "react";
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
import {
  assignStudents,
  unAssignStudents,
} from "../../services/classes";
import { getAllStudent } from "../../services/student";
import { useQuery } from "@tanstack/react-query";
import { Class } from "src/interfaces/classes";

export const ParticipantsTable = ({
  classData,
  refresh,
}: {
  classData: Class;
  refresh: () => void;
}) => {
  const classId = classData.id;
  const [isEditing, setIsEditing] = useState(false);
  const [isAdding, setIsAdding] = useState(false);

  const toggleEditing = () => {
    setIsEditing(!isEditing);
  };

  const switchEdits = () => {
    setIsAdding(!isAdding);
  };

  const saveChanges = () => {
    if (classData) {
      if (unassignList.length > 0) {
        unAssignStudents(classId, {
          studentIds: unassignList,
        }).then(() => {
          refresh();
        });
      }

      if (assignList.length > 0) {
        assignStudents(classId, {
          studentIds: assignList,
        }).then(() => {
          refresh();
        });
      }
    }

    setUnassignList([]);
    setAssignList([]);

    setIsEditing(false);
    setIsAdding(false);
  };

  const [page, setPage] = useState(1);
  const pageSize = 10;

  const classStudents = classData?.students;

  const totalPage = classStudents ? Math.ceil(classStudents.length / 10) : 0;

  const [unassignList, setUnassignList] = useState<string[]>([]);
  const [assignList, setAssignList] = useState<string[]>([]);

  // used for adding new students
  const [addPage, setAddPage] = useState(1);
  const { data: studentData } = useQuery({
    queryKey: ["students", addPage],
    queryFn: () => getAllStudent({ page: addPage, limit: 10 }),
  });
  useEffect(() => {
    if (studentData) {
      if (studentData.meta.page != addPage) {
        setAddPage(studentData.meta.page);
      }
    }
  }, [studentData, addPage]);

  useEffect(() => {
    console.log(unassignList);
    console.log(assignList);
  });

  return (
    <>
      <div className="w-full mb-2">
        <div className="w-full flex flex-row font-normal gap-x-3 justify-between">
          <div className="w-full flex flex-row font-normal gap-x-3"></div>
          <div className="flex flex-row gap-x-3 justify-between">
            {isEditing ? (
              <>
                <div className="flex flex-row gap-x-3 justify-between">
                  <Button
                    disabled={!isAdding}
                    size={"sm"}
                    variant={"default"}
                    onClick={switchEdits}
                  >
                    Remove Students
                  </Button>
                  <Button
                    disabled={isAdding}
                    size={"sm"}
                    variant={"default"}
                    onClick={switchEdits}
                  >
                    Add Students
                  </Button>
                </div>
                <div className="flex flex-row gap-x-3 justify-between">
                  <Button
                    size={"sm"}
                    variant={"outline"}
                    onClick={toggleEditing}
                  >
                    Cancel
                  </Button>
                  <Button
                    size={"sm"}
                    variant={"destructive"}
                    onClick={saveChanges}
                  >
                    Save Changes
                  </Button>
                </div>
              </>
            ) : (
              <Button size={"sm"} onClick={toggleEditing}>
                Edit Participants
              </Button>
            )}
          </div>
        </div>
        {/* EXISTING CLASS PARTICIPANTS */}
        {isAdding === false && (
          <Table>
            <TableCaption>
              {classStudents && classStudents.length === 0 ? (
                <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
                  <SearchIcon className="w-10 h-10" />
                  No participants.
                </div>
              ) : (
                classStudents &&
                classStudents.length / 10 >= 1 && (
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
                <TableHead>NISN</TableHead>
                {isEditing === true && <TableHead>Actions</TableHead>}
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
                          {student.firstName} {student.lastName}
                        </TableCell>
                        <TableCell>{student.nisn}</TableCell>

                        {isEditing === true && (
                          <TableCell>
                            <div className="flex space-x-2">
                              <>
                                {unassignList.includes(student.id) ? (
                                  <Button
                                    size={"sm"}
                                    variant={"outline"}
                                    onClick={() => {
                                      setUnassignList(
                                        unassignList.filter(
                                          (item) => item !== student.id,
                                        ),
                                      );
                                    }}
                                  >
                                    Cancel
                                  </Button>
                                ) : (
                                  <Button
                                    size={"sm"}
                                    variant={"destructive"}
                                    onClick={() => {
                                      setUnassignList([
                                        ...unassignList,
                                        student.id,
                                      ]);
                                    }}
                                  >
                                    Remove
                                  </Button>
                                )}
                              </>
                            </div>
                          </TableCell>
                        )}
                      </TableRow>
                    );
                  })}
            </TableBody>
          </Table>
        )}

        {/* ADD CLASS PARTICIPANTS */}
        {isEditing === true && isAdding == true && (
          <Table>
            <TableCaption>
              {studentData && studentData.meta.totalItem === 0 ? (
                <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
                  <SearchIcon className="w-10 h-10" />
                  No students available.
                </div>
              ) : (
                studentData &&
                studentData.meta.totalPage > 1 && (
                  <div className="flex w-full justify-end">
                    <Pagination
                      page={addPage}
                      totalPage={studentData?.meta.totalPage || 1}
                      setPage={setAddPage}
                      withSearchParams={true}
                    />
                  </div>
                )
              )}
            </TableCaption>
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>NISN</TableHead>
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {studentData &&
                studentData.data &&
                studentData.data.length > 0 &&
                studentData.data.map((student) => {
                  return (
                    <TableRow key={student.id}>
                      <TableCell>
                        {student.firstName} {student.lastName}
                      </TableCell>
                      <TableCell>{student.nisn}</TableCell>

                      {isEditing === true &&
                      classStudents?.some((s) => s.id === student.id) ? (
                        <TableCell>Student is already participant</TableCell>
                      ) : (
                        <TableCell>
                          <div className="flex space-x-2">
                            <>
                              {assignList.includes(student.id) ? (
                                <Button
                                  size={"sm"}
                                  variant={"outline"}
                                  onClick={() => {
                                    setAssignList(
                                      assignList.filter(
                                        (item) => item !== student.id,
                                      ),
                                    );
                                  }}
                                >
                                  X
                                </Button>
                              ) : (
                                <Button
                                  size={"sm"}
                                  variant={"secondary"}
                                  onClick={() => {
                                    setAssignList([...assignList, student.id]);
                                  }}
                                >
                                  +
                                </Button>
                              )}
                            </>
                          </div>
                        </TableCell>
                      )}
                    </TableRow>
                  );
                })}
            </TableBody>
          </Table>
        )}
      </div>
    </>
  );
};
