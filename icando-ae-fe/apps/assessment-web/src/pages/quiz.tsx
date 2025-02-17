import { Layout } from "../layouts/layout.tsx";
import { useStudentQuiz } from "../context/user-context.tsx";
import { Card, CardContent } from "@ui/components/ui/card.tsx";
import { useMutation, useQuery } from "@tanstack/react-query";
import { getQuizDetail, submitQuiz, updateAnswer } from "../services/quiz.ts";
import { useEffect, useState } from "react";
import { onErrorToast } from "../components/error-toast.tsx";
import { useNavigate, useParams } from "react-router-dom";
import { Button } from "@ui/components/ui/button.tsx";
import Countdown, { zeroPad } from "react-countdown";
import { cn } from "@ui/lib/utils.ts";
import { toast } from "@ui/components/ui/use-toast.ts";
import { useAlert, useConfirm } from "../context/alert-dialog.tsx";
import { formatDate, formatHour } from "../utils/format-date.ts";

export const Quiz = () => {
  const { number } = useParams();
  const navigate = useNavigate();

  const {
    studentQuiz,
    questions,
    studentAnswers,
    setQuestions,
    setStudentAnswers,
    refresh
  } = useStudentQuiz();
  const quiz = studentQuiz?.quiz;
  const [time, setTime] = useState<Date>();

  const { data, isLoading, error } = useQuery({
    queryKey: ["studentQuiz"],
    queryFn: () => getQuizDetail(),
    retry: false,
    refetchOnWindowFocus: false,
    refetchOnReconnect: false,
    refetchOnMount: false,
  });

  const confirm = useConfirm();
  const alert = useAlert();

  useEffect(() => {
    if (data) {
      setTime(
        getRemainingTime(data.startedAt, data.quiz!.endAt, data.quiz!.duration),
      );
    }
  }, [data]);

  useEffect(() => {
    if (error) {
      onErrorToast(error);
      navigate("/");
    }

    if (!isLoading && data && data.quiz?.questions && data.studentAnswers) {
      setQuestions(data.quiz.questions);
      setStudentAnswers(data.studentAnswers);
    }
  }, [isLoading, data, error, setQuestions, setStudentAnswers, navigate]);

  const getRemainingTime = (
    startAt: string,
    endAt: string,
    duration: number,
  ) => {
    const startTime = new Date(startAt);
    const endQuizTime = new Date(endAt);
    const endTime = new Date(startTime.getTime() + duration * 60000);
    return endQuizTime.getTime() < endTime.getTime() ? endQuizTime : endTime;
  };

  const mutation = useMutation({
    mutationFn: async (choiceId: number) => {
      return await updateAnswer(questions![parseInt(number!) - 1].id, choiceId);
    },
    onSuccess: (_, choiceId) => {
      const foundIdx = studentAnswers!.findIndex(
        (answer) => answer.questionId === questions![parseInt(number!) - 1].id,
      );
      const updatedAnswers = [...studentAnswers!];
      if (foundIdx !== -1) {
        updatedAnswers[foundIdx] = {
          ...updatedAnswers[foundIdx],
          answerId: choiceId,
        };
      } else {
        updatedAnswers.push({
          questionId: questions![parseInt(number!) - 1].id,
          answerId: choiceId,
        });
      }
      setStudentAnswers(updatedAnswers);
    },
    onError: (err: Error) => {
      onErrorToast(err);
    },
  });

  const onChooseAnswer = (choiceId: number) => {
    mutation.mutate(choiceId);
  };

  const renderer = ({
    hours,
    minutes,
    seconds,
    completed,
  }: {
    hours: number;
    minutes: number;
    seconds: number;
    completed: boolean;
  }) => {
    if (completed) {
      submit();
    } else if (hours === 0 && minutes === 5 && seconds === 0) {
      alert({
        title: "Waktu tersisa 5 menit lagi!",
        body: "Perhatikan waktu mengerjakan kuis.",
        cancelButton: "Oke",
      });
      return (
        <span className="font-bold text-primary">
          {zeroPad(hours)}:{zeroPad(minutes)}:{zeroPad(seconds)}
        </span>
      );
    } else {
      return (
        <span className="font-bold text-primary">
          {zeroPad(hours)}:{zeroPad(minutes)}:{zeroPad(seconds)}
        </span>
      );
    }
  };

  const submit = async () => {
    try {
      submitQuiz().then(() => {
        return refresh();
      }).then(() => {
        navigate("/submit");
      })
    } catch (err) {
      toast({
        variant: "destructive",
        description: "Failed to submit quiz! Please try again!",
      });
    }
  };

  return (
    <Layout pageTitle={"Quiz"} showTitle={false} showNavigation={true}>
      <h1 className="text-lg font-bold">Quiz</h1>
      {quiz && (
        <h3 className="text-sm mb-2">{`Versi ${formatDate(new Date(quiz.publishedAt))} ${formatHour(new Date(quiz.publishedAt))}`}</h3>
      )}
      {questions && number && (
        <div className="flex flex-col flex-grow justify-between">
          <div className="flex flex-col gap-3">
            <div className="w-full flex justify-end">
              <Countdown date={time} renderer={renderer} />
            </div>
            <Card className="mb-2">
              <CardContent className="p-4">
                <p>{questions[parseInt(number) - 1].text}</p>
              </CardContent>
            </Card>
            {studentAnswers &&
              questions[parseInt(number) - 1].choices.map((choice) => (
                <Button
                  key={choice.id}
                  className={cn(
                    choice.id ===
                      studentAnswers.find(
                        (answer) =>
                          answer.questionId ===
                          questions![parseInt(number) - 1].id,
                      )?.answerId
                      ? "bg-blue border-blue-foreground border-2 hover:bg-blue"
                      : "hover:bg-blue bg-background ",
                    "w-full shadow-md py-3 rounded-lg text-foreground justify-start whitespace-normal h-fit",
                  )}
                  onClick={() => onChooseAnswer(choice.id)}
                >
                  {choice.text}
                </Button>
              ))}
          </div>
          <div className="flex gap-3">
            {parseInt(number) > 1 && (
              <Button
                className="text-primary bg-background hover:bg-background/80 rounded-full w-full"
                onClick={() => navigate(`/quiz/${parseInt(number) - 1}`)}
              >
                Kembali
              </Button>
            )}
            {parseInt(number) !== questions.length && (
              <Button
                className="rounded-full w-full"
                onClick={() => navigate(`/quiz/${parseInt(number) + 1}`)}
              >
                Lanjut
              </Button>
            )}
            {parseInt(number) === questions.length && (
              <Button
                className="rounded-full w-full"
                onClick={() =>
                  confirm({
                    title: "Apakah kamu yakin ingin mengumpulkan jawaban?",
                    body: "Jawaban tidak dapat diubah ketika sudah dikumpulkan",
                    cancelButton: "Tidak",
                    actionButton: "Ya",
                  }).then((res) => {
                    if (res) submit();
                  })
                }
              >
                Kirim
              </Button>
            )}
          </div>
        </div>
      )}
    </Layout>
  );
};
