import { Layout } from "../layouts/layout.tsx";
import { useStudentQuiz } from "../context/user-context.tsx";

export const Quiz = () => {
  const { studentQuiz } = useStudentQuiz()
  return (
    <Layout pageTitle="Quiz" showTitle={true} showNavigation={true}>
      {/* todo: quiz layout */}
      <p>
        {studentQuiz ? `${studentQuiz.quiz?.startAt} ${studentQuiz.quiz?.endAt}` : "Quiz invalid"}
      </p>
    </Layout>
  );
};
