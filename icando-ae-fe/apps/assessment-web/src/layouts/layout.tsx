import { ReactNode, useState } from "react";
import { Navigation, SideBar } from "./navigation.tsx";
import { Helmet } from "react-helmet-async";
import { ProtectedRoute } from "../components/protected-route";

type LayoutProps = {
  children: ReactNode;
  pageTitle: string;
  showTitle: boolean;
};

export const Layout = ({ children, pageTitle, showTitle }: LayoutProps) => {
  const [sidebarOpen, setSidebarOpen] = useState(false);

  const toggleSidebar = () => {
    setSidebarOpen(!sidebarOpen);
  };

  return (
    <ProtectedRoute>
      <>
        <Helmet>
          <title>{pageTitle}</title>
        </Helmet>
        <div className="flex flex-col items-center w-full max-w-md mx-auto min-h-screen bg-primary overflow-hidden">
          <Navigation pageTitle={pageTitle} toggleSidebar={toggleSidebar} />
          <div
            className={`relative w-full flex-grow bg-[#EDF3FF] overflow-hidden p-5 ${sidebarOpen ? "rounded-tl-3xl" : "rounded-t-3xl"}`}
          >
            <SideBar sidebarOpen={sidebarOpen} />
            <main className="w-full p-5 rounded-t-3xl">
              {showTitle && (
                <h1 className="text-lg font-bold mb-2">{pageTitle}</h1>
              )}
              {children}
            </main>
          </div>
        </div>
      </>
    </ProtectedRoute>
  );
};

export default Layout;
