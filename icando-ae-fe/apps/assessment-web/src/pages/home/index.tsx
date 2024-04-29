import { Layout } from "../../layouts/layout.tsx";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@ui/components/ui/card.tsx";
import { Button } from "@ui/components/ui/button.tsx";
import { useNavigate } from "react-router-dom";
import { Badge } from "@ui/components/ui/badge.tsx";
import {
  useStudentProfile,
  useStudentQuiz,
} from "../../context/user-context.tsx";
import { startQuiz } from "../../services/quiz.ts";
import { onErrorToast } from "../../components/error-toast.tsx";
import { formatDate, formatHour } from "../../utils/format-date.ts";

export const Home = () => {
  const { quiz } = useStudentQuiz();
  const { student } = useStudentProfile();
  const currentDate = new Date();
  const navigate = useNavigate();

  const handleStartQuiz = async () => {
    try {
      const quizAttempt = await startQuiz();
      if (quizAttempt) {
        navigate(`/quiz/1`);
      }
    } catch (error) {
      console.error("Failed to start quiz:", error);
      onErrorToast(error as Error);
    }
  };
  return (
    <Layout pageTitle={"Home"} showTitle={false} showNavigation={false}>
      {quiz && student ? (
        <>
          <h1 className="text-lg mb-2">Selamat datang, {student.firstName}</h1>
          <Card className="space-x-2">
            <CardHeader className="flex flex-row justify-between items-center">
              <CardTitle>{quiz.name ? quiz.name : "Untitled Quiz"}</CardTitle>
            </CardHeader>
            <CardContent className="space-y-2">
              <CardDescription>
                {quiz.subject && quiz.subject.length > 0 && (
                  <div className="flex flex-wrap gap-2">
                    {quiz.subject.map((subject) => (
                      <Badge key={subject}>{subject}</Badge>
                    ))}
                  </div>
                )}
              </CardDescription>
              <div>
                <div className="text-left font-medium text-gray-700">
                  Durasi Pengerjaan:
                </div>
                <div className="text-left font-semibold text-black flex flex-wrap">
                  {quiz.duration} menit
                </div>
              </div>
              <div>
                <div className="text-left font-medium text-gray-700">
                  Mulai Pengerjaan:
                </div>
                <div className="text-left font-semibold text-black flex flex-wrap">
                  {" "}
                  {formatDate(new Date(quiz.startAt))}{" - "}
                  {formatHour(new Date(quiz.startAt))}
                </div>
              </div>
              <div>
                <div className="text-left font-medium text-gray-700">
                  Batas Pengerjaan:
                </div>
                <div className="text-left font-semibold text-black flex flex-wrap">
                  {" "}
                  {formatDate(new Date(quiz.endAt))}{" - "}
                  {formatHour(new Date(quiz.endAt))}
                </div>
              </div>
            </CardContent>
            <CardFooter className={
              `flex 
              ${new Date(quiz.startAt) <= currentDate && new Date(quiz.endAt) >= currentDate
                ? "justify-end" : "justify-center"}`}>
              {new Date(quiz.startAt) <= currentDate && new Date(quiz.endAt) >= currentDate ? (
                <Button
                  className="flex flex-row justify-between space-x-2"
                  onClick={handleStartQuiz}
                >
                  Mulai
                </Button>
              ) : (
                <div className="text-red-400">
                  {new Date(quiz.startAt) > currentDate ? "Kuis belum dimulai" : "Kuis telah berakhir"}
                </div>
              )}
            </CardFooter>
          </Card>
        </>
      ) : null}
    </Layout>
  );
};
