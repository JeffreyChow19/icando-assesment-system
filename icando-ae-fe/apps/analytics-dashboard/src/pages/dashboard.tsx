import { Layout } from "../layouts/layout.tsx";
import {
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@ui/components/ui/card.tsx";
import { CustomCard } from "../components/ui/custom-card.tsx";
import { useQueries } from "@tanstack/react-query";
import { getLatestSubmissions, getOverview, getPerformance } from "../services/analytics.ts";
import { PieChart } from "@mui/x-charts/PieChart";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@ui/components/ui/table.tsx";
import {
  SearchIcon,
} from "lucide-react";
import { formatDate, formatHour } from "../utils/format-date.ts";
import { StudentSubmissions } from "src/interfaces/student.ts";

export const Dashboard = () => {
  const breadcrumbs = [
    {
      title: "Home",
      link: "/",
    }
  ];

  const [overview, performance, latestSubmissions] = useQueries({
    queries: [
      {
        queryKey: ["overview"],
        queryFn: () => getOverview(),
      },
      {
        queryKey: ["performance"],
        queryFn: () => getPerformance({}),
      },
      {
        queryKey: ["latest-submissions"],
        queryFn: () => getLatestSubmissions(),
      },
    ],
  });

  const processData = (submissions: StudentSubmissions[]) => {
    const reversedData = [...submissions].reverse();

    const quizCount = new Map();
    const processed = reversedData.map(submission => {
      const baseName = submission.quizName;
      const count = quizCount.get(baseName) || 0;
      quizCount.set(baseName, count + 1);
      return { ...submission, quizName: `${baseName} - v${count}` };
    });

    return processed.reverse().slice(-5);
  };

  const processedData = processData(latestSubmissions.data || []);

  return (
    <Layout pageTitle="Dashboard" showTitle={true} breadcrumbs={breadcrumbs}>
      <div className="flex flex-col gap-2">
        <div className="grid grid-cols-3 gap-2 w-full">
          <CustomCard>
            <CardHeader>
              <CardTitle>Total Students</CardTitle>
              <CardDescription>Current count of all students enrolled</CardDescription>
            </CardHeader>
            <CardContent className="text-primary text-2xl font-bold">{overview.data?.totalStudent}</CardContent>
          </CustomCard>
          <CustomCard>
            <CardHeader>
              <CardTitle>Total Classes</CardTitle>
              <CardDescription>Number of active classes available</CardDescription>
            </CardHeader>
            <CardContent className="text-primary text-2xl font-bold">{overview.data?.totalClass}</CardContent>
          </CustomCard>
          <CustomCard>
            <CardHeader>
              <CardTitle>Ongoing Quizzes</CardTitle>
              <CardDescription>Count of quizzes currently active</CardDescription>
            </CardHeader>
            <CardContent className="text-primary text-2xl font-bold">{overview.data?.totalOngoingQuiz}</CardContent>
          </CustomCard>
        </div>
        <div className="grid grid-cols-2 gap-2 w-full">
          {performance.data && (
            <CustomCard >
                <CardHeader className="items-center">
                  <CardTitle>Quizzes Performance</CardTitle>
                </CardHeader>
                <div className="flex justify-center items-center h-auto">
                  <PieChart
                    series={[
                      {
                        data: [
                          {
                            id: 0,
                            value: performance.data.quizzesPassed,
                            label: "Passed",
                          },
                          {
                            id: 1,
                            value: performance.data.quizzesFailed,
                            label: "Failed",
                          }
                        ],
                      },
                    ]}
                    width={500}
                    height={300}
                  />
                </div>
            </CustomCard>
          )}
          {latestSubmissions.data && (
            <CustomCard >
              <CardHeader className="items-center">
                <CardTitle>Students Latest Submissions</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="flex flex-col gap-y-2 items-center">
                  <Table>
                    <TableCaption>
                      {latestSubmissions.data ? (
                        null
                      ) :
                        <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
                          <SearchIcon className="w-10 h-10" />
                          No submissions or quizzes yet.
                        </div>
                      }
                    </TableCaption>
                    <TableHeader>
                      <TableRow>
                        <TableHead>Name</TableHead>
                        <TableHead className="text-center">Class</TableHead>
                        <TableHead className="text-center">Grade</TableHead>
                        <TableHead className="text-center">Quiz Name</TableHead>
                        <TableHead>Completed At</TableHead>
                      </TableRow>
                    </TableHeader>
                    <TableBody>
                      {latestSubmissions.data && processedData.map((submission) => {
                        return (
                          <TableRow key={submission.quizName + submission.completedAt}>
                            <TableCell>{submission.firstName}{" "}{submission.lastName}</TableCell>
                            <TableCell className="text-center">{submission.className}</TableCell>
                            <TableCell className="text-center">{submission.grade}</TableCell>
                            <TableCell className="text-center">{submission.quizName}</TableCell>
                            <TableCell>{formatDate(submission.completedAt)}{", "}{formatHour(submission.completedAt)}</TableCell>
                          </TableRow>
                        );
                      })}
                    </TableBody>
                  </Table>
                </div>
              </CardContent>
            </CustomCard>
          )}
        </div>
      </div>
    </Layout>
  );
};
