import { Link, useParams } from "react-router-dom";
import { Layout } from "../../layouts/layout.tsx";
import { ParticipantsTable } from "../../components/classes/participants.tsx";
import { useQuery } from "@tanstack/react-query";
import { LoadingComponent } from "../../components/loading.tsx";
import { getClass } from "../../services/classes.ts";
import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from "@ui/components/ui/tabs.tsx";
import { Button } from "@ui/components/ui/button.tsx";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@ui/components/ui/table.tsx";

export const ClassDetailPage = () => {
  const { id } = useParams<{ id: string }>();

  const { data, isLoading, refetch } = useQuery({
    queryKey: ["class-participants", id],
    queryFn: () => getClass({ withTeacher: true, withStudents: true }, id!),
  });

  return !isLoading && data ? (
    <Layout pageTitle={data.data.name} showTitle={true} withBack={true}>
      <Tabs className="w-full" defaultValue="information">
        <TabsList>
          <TabsTrigger value="information">Information</TabsTrigger>
          <TabsTrigger value="participants">Participants</TabsTrigger>
        </TabsList>
        <TabsContent value="information">
          <div className="flex flex-row gap-x-3 justify-between">
            <div className="w-full flex flex-row font-normal gap-x-3"></div>
            <Button>
              <Link to={`/classes/edit/${data.data.id}`}>Edit Class</Link>
            </Button>
          </div>
          <Table>
            <TableBody>
              <TableRow>
                <TableHead className="w-[10vw]">
                  <TableHeader>Name</TableHeader>
                </TableHead>
                <TableCell>{data.data.name}</TableCell>
              </TableRow>
              <TableRow>
                <TableHead>
                  <TableHeader>Grade</TableHeader>
                </TableHead>
                <TableCell>{data.data.grade}</TableCell>
              </TableRow>
              <TableRow>
                <TableHead>
                  <TableHeader>Teachers</TableHeader>
                </TableHead>
                <TableCell>
                  {data.data.teachers && data.data.teachers.length > 0 ? (
                    <Table>
                      {data.data.teachers.map((teach) => {
                        return (
                          <TableRow>
                            <TableCell>{`${teach.firstName} ${teach.lastName}`}</TableCell>
                            <TableCell>{teach.email}</TableCell>
                          </TableRow>
                        );
                      })}
                    </Table>
                  ) : (
                    `No teachers assigned to class "${data.data.name}"`
                  )}
                </TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </TabsContent>
        <TabsContent value="participants">
          <ParticipantsTable classData={data.data} refresh={refetch} />
        </TabsContent>
      </Tabs>
    </Layout>
  ) : (
    <LoadingComponent />
  );
};
