import { ClassesForm } from "../../components/classes/classes-form";
import { Layout } from "../../layouts/layout";

export const ClassNew = () => {
  return (
    <>
      <Layout pageTitle="New Class" showTitle={true}>
        <ClassesForm />
      </Layout>
    </>
  );
};
