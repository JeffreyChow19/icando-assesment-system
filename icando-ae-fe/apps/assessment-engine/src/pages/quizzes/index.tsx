import { Layout } from '../../layouts/layout.tsx';
import { QuizzesTable } from '../../components/quizzes/quizzes-table.tsx';

export const Quizzes = () => {
  return (
    <Layout pageTitle={'Quizzes'} showTitle={true}>
      <QuizzesTable />
    </Layout>
  )
}
