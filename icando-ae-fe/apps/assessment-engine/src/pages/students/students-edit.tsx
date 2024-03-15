import { Layout } from '../../layouts/layout.tsx';
import { StudentsForm } from '../../components/students/students-form.tsx';
import { useParams } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { getStudent } from '../../services/student.ts';
import { LoadingComponent } from '../../components/loading.tsx';

export const StudentsEdit = () => {
  const { id } = useParams<{ id: string }>();

  const { data, isLoading } = useQuery({
    queryKey: ['student', id],
    queryFn: () => getStudent(id!),
  });

  return (
    <Layout pageTitle={'Edit Student'} showTitle={true}>
      {isLoading || !data ? <LoadingComponent /> : <StudentsForm student={data} />}
    </Layout>
  );
};
