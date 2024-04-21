import { Layout } from "../layouts/layout.tsx";
import { useStudentQuiz } from "../context/user-context.tsx";
import { Card, CardContent } from "@ui/components/ui/card.tsx";
import { useMutation, useQuery } from "@tanstack/react-query";
import { getQuizDetail, updateAnswer, submitQuiz } from "../services/quiz.ts";
import { useEffect, useState } from "react";
import { onErrorToast } from "../components/error-toast.tsx";
import { useNavigate, useParams } from "react-router-dom";
import { Button } from "@ui/components/ui/button.tsx";
import Countdown from "react-countdown";
import { toast } from "@ui/components/ui/use-toast.ts";

export const Quiz = () => {
  const { number } = useParams();
  const navigate = useNavigate();

  const { studentQuiz, setStudentQuiz } = useStudentQuiz();
  const [time, setTime] = useState<Date>();

  const { data, isLoading, error, refetch } = useQuery({
    queryKey: ["studentQuiz"],
    queryFn: () => getQuizDetail(),
    retry: false,
  });

  useEffect(() => {
    if (data) {
      setTime(getRemainingTime(data.startedAt, data.quiz!.duration));
    }
  }, [data]);

  useEffect(() => {
    if (error) {
      onErrorToast(error);
    }

    if (!isLoading && data) {
      setStudentQuiz(data);
      console.log(data);
    }
  }, [isLoading, data, error]);

  const getRemainingTime = (startAt: string, duration: number) => {
    const startTime = new Date(startAt);
    return new Date(startTime.getTime() + duration * 60000);
  };

  const mutation = useMutation({
    mutationFn: async (choiceId: number) => {
      return await updateAnswer(
        studentQuiz!.quiz!.questions![parseInt(number!) - 1].id,
        choiceId,
      );
    },
    onSuccess: (response) => {
      console.log(response);
      refetch();
    },
    onError: (err: Error) => {
      onErrorToast(err);
    },
  });

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
      console.log("time ended");
      // TODO: submit otomatis
    } else {
      return (
        <span className="font-bold text-primary">
          {hours}:{minutes}:{seconds}
        </span>
      );
    }
  };

  const submit = async () => {
    try {
      await submitQuiz();
      toast({
        description: "Quiz submitted!",
      });
      navigate("/");
    } catch (err) {
      toast({
        variant: "destructive",
        description: "Failed to submit quiz! Please try again!",
      });
    }
  };

  return (
    <Layout pageTitle="Quiz" showTitle={true} showNavigation={true}>
      {studentQuiz &&
        studentQuiz.quiz &&
        studentQuiz.quiz.questions &&
        number && (
          <div className="flex flex-col flex-grow justify-between">
            <div className="flex flex-col gap-3">
              <div className="w-full flex justify-end">
                <Countdown date={time} renderer={renderer} />
              </div>
              <Card className="mb-2">
                <CardContent className="p-4">
                  <p>{studentQuiz.quiz.questions[parseInt(number) - 1].text}</p>
                </CardContent>
              </Card>
              {studentQuiz.quiz.questions[parseInt(number) - 1].choices.map(
                (choice) => (
                  <Button
                    className="w-full shadow-md py-3 rounded-lg bg-background text-foreground hover:bg-blue"
                    onClick={() => mutation.mutate(choice.id)}
                  >
                    {choice.text}
                  </Button>
                ),
              )}
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
              {parseInt(number) !== studentQuiz.quiz.questions.length && (
                <Button
                  className="rounded-full w-full"
                  onClick={() => navigate(`/quiz/${parseInt(number) + 1}`)}
                >
                  Lanjut
                </Button>
              )}
              {parseInt(number) === studentQuiz.quiz.questions.length && (
                <Button
                  className="rounded-full w-full"
                  onClick={() => submit()}
                >
                  Submit
                </Button>
              )}
            </div>
          </div>
        )}
    </Layout>
  );
};
