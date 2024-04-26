import { Fragment, ReactNode } from "react";
import { Navigation, SideBar } from "./navigation.tsx";
import { Helmet } from "react-helmet-async";
import { ProtectedRoute } from "../components/protected-route.tsx";
import { Button } from "@ui/components/ui/button.tsx";
import { ArrowLeftIcon } from "lucide-react";
import { useNavigate } from "react-router-dom";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@ui/components/ui/breadcrumb.tsx";

export const Layout = ({
  children,
  pageTitle,
  showTitle,
  withBack,
  breadcrumbs,
}: {
  children: ReactNode;
  pageTitle: string;
  showTitle: boolean;
  withBack?: boolean;
  breadcrumbs?: {
    title: string;
    link: string;
  }[];
}) => {
  const completeTitle = `${pageTitle} - ICANDO Analytics Dashboard`;

  const navigate = useNavigate();
  const breadcrumbsMarkup = (
    <Breadcrumb>
      <BreadcrumbList>
        {breadcrumbs?.map((breadcrumb, index) => (
          <Fragment key={index}>
            <BreadcrumbItem>
              <BreadcrumbLink href={breadcrumb.link}>
                {breadcrumb.title}
              </BreadcrumbLink>
            </BreadcrumbItem>
            <BreadcrumbSeparator />
          </Fragment>
        ))}
        <BreadcrumbItem>
          <BreadcrumbPage>{pageTitle}</BreadcrumbPage>
        </BreadcrumbItem>
      </BreadcrumbList>
    </Breadcrumb>
  );

  return (
    <ProtectedRoute>
      <Helmet>
        <title>{completeTitle}</title>
      </Helmet>
      <div className="relative min-h-screen flex flex-row">
        <SideBar />
        <div className="flex flex-col flex-grow w-full">
          <Navigation />
          <div className="flex flex-col py-8 px-4 lg:px-16 h-full w-full bg-muted">
            {withBack && (
              <Button
                className="mb-2"
                size={"icon"}
                variant={"ghost"}
                onClick={() => navigate(-1)}
              >
                <ArrowLeftIcon />
              </Button>
            )}
            {showTitle && (
              <h1 className="font-bold text-2xl mb-2">{pageTitle}</h1>
            )}
            {breadcrumbs && breadcrumbsMarkup}
            <div className="mt-4">{children}</div>
          </div>
        </div>
      </div>
    </ProtectedRoute>
  );
};
