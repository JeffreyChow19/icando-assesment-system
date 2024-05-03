import { QuizzesTable } from "../../components/quiz/quizzes-table";
import { Layout } from "../../layouts/layout";

export const Quizzes = () => {
  return (
    <Layout pageTitle="Quizzes" showTitle={true}>
      <QuizzesTable />
    </Layout>
  );
};
