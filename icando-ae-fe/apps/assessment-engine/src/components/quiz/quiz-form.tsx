import {QuestionList} from "./questions/question-list.tsx";
import {QuestionForm} from "./questions/question-form.tsx";

export const QuizForm = () => {
  const dummyQuestions = []
  for (let i = 0; i < 20; i++) {
    dummyQuestions.push({
      id: "7a32d5af-aee3-49e3-8f7d-82e57d026791",
      choices: [],
      text: "Apakah bumi itu bulat? Apakah bumi itu bulat?",
      answerId: 0,
      quizId: "7a32d5af-aee3-49e3-8f7d-82e57d026791"
    })
  }
  return (
    <>
      <div className="flex w-full justify-end">
        <QuestionForm type="new"/>
      </div>
      <QuestionList questions={dummyQuestions}/>
    </>
  )
}
