import { Badge } from "@ui/components/ui/badge";
import { QuizDetail } from "../../../interfaces/quiz";
import { formatDate, formatHour } from "../../../utils/format-date.ts";


export const QuizInfo = ({ quiz }: { quiz: QuizDetail }) => {
    return (
        <div className="flex flex-col gap-6 max-w-[500px]">
            <div>
                <strong>Quiz Name:</strong>
                <p>{quiz.name ? quiz.name : "Untitled Quiz"}</p>
            </div>
            <div>
                <strong>Quiz Subject:</strong>
                {quiz.subject && quiz.subject.length > 0 ? (
                    <div className="flex flex-wrap gap-2">
                        {quiz.subject.map((subject, index) => (
                            <Badge key={index}>{subject}</Badge>
                        ))}
                    </div>
                ) : <p>No subjects listed.</p>}
            </div>
            <div>
                <strong>Passing Grade:</strong>
                <p>{quiz.passingGrade ?? 'No passing grade set'}</p>
            </div>
            <div>
                <strong>Duration:</strong>
                <p>{quiz.duration ? `${quiz.duration} minutes` : "No duration set"}</p>
            </div>
            <div className="flex flex-col">
                <strong>Start At:</strong>
                {quiz.startAt ?
                    <div className="flex flex-row">
                        <Badge key={formatDate(new Date(quiz.startAt))} className="mr-2" variant={"outline"}>{formatDate(new Date(quiz.startAt))}</Badge>
                        <Badge key={formatHour(new Date(quiz.startAt))} variant={"outline"}>{formatHour(new Date(quiz.startAt))}</Badge>
                    </div> : <p>No start date set</p>
                }
            </div>
            <div className="flex flex-col">
                <strong>End At:</strong>
                {quiz.endAt ?
                    <div className="flex flex-row">
                        <Badge key={formatDate(new Date(quiz.endAt))} className="mr-2" variant={"outline"}>{formatDate(new Date(quiz.endAt))}</Badge>
                        <Badge key={formatHour(new Date(quiz.endAt))} variant={"outline"}>{formatHour(new Date(quiz.endAt))}</Badge>
                    </div> : <p>No start date set</p>
                }
            </div>
            <div>
                <strong>CLasses:</strong>
                {quiz.classes && quiz.classes.length > 0 ? (
                    <div className="flex flex-wrap gap-2">
                        {quiz.classes.map((classItem, index) => (
                            <Badge key={index}>{classItem.name}</Badge>
                        ))}
                    </div>
                ) : <p>No classes listed.</p>}
            </div>
        </div>
    );
};