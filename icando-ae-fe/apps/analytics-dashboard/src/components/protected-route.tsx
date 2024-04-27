// import { useUser } from "../context/user-context.tsx";
// import { useNavigate } from "react-router-dom";
import { ReactNode } from "react";
// import { LoadingPage } from "../pages/loading.tsx";

export const ProtectedRoute = ({ children }: { children: ReactNode }) => {
  // const { user, loading } = useUser();
  // const navigate = useNavigate();
  // if (loading) {
  //   return <LoadingPage />;
  // }
  // if (user) {
  //   return children;
  // }
  // navigate("/login");
  return children;
};
