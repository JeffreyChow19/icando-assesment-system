import { useQuery } from "@tanstack/react-query";
import { Button } from "@ui/components/ui/button";
import { useState } from "react";
import { getClass } from "../../services/classes";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@ui/components/ui/dialog";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@ui/components/ui/table";

interface ClassTeachersCollapsibleProps {
  classId: string;
}

export const ClassTeachersDialog = ({
  classId,
}: ClassTeachersCollapsibleProps) => {
  const [open, setOpen] = useState(false);

  const { data, isLoading, refetch } = useQuery({
    enabled: false,
    queryKey: ["classes", classId],
    queryFn: () =>
      getClass({ withTeacher: true, withStudents: false }, classId!),
  });

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>{data?.data.name} Teachers</DialogTitle>
        </DialogHeader>
        {open &&
          !isLoading &&
          (data &&
          data.data &&
          data.data.teachers &&
          data.data.teachers.length > 0 ? (
            <Table>
              <TableHeader>
                <TableHead>Name</TableHead>
                <TableHead>Email</TableHead>
              </TableHeader>
              <TableBody>
                {data.data.teachers.map((item) => {
                  return (
                    <TableRow>
                      <TableCell>
                        {item.firstName} {item.lastName}
                      </TableCell>
                      <TableCell>{item.email}</TableCell>
                    </TableRow>
                  );
                })}
              </TableBody>
            </Table>
          ) : (
            "No teachers assigned to this class"
          ))}
      </DialogContent>
      <DialogTrigger asChild>
        <Button
          size="sm"
          onClick={() => {
            if (!open) {
              if (!isLoading) {
                refetch();
              }
              setOpen(!open);
            }
          }}
        >
          Show Teachers
        </Button>
      </DialogTrigger>
    </Dialog>
  );
};
