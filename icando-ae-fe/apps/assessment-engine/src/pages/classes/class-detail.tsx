import { useParams } from "react-router-dom";
import { Layout } from "../../layouts/layout.tsx";
import { useQuery } from "@tanstack/react-query";
import { LoadingComponent } from "../../components/loading.tsx";
import { getClass } from "../../services/classes.ts";
import { ClassInfo } from "../../components/classes/class-info.tsx";

export const ClassDetailPage = () => {
  const { id } = useParams<{ id: string }>();

  const { data, isLoading, refetch } = useQuery({
    queryKey: ["classes", id],
    queryFn: () => getClass({ withTeacher: true, withStudents: true }, id!),
  });

  return (
    <Layout
      pageTitle={data?.data.name ?? "Class Detail"}
      showTitle={true}
      withBack={true}
    >
      {isLoading && <LoadingComponent />}
      {data && <ClassInfo refetch={refetch} classData={data.data} />}
    </Layout>
  );
};
