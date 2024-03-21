import { QuestionList } from "./questions/question-list.tsx";
import { QuestionForm } from "./questions/question-form.tsx";
import { QuizDetail } from "src/interfaces/quiz.ts";

export const QuizForm = ({ quiz }: { quiz: QuizDetail }) => {
  console.log(quiz);
  const dummyQuestions = [];
  for (let i = 0; i < 20; i++) {
    dummyQuestions.push({
      id: "7a32d5af-aee3-49e3-8f7d-82e57d026791",
      choices: [
        { text: "Option A", id: 0 },
        { text: "Option B", id: 1 },
        { text: "Option C", id: 2 },
        { text: "Option D", id: 3 },
      ],
      text: "Apakah bumi itu bulat? Apakah bumi itu bulat?",
      answerId: 0,
      quizId: "7a32d5af-aee3-49e3-8f7d-82e57d026791",
      competencies: [
        {
          id: "72d2376d-0da8-46a7-a27c-af1255e9e725",
          numbering: "C01",
          name: "Competency 1",
          description: "Description for competency 1",
        },
      ],
    });
  }
  return (
    <>
      <div className="flex w-full justify-end">
        <QuestionForm type="new" />
      </div>
      <QuestionList questions={dummyQuestions} />
    </>
  );
};
