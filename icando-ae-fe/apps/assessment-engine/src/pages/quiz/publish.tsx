import { Layout } from "../../layouts/layout.tsx";
import { useQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";
import { getQuiz } from "../../services/quiz.ts";
import { LoadingComponent } from "../../components/loading.tsx";
import { QuizPublishForm } from "../../components/quiz/quiz-publish.tsx";

export const PublishQuizPage = () => {
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
    <Layout pageTitle={`Publish Quiz${data?.name? ` "${data?.name}"` : ""}`} showTitle={true} withBack={true}>
      {isLoading && <LoadingComponent />}
      {data && <QuizPublishForm id={params.id!} />}
    </Layout>
  );
};
