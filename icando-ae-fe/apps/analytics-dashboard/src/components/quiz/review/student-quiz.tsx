import { useQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";
import { QuestionList } from "./question-list";
import { useMemo } from "react";
import { QuestionWithAnswer } from "../../../interfaces/quiz";
import { QuizInfo } from "./quiz-info";
import { PieChart } from "@mui/x-charts/PieChart";
import CompetencyChart from "../competency-chart";
import { getStudentQuizReview } from "../../../services/student-quiz.ts";

export const StudentQuiz = () => {
  const params = useParams<{ quizid: string; studentquizid: string }>();

  const { data, isLoading } = useQuery({
    queryKey: ["studentQuiz", params.studentquizid],
    queryFn: () => getStudentQuizReview(params.studentquizid!),
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
    <div>
      {data && !isLoading && (
        <>
          <QuizInfo data={data} />
        </>
      )}
      <div className="flex gap-4 flex-wrap items-top my-4">
        {data && !isLoading && (
          <div>
            <p className="text-center text-xl font-medium">
              Competency Statistics
            </p>
            <CompetencyChart data={data.competency} />
          </div>
        )}
        {questionCorrectStats && (
          <div>
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
              width={600}
              height={300}
            />
          </div>
        )}
      </div>
      <QuestionList questions={questionWithAnswer} />
    </div>
  );
};
