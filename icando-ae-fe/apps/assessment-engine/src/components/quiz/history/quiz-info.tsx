import { Badge } from "@ui/components/ui/badge";
import { QuizDetail } from "src/interfaces/quiz";

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
        </div>
    );
};