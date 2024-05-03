import { useQuery } from "@tanstack/react-query";
import { useNavigate, useParams, useSearchParams } from "react-router-dom";
import { QuizDetailCard } from "../../components/quiz/quiz-detail-card.tsx";
import { Layout } from "../../layouts/layout.tsx";
import { getQuiz } from "../../services/quiz.ts";

export const QuizDetail = () => {
  const breadcrumbs = [
    {
      title: "Quiz",
      link: "/quiz",
    },
  ];
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();

  const { data, loading } = useQuery({
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
      {data && <QuizDetailCard quiz={data.quiz} />}
    </Layout>
  );
};
