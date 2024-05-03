import dayjs from "dayjs";
import { StudentQuizReviewResponseData } from "../../../services/quiz";
import { timeDiff } from "../../../utils/date";

interface QuizInfoProps {
  data: StudentQuizReviewResponseData;
}

export const QuizInfo = ({ data }: QuizInfoProps) => {
  return (
    <div className="grid grid-cols-3 gap-y-2 gap-x-4 max-w-[400px] p-2 divide-y">
      <div className="text-right font-semibold">Student Name</div>
      <div className="col-span-2">
        {data.quiz.student
          ? `${data.quiz.student.firstName} ${data.quiz.student.lastName}`
          : ""}
      </div>
      <div className="text-right font-semibold">Started At</div>
      <div className="col-span-2">
        {dayjs(data.quiz.startedAt).format("L LT")}
      </div>
      <div className="text-right font-semibold">Submitted At</div>
      <div className="col-span-2">
        {dayjs(data.quiz.completedAt).format("L LT")}
      </div>
      <div className="text-right font-semibold">Time Taken</div>
      <div className="col-span-2">
        {timeDiff(dayjs(data.quiz.startedAt), dayjs(data.quiz.completedAt))}
      </div>
      <div className="text-right font-semibold">Grade</div>
      <div className="col-span-2">{data.quiz.totalScore}</div>
      <div className="text-right font-semibold">Status</div>
      <div className="col-span-2">
        {data.quiz.totalScore! >= data.quiz.quiz!.passingGrade
          ? "Passed"
          : `Not passed (passing grade ${data.quiz.quiz!.passingGrade})`}
      </div>
    </div>
  );
};
