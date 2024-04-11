import { useParams } from "react-router-dom";
import { Layout } from "../../layouts/layout.tsx";
import { ParticipantsTable } from "../../components/classes/participants.tsx";
import { useQuery } from "@tanstack/react-query";
import { LoadingComponent } from "../../components/loading.tsx";
import { getClass } from "../../services/classes.ts";

export const ClassParticipants = () => {
  const { id } = useParams<{ id: string }>();

  const { data, isLoading, refetch } = useQuery({
    queryKey: ["class-participants", id],
    queryFn: () => getClass({ withTeacher: false, withStudents: true }, id!),
  });

  return !isLoading && data ? (
    <Layout pageTitle={"Class Participants"} showTitle={true} withBack={true}>
      <ParticipantsTable classData={data.data} refresh={refetch} />
    </Layout>
  ) : (
    <LoadingComponent />
  );
};
