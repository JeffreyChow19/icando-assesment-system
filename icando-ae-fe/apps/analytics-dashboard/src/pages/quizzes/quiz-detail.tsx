import { useQueries, useQuery } from "@tanstack/react-query";
import { useNavigate, useParams } from "react-router-dom";
import { QuizDetailCard } from "../../components/quiz/quiz-detail-card.tsx";
import { Layout } from "../../layouts/layout.tsx";
import { getQuiz } from "../../services/quiz.ts";
import { StudentQuizTable } from "../../components/quiz/student-quiz-table.tsx";
import { getPerformance } from "../../services/analytics.ts";
import { QuizStatisticsChart } from "../../components/quiz/quiz-statistics-chart.tsx";
import { useEffect } from "react";
import { Card, CardContent } from "@ui/components/ui/card.tsx";

export const QuizDetail = () => {
  const breadcrumbs = [
    {
      title: "Quiz",
      link: "/quiz",
    },
  ];
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();

  const [quizQuery, statisticsQuery] = useQueries({
    queries: [
      {
        queryKey: ["quiz-detail", id],
        queryFn: () => getQuiz(id || ""),
      },
      {
        queryKey: ["quiz-performance", id],
        queryFn: () => getPerformance({ quizId: id }),
      },
    ],
  });

  if (!id) {
    navigate("/quiz");
    return;
  }

  return (
    <Layout
      pageTitle={quizQuery.data?.quiz.name || ""}
      showTitle={false}
      breadcrumbs={breadcrumbs}
    >
      <div className="flex flex-col gap-6">
        <div className="grid grid-cols-2 gap-x-6">
          <QuizDetailCard
            quiz={quizQuery.data?.quiz}
            isLoading={quizQuery.isLoading}
          />
          <Card>
            <CardContent className="flex justify-center mt-4">
              <QuizStatisticsChart
                pass={statisticsQuery?.data?.quizzesPassed}
                fail={statisticsQuery?.data?.quizzesFailed}
                isLoading={statisticsQuery.isLoading}
              />
            </CardContent>
          </Card>
        </div>
        {quizQuery.data && <StudentQuizTable quiz={quizQuery.data.quiz} />}
      </div>
    </Layout>
  );
};
