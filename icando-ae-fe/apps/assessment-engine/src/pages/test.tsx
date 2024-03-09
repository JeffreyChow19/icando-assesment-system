import { Layout } from "../layouts/layout.tsx";
import { Button } from "@ui/components/ui/button.tsx";

export const TestPage = () => {
  return (
    <Layout pageTitle={"Tes"} showTitle={true}>
      <Button size={"lg"} variant={"destructive"}>
        Assessment Engine
      </Button>
    </Layout>
  );
};
