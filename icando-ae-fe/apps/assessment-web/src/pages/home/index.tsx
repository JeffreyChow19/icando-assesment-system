import { Layout } from '../../layouts/layout.tsx';
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@ui/components/ui/card.tsx";
import { Button } from "@ui/components/ui/button.tsx";
import { Link, useNavigate } from "react-router-dom";
import { Badge } from "@ui/components/ui/badge.tsx";
import { useStudentProfile, useStudentQuiz } from "../../context/user-context.tsx";
import { startQuiz } from '../../services/quiz.ts';


export const Home = () => {
  const { studentQuiz } = useStudentQuiz()
  const { student } = useStudentProfile()
  const currentDate = new Date();
  const navigate = useNavigate();

  const handleStartQuiz = async () => {
    try {
      const quizAttempt = await startQuiz(); 
      // todo @ livia change navigation
      if (quizAttempt) {
        navigate(`/`);
      }
    } catch (error) {
      console.error('Failed to start quiz:', error);
    }
  };
  
  function formatDate(date: Date): string {
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0'); 
    const year = date.getFullYear().toString().slice(-2);

    return `${day}-${month}-${year}`;
  }
  function formatHour(date: Date): string {
    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');

    return `${hours}:${minutes}`;
  }
  return (
    <Layout pageTitle={'Home'} showTitle={false} showNavigation={false}>
      {studentQuiz && student ? (
        <>
          <h1 className="text-lg mb-2">Selamat datang, {student.firstName}</h1>
          {new Date(studentQuiz.endAt) < currentDate ? (
            <p>Kuis sudah berakhir.</p>
          ) : new Date(studentQuiz.startAt) > currentDate ? (
            <p>Kuis belum dimulai. Kuis akan dimulai tanggal {formatDate(new Date(studentQuiz.startAt))} jam {formatHour(new Date(studentQuiz.startAt))}</p>
          ) : (
            <Card className="space-x-2">
              <CardHeader className="flex flex-row justify-between">
                <CardTitle>{studentQuiz.name ? studentQuiz.name : "Untitled Quiz"}</CardTitle>
                {studentQuiz.subject && studentQuiz.subject.length > 0 && (
                  <div className="flex flex-wrap gap-2">
                    {studentQuiz.subject.map((subject) => (
                      <Badge key={subject}>{subject}</Badge>
                    ))}
                  </div>
                )}
              </CardHeader>
              <CardContent>
                <div className="grid grid-cols-2 gap-x-4 py-2">
                  <div className="text-left font-medium text-gray-700">Durasi Pengerjaan:</div>
                  <div className="text-left text-lg font-semibold text-black">{studentQuiz.duration} menit</div>

                  <div className="text-left font-medium text-gray-700">Batas Pengerjaan:</div>
                  <div className="text-left text-lg font-semibold text-black"> {formatDate(new Date(studentQuiz.endAt))} {formatHour(new Date(studentQuiz.endAt))}</div>
                </div>
              </CardContent>
              <CardFooter className="flex justify-end">
                  <Button className="flex flex-row justify-between space-x-2" onClick={handleStartQuiz}>
                    <Link to={``}>Start</Link>
                  </Button>
              </CardFooter>
            </Card>
          )}
        </>
      ) : <p>Quiz invalid or student information not available.</p>}
    </Layout>
  );
}
