import { Layout } from '../../layouts/layout.tsx';
import { StudentsTable } from '../../components/students/students-table.tsx';

export const Students = () => {
  return (
    <Layout pageTitle={'Students'} showTitle={true}>
      <StudentsTable />
    </Layout>
  )
}
