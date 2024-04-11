import { useParams } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { LoadingComponent } from "../../components/loading.tsx";
import { getClass } from "../../services/classes.ts";
import { ClassesForm } from "../../components/classes/classes-form.tsx";
import { Layout } from "../../layouts/layout.tsx";

export const ClassEdit = () => {
  const { id } = useParams<{ id: string }>();

  const { data, isLoading } = useQuery({
    queryKey: ["classes", id],
    queryFn: () => getClass({ withTeacher: true, withStudents: false }, id!),
  });

  return (
    <Layout pageTitle={"Edit Class"} showTitle={true} withBack={true}>
      {isLoading || !data ? (
        <LoadingComponent />
      ) : (
        <ClassesForm classes={data.data} />
      )}
    </Layout>
  );
};
