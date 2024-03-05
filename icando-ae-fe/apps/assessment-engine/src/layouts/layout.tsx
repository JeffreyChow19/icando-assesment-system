import {ReactNode} from "react";
import {Navigation, SideBar} from "./navigation.tsx";
import {Helmet} from "react-helmet-async";

export const Layout = ({
                         children,
                         pageTitle,
                         showTitle,
                       }: {
  children: ReactNode;
  pageTitle: string;
  showTitle: boolean;
}) => {
  return (
    <>
        <Helmet>
          <title>{pageTitle}</title>
        </Helmet>
        <div className="relative min-h-screen flex flex-row">
          <SideBar/>
          <div className="flex flex-col flex-grow w-full">
            <Navigation/>
            <div className="flex flex-col py-8 px-4 lg:px-16 min-h-full w-full">
              {showTitle && (
                <h1 className="font-bold text-lg mb-2">{pageTitle}</h1>
              )}
              {children}
            </div>
          </div>
        </div>
    </>
  );
};
