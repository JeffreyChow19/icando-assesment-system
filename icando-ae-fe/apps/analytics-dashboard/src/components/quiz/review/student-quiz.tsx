import { useQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";
import { getStudentQuizReview } from "../../../services/quiz";
import { QuestionList } from "./question-list";
import { useMemo } from "react";
import { QuestionWithAnswer } from "../../../interfaces/quiz";

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

  return (
    <div>
      <QuestionList questions={questionWithAnswer} />
    </div>
  );
};
