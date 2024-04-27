import { Layout } from "../layouts/layout.tsx";
import {
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@ui/components/ui/card.tsx";
import { CustomCard } from "../components/shared/custom-card.tsx";

export const Dashboard = () => {
  const breadcrumbs = [
    {
      title: "Page",
      link: "/",
    },
    {
      title: "Subpage",
      link: "/subpage",
    },
  ];
  return (
    <Layout pageTitle="Dashboard" showTitle={true} breadcrumbs={breadcrumbs}>
      <div className="grid grid-cols-3 gap-2 w-full">
        <CustomCard>
          <CardHeader>
            <CardTitle>Stats 1</CardTitle>
            <CardDescription>Desc 1</CardDescription>
          </CardHeader>
          <CardContent>Another fortnight lost in America</CardContent>
        </CustomCard>
        <CustomCard>
          <CardHeader>
            <CardTitle>Stats 1</CardTitle>
            <CardDescription>Desc 1</CardDescription>
          </CardHeader>
          <CardContent>Another fortnight lost in America</CardContent>
        </CustomCard>
        <CustomCard>
          <CardHeader>
            <CardTitle>Stats 1</CardTitle>
            <CardDescription>Desc 1</CardDescription>
          </CardHeader>
          <CardContent>Another fortnight lost in America</CardContent>
        </CustomCard>
      </div>
    </Layout>
  );
};
