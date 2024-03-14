import {Layout} from "../../layouts/layout.tsx";
import {QuizForm} from "../../components/quiz/quiz-form.tsx";

export const NewQuizPage = () => {
  return (
    <Layout pageTitle="New Quiz" showTitle={true}>
      <QuizForm/>
    </Layout>
  )
}

