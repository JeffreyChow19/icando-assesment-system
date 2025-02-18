import { ReactNode, useState } from "react";
import { Navigation, SideBar } from "./navigation.tsx";
import { Helmet } from "react-helmet-async";
import { ProtectedRoute } from "../components/protected-route";

type LayoutProps = {
  children: ReactNode;
  pageTitle: string;
  showTitle: boolean;
  showNavigation: boolean;
};

export const Layout = ({
  children,
  pageTitle,
  showTitle,
  showNavigation,
}: LayoutProps) => {
  const [sidebarOpen, setSidebarOpen] = useState(false);

  const toggleSidebar = () => {
    setSidebarOpen(!sidebarOpen);
  };

  const completeTitle = `${pageTitle} - ICANDO`;

  return (
    <ProtectedRoute>
      <Helmet>
        <title>{completeTitle}</title>
      </Helmet>
      <div className="flex flex-col items-center w-full max-w-md mx-auto min-h-[100vh] bg-primary overflow-hidden">
        <Navigation
          pageTitle={pageTitle}
          toggleSidebar={toggleSidebar}
          showNavigation={showNavigation}
        />

        <div className="relative w-full h-full flex flex-col flex-grow bg-[#EDF3FF] overflow-hidden p-5">
          <SideBar sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen} />
          <main className="w-full p-5 rounded-t-3xl flex flex-col flex-grow">
            {showTitle && (
              <h1 className="text-lg font-bold mb-2">{pageTitle}</h1>
            )}
            {children}
          </main>
        </div>
      </div>
    </ProtectedRoute>
  );
};

export default Layout;
