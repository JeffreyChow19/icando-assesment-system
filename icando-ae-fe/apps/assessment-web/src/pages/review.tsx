import Layout from "../layouts/layout.tsx";
import { useQuery } from "@tanstack/react-query";
import { getQuizReview } from "../services/quiz.ts";
import { formatDate } from "../utils/format-date.ts";
import { useEffect, useMemo } from "react";
import { QuestionAnswerCard } from "../components/question-answer-card.tsx";
import { LoadingComponent } from "../components/loading.tsx";
import { useNavigate } from "react-router-dom";
import { AxiosError } from "axios";
import { AlertTriangleIcon } from "lucide-react";

export const Review = () => {
  const { data, isLoading, error } = useQuery({
    queryKey: ["studentQuiz"],
    queryFn: () => getQuizReview(),
    retry: false,
    refetchOnWindowFocus: false,
    refetchOnReconnect: false,
    refetchOnMount: false,
  });

  const navigate = useNavigate();

  const orderedQuestions = useMemo(
    () => data?.quiz?.questions?.sort((a, b) => a.id.localeCompare(b.id)) || [],
    [data],
  );
  const orderedAnswers = useMemo(
    () =>
      data?.studentAnswers?.sort((a, b) =>
        a.questionId.localeCompare(b.questionId),
      ) || [],
    [data],
  );

  useEffect(() => {
    if (
      error &&
      error instanceof AxiosError &&
      error.response &&
      error.response.data.errors.includes("start")
    ) {
      navigate("/");
    }
  }, [error]);

  return (
    <Layout pageTitle="Review" showTitle={false} showNavigation={false}>
      <div className="flex flex-col gap-2">
        {error &&
          error instanceof AxiosError &&
          error.response &&
          error.response.data.errors.includes("end") && (
            <div className="flex flex-col items-center justify-center gap-1 text-center">
              <AlertTriangleIcon className="text-yellow-400 size-14" />
              <h1 className="font-bold text-2xl">Maaf!</h1>
              <h2 className="font-semibold">
                Anda belum dapat me-review kuis ini karena kuis masih
                berlangsung.
              </h2>
            </div>
          )}
        {!error && data && data?.quiz && (
          <>
            <h3 className="text-xl font-bold">{data?.quiz.name}</h3>
            <h2 className="text-4xl text-primary font-bold">
              {data?.totalScore}
              <span className="text-muted-foreground font-normal">/100 </span>
            </h2>
            <p className="text-muted-foreground">
              Dikumpulkan pada{" "}
              {data.completedAt ? formatDate(new Date(data.completedAt)) : "-"}
            </p>
            {orderedQuestions.map((question, index) => (
              <QuestionAnswerCard
                key={index}
                question={question}
                answerId={orderedAnswers[index]? orderedAnswers[index].answerId:null}
                questionNumber={index + 1}
              />
            ))}
          </>
        )}
        {isLoading && <LoadingComponent />}
      </div>
    </Layout>
  );
};
