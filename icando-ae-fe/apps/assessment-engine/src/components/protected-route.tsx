import { useUser } from "../context/user-context.tsx";
import { Navigate } from "react-router-dom";
import { ReactNode } from "react";
import { LoadingPage } from "../pages/loading.tsx";

export const ProtectedRoute = ({ children }: { children: ReactNode }) => {
  const { user, loading } = useUser();
  if (loading) {
    return <LoadingPage />;
  }
  if (user) {
    return children;
  }

  return <Navigate to={"/login"} />;
};
