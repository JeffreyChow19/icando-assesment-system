import { Layout } from '../../layouts/layout.tsx';
import { StudentsForm } from '../../components/students/students-form.tsx';

export const StudentsNew = () => {
  return (
    <Layout pageTitle={'New Student'} showTitle={true}>
      <StudentsForm/>
    </Layout>
  )
}
