import { useQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";
import { getStudentQuizReview } from "../../../services/quiz";
import { QuestionList } from "./question-list";
import { useMemo } from "react";
import { QuestionWithAnswer } from "../../../interfaces/quiz";
import { QuizInfo } from "./quiz-info";
import { PieChart } from "@mui/x-charts/PieChart";
import CompetencyChart from "../competency-chart";
import { StatsCard } from "../../ui/stats-card.tsx";

export const StudentQuiz = () => {
  const params = useParams<{ quizid: string; studentquizid: string }>();

  const { data, isLoading } = useQuery({
    queryKey: ["studentQuiz", params.studentquizid],
    queryFn: () => getStudentQuizReview(params.quizid!, params.studentquizid!),
    enabled: !!params.quizid && !!params.studentquizid,
  });

  const questionWithAnswer = useMemo(() => {
    if (!data || isLoading) {
      return [];
    }

    const result: QuestionWithAnswer[] = data.quiz.quiz!.questions!.map(
      (question) => {
        if (data.quiz.studentAnswers === null) {
          return {
            ...question,
            studentAnswer: null,
          };
        }

        const answer = data.quiz.studentAnswers.find(
          (answer) => answer.questionId === question.id,
        );

        return {
          ...question,
          studentAnswer: answer || null,
        };
      },
    );

    return result;
  }, [data, isLoading]);

  const questionCorrectStats = useMemo(() => {
    if (!data || isLoading) {
      return null;
    }

    const totalUnanswered =
      data.quiz.quiz!.questions!.length - data.quiz.studentAnswers!.length;
    const totalIncorrect =
      data.quiz.quiz!.questions!.length -
      totalUnanswered -
      data.quiz.correctCount!;

    return {
      totalUnanswered,
      totalIncorrect,
      totalCorrect: data.quiz.correctCount!,
    };
  }, [data, isLoading]);

  return (
    <div className="flex flex-col gap-10">
      {data && !isLoading && (
        <StatsCard className="w-fit">
          <QuizInfo data={data} />
        </StatsCard>
      )}
      <div className="flex gap-10 flex-wrap items-top">
        {data && !isLoading && (
          <StatsCard className="w-fit">
            <p className="text-center text-xl font-medium">
              Competency Statistics
            </p>
            <CompetencyChart data={data.competency} />
          </StatsCard>
        )}
        {questionCorrectStats && (
          <StatsCard className="w-fit">
            <p className="text-center text-xl font-medium">
              Question Statistics
            </p>
            <PieChart
              series={[
                {
                  data: [
                    {
                      id: 0,
                      value: questionCorrectStats.totalCorrect,
                      label: "Correct",
                    },
                    {
                      id: 1,
                      value: questionCorrectStats.totalIncorrect,
                      label: "Incorrect",
                    },
                    {
                      id: 2,
                      value: questionCorrectStats.totalUnanswered,
                      label: "Unanswered",
                    },
                  ],
                },
              ]}
              width={500}
              height={300}
            />
          </StatsCard>
        )}
      </div>
      <StatsCard className="w-full">
        <QuestionList questions={questionWithAnswer} />{" "}
      </StatsCard>
    </div>
  );
};
