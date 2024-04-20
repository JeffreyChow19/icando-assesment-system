import { Layout } from "../../layouts/layout.tsx";
import { QuizForm } from "../../components/quiz/history/quiz-form.tsx";
import { useQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";
import { getQuiz } from "../../services/quiz.ts";
import { LoadingComponent } from "../../components/loading.tsx";

export const ViewQuizPage = () => {
  const params = useParams<{ id: string }>();

  const { data, isLoading } = useQuery({
    queryKey: ["quiz", params.id || "no-id"],
    queryFn: () => getQuiz(params.id!),
    enabled: !!params.id,
    refetchOnMount: false,
    refetchOnReconnect: false,
    refetchOnWindowFocus: false,
  });

  return (
    <Layout pageTitle="View Quiz" showTitle={false} withBack={true}>
      {isLoading && <LoadingComponent />}
      {data && <QuizForm quiz={data} />}
    </Layout>
  );
};
