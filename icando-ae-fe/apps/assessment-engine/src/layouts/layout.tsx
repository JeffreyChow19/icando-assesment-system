import { ReactNode } from "react";
import { Navigation, SideBar } from "./navigation.tsx";
import { Helmet } from "react-helmet-async";
import { ProtectedRoute } from "../components/protected-route.tsx";
import { Button } from "@ui/components/ui/button.tsx";
import { ArrowLeftIcon } from "lucide-react";
import { useNavigate } from "react-router-dom";

export const Layout = ({
  children,
  pageTitle,
  showTitle,
  withBack,
}: {
  children: ReactNode;
  pageTitle: string;
  showTitle: boolean;
  withBack?: boolean;
}) => {
  const navigate = useNavigate();

  return (
    <ProtectedRoute>
      <Helmet>
        <title>{pageTitle}</title>
      </Helmet>
      <div className="relative min-h-screen flex flex-row">
        <SideBar />
        <div className="flex flex-col flex-grow w-full">
          <Navigation />
          <div className="flex flex-col py-8 px-4 lg:px-16 h-full w-full">
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
              <h1 className="font-bold text-lg mb-2">{pageTitle}</h1>
            )}
            {children}
          </div>
        </div>
      </div>
    </ProtectedRoute>
  );
};
