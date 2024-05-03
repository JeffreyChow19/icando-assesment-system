import { useParams } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { getStudentStatistics } from "../../services/student.ts";
import { StudentInfo } from "./student-info.tsx";
import CompetencyChart from "../quiz/competency-chart.tsx";
import { QuizStatisticsChart } from "../quiz/quiz-statistics-chart.tsx";
import { QuizList } from "../quiz/quiz-list.tsx";

export const Statistics = () => {
  const params = useParams<{ studentId: string }>();

  const { data, isLoading } = useQuery({
    queryKey: ["studentId", params.studentId],
    queryFn: () => getStudentStatistics(params.studentId!),
    enabled: !!params.studentId,
  });

  return (
    <div className="flex flex-col gap-12">
      {data && !isLoading && (
        <>
          <StudentInfo data={data.student_info} />

          <div className="flex flex-col items-center gap-4 w-fit">
            <p className="text-center text-xl font-medium">Quiz Statistics</p>
            <QuizStatisticsChart
              pass={data.performance.quizzesPassed}
              fail={data.performance.quizzesFailed}
              isLoading={isLoading}
            />
          </div>

          <div className="flex flex-col items-center gap-4 w-fit">
            <p className="text-center text-xl font-medium">
              Competency Statistics
            </p>
            <CompetencyChart data={data.competency} />
          </div>

          <div className="flex flex-col items-center gap-4 w-fit">
            <p className="text-center text-xl font-medium">Quiz History</p>
            <QuizList quizzes={data.quizzes} />
          </div>
        </>
      )}
    </div>
  );
};
