import { ReactNode } from 'react';
// import { Navigation, SideBar } from "../layouts/navigation.tsx";
import { Helmet } from "react-helmet-async";
// import { ProtectedRoute } from "../components/protected-route.tsx";

type LayoutProps = {
  children: ReactNode;
  pageTitle: string;
  showTitle: boolean;
};

export const Layout = ({ children, pageTitle, showTitle }: LayoutProps) => {
  return (
    // <ProtectedRoute>
    <>
      <Helmet>
        <title>{pageTitle}</title>
      </Helmet>
      <div className="flex flex-col items-center w-full max-w-md mx-auto min-h-screen bg-primary overflow-hidden">
        <header className="w-full p-2.5 bg-primary text-white text-center text-2xl font-bold ">
          {pageTitle}
        </header>
        <div className="w-full flex-grow bg-[#EDF3FF] overflow-hidden p-5 rounded-t-3xl">
          <main className="w-full p-5 rounded-t-3xl">
            {showTitle && (
              <h1 className="text-lg font-bold mb-2">{pageTitle}</h1>
            )}
            {children}
          </main>
        </div>
      </div>
    </>
    // </ProtectedRoute>
  );
};

export default Layout;
