import { useQuery } from "@tanstack/react-query";
import { Layout } from "../layouts/layout.tsx";
import { getQuizAvailability } from "../services/quiz.ts";

export const Quiz = () => {
  const { data: quizData } = useQuery({
    queryKey: ["quiz"],
    queryFn: () => getQuizAvailability(),
  });
  return (
    <Layout pageTitle="Quiz" showTitle={true}>
      {/* todo: quiz layout */}
      <p>
        {quizData ? `${quizData.startAt} ${quizData.endAt}` : "Quiz invalid"}
      </p>
    </Layout>
  );
};
