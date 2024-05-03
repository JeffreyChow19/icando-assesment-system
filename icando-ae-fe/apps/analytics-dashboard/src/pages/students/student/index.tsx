import { Layout } from "../../../layouts/layout.tsx";
import { Statistics } from "../../../components/student/statistics.tsx";

export const StudentStatistics = () => {
  const breadcrumbs = [
    {
      title: "Student",
      link: "/student",
    },
  ];

  return (
    <Layout
      pageTitle="Student Statistics"
      showTitle={true}
      breadcrumbs={breadcrumbs}
      withBack={true}
    >
      <Statistics />
    </Layout>
  );
};
