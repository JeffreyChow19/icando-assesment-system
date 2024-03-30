import { Link, useSearchParams } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { deleteStudent, getAllStudent } from '../../services/student.ts';
import { useConfirm } from '../../context/alert-dialog.tsx';
import { Button } from '@ui/components/ui/button.tsx';
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@ui/components/ui/table.tsx';
import { SearchIcon } from 'lucide-react';
import { Pagination } from '../pagination.tsx';
import { useDebouncedCallback } from 'use-debounce';
import { Input } from '@ui/components/ui/input.tsx';
import {
  Select,
  SelectContent,
  SelectItem, SelectLabel,
  SelectTrigger,
  SelectValue,
} from '@ui/components/ui/select.tsx';
import { getAllClasses } from '../../services/classes.ts';
import { Class } from '../../interfaces/classes.ts';


export function StudentsTable() {
  const [searchParams] = useSearchParams();
  const pageParams = searchParams.get('page');
  const nameParams = searchParams.get('name');
  const classIdParams = searchParams.get('classId');

  const [page, setPage] = useState(pageParams ? parseInt(pageParams) : 1);
  const [name, setName] = useState(nameParams ? nameParams : '');
  const [classId, setClassId] = useState(classIdParams ? classIdParams : '');

  const debouncedName = useDebouncedCallback(
    (value) => {
      setName(value);
    },
    1000,
  );

  const { data, isLoading, refetch } = useQuery({
    queryKey: ['students', page],
    queryFn: () => getAllStudent({ page: page, limit: 10, name: name, classId: classId }),
  });

  const { data: classes } = useQuery({
    queryKey: ['classes'],
    queryFn: () => getAllClasses(),
  });

  useEffect(() => {
    if (data) {
      if (data.meta.page != page) {
        setPage(data.meta.page);
      }
    }
  }, [data, page]);

  const confirm = useConfirm();


  return (
    <div className="w-full mb-2">
      <div className="w-full flex flex-row font-normal gap-x-3 justify-between">
        <div className="w-full flex flex-row font-normal gap-x-3">
          <Input className={"w-[360px]"} placeholder={"Search name ..."} value={name} onChange={(e) => debouncedName(e.target.value)} />
          <Select value={classId} onValueChange={setClassId}>
            <SelectTrigger className="w-[240px]">
              <SelectValue placeholder="Select class" />
            </SelectTrigger>
            <SelectContent>
              {classes && classes.data.map((e: Class) => (<SelectItem value={e.id} key={e.id}>
                {e.name} - {e.grade}
              </SelectItem>))
              }
              {classes && classes.data.length === 0 && <SelectLabel>No class yet</SelectLabel>}
            </SelectContent>
          </Select>
        </div>
        <div>
          <Button size={'sm'}>
            <Link to={'/students/new'}>New Student</Link>
          </Button>
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
                <Pagination page={page} totalPage={data?.meta.totalPage || 1} setPage={setPage}
                            withSearchParams={true} />
              </div>
            )
          )}
        </TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead>Name</TableHead>
            <TableHead>NISN</TableHead>
            <TableHead>Email</TableHead>
            <TableHead>Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {data &&
            data.data.length > 0 &&
            data.data.map(student => {
              return (
                <TableRow key={student.id}>
                  <TableCell>{student.firstName + ' ' + student.lastName}</TableCell>
                  <TableCell>{student.email}</TableCell>
                  <TableCell>{student.nisn}</TableCell>
                  <TableCell>
                    <div className="flex space-x-2">
                      <>
                        <Button size={'sm'}>
                          <Link to={`/students/edit/${student.id}`}>Edit</Link>
                        </Button>
                        <Button
                          size={'sm'}
                          variant={'destructive'}
                          onClick={() => {
                            confirm({
                              title: 'Are you sure?',
                              body: 'Are you sure want to delete this student?',
                            }).then(result => {
                              if (result) {
                                deleteStudent(student.id).then(() => {
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
