import { useUser } from "../context/user-context.tsx";
import { Navigate, useSearchParams } from "react-router-dom";
import { ReactNode } from "react";
import { LoadingPage } from "../pages/loading.tsx";
import { saveQuizToken } from "../services/auth.ts";

export const ProtectedRoute = ({ children }: { children: ReactNode }) => {
  const { studentQuiz: user, loading, refresh } = useUser();
  const [searchParams, setSearchParams] = useSearchParams();

  // fetch token from param
  const token = searchParams.get("token");

  // save token
  if (token !== null) {
    saveQuizToken(token).then(() => {
      refresh();
      searchParams.delete("token");
      setSearchParams(searchParams);
    });
    // console.log(token);
  }

  if (loading) {
    return <LoadingPage />;
  }
  if (user) {
    return children;
  }
  return <Navigate to={"/"} />;
};
