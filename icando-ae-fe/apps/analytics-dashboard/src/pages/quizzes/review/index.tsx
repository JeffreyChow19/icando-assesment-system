import { useParams } from "react-router-dom";
import { Layout } from "../../../layouts/layout.tsx";
import { StudentQuiz } from "../../../components/quiz/review/student-quiz.tsx";

export const StudentQuizReview = () => {
  const params = useParams<{ quizid: string; studentquizid: string }>();

  const breadcrumbs = [
    {
      title: "Quiz",
      link: "/quiz",
    },
    {
      title: "Detail",
      link: `/quiz/${params.quizid}`,
    },
  ];

  return (
    <Layout
      pageTitle="Student Quiz Review"
      showTitle={true}
      breadcrumbs={breadcrumbs}
    >
      <StudentQuiz />
    </Layout>
  );
};
