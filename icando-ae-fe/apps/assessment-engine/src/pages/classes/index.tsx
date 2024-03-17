import { Layout } from '../../layouts/layout.tsx';
import { ClassesTable } from '../../components/classes/classes-table.tsx';

export const Classes = () => {
  return (
    <Layout pageTitle={'Classes'} showTitle={true}>
      <ClassesTable />
    </Layout>
  )
}
