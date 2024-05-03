import { Link, useSearchParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { getAllClasses, getAllStudent } from "../../services/student.ts";
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
import { useDebounce } from "use-debounce";
import { Input } from "@ui/components/ui/input.tsx";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@ui/components/ui/select.tsx";
import { Class } from "../../interfaces/class.ts";
import { Badge } from "@ui/components/ui/badge.tsx";
import { Button } from "@ui/components/ui/button.tsx";

interface ClassMap {
  [id: string]: string;
}

export function StudentsTable() {
  const [searchParams, setSearchParams] = useSearchParams();
  const pageParams = searchParams.get("page");
  const nameParams = searchParams.get("name");
  const classIdParams = searchParams.get("classId");

  const [page, setPage] = useState(pageParams ? parseInt(pageParams) : 1);
  const [name, setName] = useState(nameParams ? nameParams : "");
  const [classId, setClassId] = useState(classIdParams ? classIdParams : "");
  const [query] = useDebounce([name, classId], 300);
  const [classMap] = useState<ClassMap>({});

  useEffect(() => {
    setPage(1);
  }, [name, classId]);

  useEffect(() => {
    if (name != "") {
      searchParams.set("name", name);
    } else {
      searchParams.delete("name");
    }
    if (classId != "") {
      searchParams.set("classId", classId);
    } else {
      searchParams.delete("classId");
    }
    setSearchParams(searchParams);
  }, [name, classId]);

  const { data, isLoading } = useQuery({
    queryKey: ["students", page, ...query],
    queryFn: () =>
      getAllStudent({
        page: page,
        limit: 10,
        name,
        classId,
        orderBy: "first_name",
        asc: true,
      }),
  });

  const { data: classes } = useQuery({
    queryKey: ["classes"],
    queryFn: () => getAllClasses(),
  });

  useEffect(() => {
    if (classes) {
      classes.forEach((cls) => {
        classMap[cls.id] = `${cls.name} - ${cls.grade}`;
      });
    }
  }, [classes]);

  useEffect(() => {
    if (data) {
      if (data.meta.page != page) {
        setPage(data.meta.page);
      }
    }
  }, [data, page]);

  return (
    <div className="w-full mb-2">
      <div className="w-full flex flex-row font-normal gap-x-3 justify-between">
        <div className="w-full flex flex-row font-normal gap-x-3">
          <Input
            className={"w-[360px]"}
            placeholder={"Search name ..."}
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <Select value={classId} onValueChange={setClassId}>
            <SelectTrigger className="w-[240px]">
              <SelectValue placeholder="Select class" />
            </SelectTrigger>
            <SelectContent>
              {classes &&
                classes.map((e: Class) => (
                  <SelectItem value={e.id} key={e.id}>
                    {e.name} - {e.grade}
                  </SelectItem>
                ))}
              {classes && classes.length === 0 && (
                <SelectLabel>No class yet</SelectLabel>
              )}
            </SelectContent>
          </Select>
        </div>
      </div>
      <Table>
        <TableCaption>
          {data && data.meta.totalItem === 0 ? (
            <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
              <SearchIcon className="w-10 h-10" />
              No students yet.
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
        </TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead>Name</TableHead>
            <TableHead>Email</TableHead>
            <TableHead>NISN</TableHead>
            <TableHead>Class - Grade</TableHead>
            <TableHead>Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {data &&
            data.data.length > 0 &&
            data.data.map((student) => {
              return (
                <TableRow key={student.id}>
                  <TableCell>
                    {student.firstName + " " + student.lastName}
                  </TableCell>
                  <TableCell>{student.email}</TableCell>
                  <TableCell>{student.nisn}</TableCell>
                  <TableCell>
                    <Badge className="bg-primary/20 text-primary text-md px-2 mb-2">
                      {classes && student.classId
                        ? classMap[student.classId]
                        : ""}
                    </Badge>
                  </TableCell>
                  <TableCell>
                    <Link to={`/student/${student.id}`}>
                      <Button>View Statistics</Button>
                    </Link>
                  </TableCell>
                </TableRow>
              );
            })}
        </TableBody>
      </Table>
    </div>
  );
}
