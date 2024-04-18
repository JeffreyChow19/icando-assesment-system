import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@ui/components/ui/card.tsx";
import { QuizDetail } from "../../interfaces/quiz.ts";
import { Button } from "@ui/components/ui/button.tsx";
import { Link } from "react-router-dom";
import { Badge } from "@ui/components/ui/badge.tsx";

export function QuizCard({ quiz }: { quiz: QuizDetail | null }) {
  function formatDate(date: Date): string {
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0'); // getMonth() is zero-indexed
    const year = date.getFullYear().toString().slice(-2);

    return `${day}-${month}-${year}`;
  }
  function formatHour(date: Date): string {
    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');

    return `${hours}:${minutes}`;
  }
  return (
    <Card className="space-x-2">
      <CardHeader className="flex flex-row justify-between">
        <CardTitle>
          Untitled
        </CardTitle>
        <Badge key={"subject"}>{"subject"}</Badge>
      </CardHeader>
      <CardContent>
        <div className="grid grid-cols-2 gap-x-4 py-2"> {/* Creates a two-column grid layout with a gap */}
          <div className="text-left font-medium text-gray-700">Jumlah Soal:</div>
          <div className="text-left text-lg font-semibold text-black">50 soal</div>

          <div className="text-left font-medium text-gray-700">Durasi:</div>
          <div className="text-left text-lg font-semibold text-black">50 menit</div>

          <div className="text-left font-medium text-gray-700">Batas Pengerjaan:</div>
          <div className="text-left text-lg font-semibold text-black">Senin</div>
        </div>
      </CardContent>
      <CardFooter className="flex justify-end">

        <div className="flex flex-row justify-between space-x-2">
          <Button>
            <Link to={`/quiz/publish`}>Start</Link>
          </Button>
        </div>
      </CardFooter>
    </Card>
  );
}
