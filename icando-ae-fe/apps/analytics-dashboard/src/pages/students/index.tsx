import { Layout } from "../../layouts/layout";
import { StudentsTable } from '../../components/student/students-table.tsx';

export const Students = () => {
  return (
    <Layout pageTitle="Students" showTitle={true}>
      <StudentsTable />
    </Layout>
  );
};
