import { useParams } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { getStudentStatistics } from "../../services/student.ts";
import { StudentInfo } from "./student-info.tsx";
import CompetencyChart from "../quiz/competency-chart.tsx";
import { QuizStatisticsChart } from "../quiz/quiz-statistics-chart.tsx";
import { QuizList } from "../quiz/quiz-list.tsx";
import { StatsCard } from "../ui/stats-card.tsx";

export const Statistics = () => {
  const params = useParams<{ studentId: string }>();

  const { data, isLoading } = useQuery({
    queryKey: ["studentId", params.studentId],
    queryFn: () => getStudentStatistics(params.studentId!),
    enabled: !!params.studentId,
  });

  return (
    <div className="flex flex-col gap-10 w-full">
      {data && !isLoading && (
        <>
          <StatsCard className="w-fit">
            <StudentInfo data={data.studentInfo} />
          </StatsCard>

          <div className="flex gap-10 flex-wrap">
            <StatsCard className="w-fit">
              <p className="text-center text-xl font-medium">Quiz Statistics</p>
              <QuizStatisticsChart
                pass={data.performance.quizzesPassed}
                fail={data.performance.quizzesFailed}
                isLoading={isLoading}
              />
            </StatsCard>

            <StatsCard className="w-fit">
              <p className="text-center text-xl font-medium">
                Competency Statistics
              </p>
              <CompetencyChart data={data.competency} />
            </StatsCard>
          </div>

          <StatsCard>
            <p className="text-center text-xl font-medium">Quiz History</p>
            <QuizList quizzes={data.quizzes} />
          </StatsCard>
        </>
      )}
    </div>
  );
};
