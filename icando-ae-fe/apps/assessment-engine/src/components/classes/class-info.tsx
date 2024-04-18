import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from "@ui/components/ui/tabs";
import { Class } from "../../interfaces/classes";
import { Button } from "@ui/components/ui/button";
import { Link } from "react-router-dom";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@ui/components/ui/table";
import { ParticipantsTable } from "./participants";

export const ClassInfo = ({
  classData,
  refetch,
}: {
  classData: Class;
  refetch: () => void;
}) => {
  return (
    <Tabs className="w-full" defaultValue="information">
      <TabsList>
        <TabsTrigger value="information">Information</TabsTrigger>
        <TabsTrigger value="participants">Participants</TabsTrigger>
      </TabsList>
      <TabsContent value="information">
        <div className="flex flex-row gap-x-3 justify-between">
          <div className="w-full flex flex-row font-normal gap-x-3"></div>
          <Button>
            <Link to={`/classes/edit/${classData.id}`}>Edit Class</Link>
          </Button>
        </div>
        <Table>
          <TableBody>
            <TableRow>
              <TableHead className="w-[10vw]">
                <TableHeader>Name</TableHeader>
              </TableHead>
              <TableCell>{classData.name}</TableCell>
            </TableRow>
            <TableRow>
              <TableHead>
                <TableHeader>Grade</TableHeader>
              </TableHead>
              <TableCell>{classData.grade}</TableCell>
            </TableRow>
            <TableRow>
              <TableHead>
                <TableHeader>Teachers</TableHeader>
              </TableHead>
              <TableCell>
                {classData.teachers && classData.teachers.length > 0 ? (
                  <Table>
                    {classData.teachers.map((teach) => {
                      return (
                        <TableRow>
                          <TableCell>{`${teach.firstName} ${teach.lastName}`}</TableCell>
                          <TableCell>{teach.email}</TableCell>
                        </TableRow>
                      );
                    })}
                  </Table>
                ) : (
                  `No teachers assigned to class "${classData.name}"`
                )}
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </TabsContent>
      <TabsContent value="participants">
        <ParticipantsTable classData={classData} refresh={refetch} />
      </TabsContent>
    </Tabs>
  );
};
