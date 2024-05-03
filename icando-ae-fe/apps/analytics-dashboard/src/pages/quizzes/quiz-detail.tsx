import { useQuery } from "@tanstack/react-query";
import { useNavigate, useParams } from "react-router-dom";
import { QuizDetailCard } from "../../components/quiz/quiz-detail-card.tsx";
import { Layout } from "../../layouts/layout.tsx";
import { getQuiz } from "../../services/quiz.ts";
import { StudentQuizTable } from "../../components/quiz/student-quiz-table.tsx";

export const QuizDetail = () => {
  const breadcrumbs = [
    {
      title: "Quiz",
      link: "/quiz",
    },
  ];
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();

  const { data, isLoading } = useQuery({
    queryKey: ["quiz", id],
    queryFn: () => getQuiz(id || ""),
  });

  if (!id) {
    navigate("/quiz");
    return;
  }

  return (
    <Layout
      pageTitle={data?.quiz.name || ""}
      showTitle={false}
      breadcrumbs={breadcrumbs}
    >
      <div className="flex flex-col gap-6">
        <QuizDetailCard quiz={data?.quiz} isLoading={isLoading} />
        {data && <StudentQuizTable quiz={data.quiz} />}
      </div>
    </Layout>
  );
};
