import { useStudentQuiz } from "../context/user-context.tsx";
import { Navigate, useSearchParams } from "react-router-dom";
import { ReactNode } from "react";
import { LoadingPage } from "../pages/loading.tsx";
import { setToken } from "../utils/local-storage.ts";

export const ProtectedRoute = ({ children }: { children: ReactNode }) => {
  const [searchParams] = useSearchParams();

  // fetch token from param
  const token = searchParams.get("token");

  // save token
  if (token !== null) {
    setToken(token);

    // remove token search params from URL without reloading page or adding to history
    const noParams = window.location.origin + window.location.pathname;
    window.history.replaceState({ path: noParams }, "", noParams);
  }
  const { studentQuiz, loading } = useStudentQuiz();

  if (loading) {
    return <LoadingPage />;
  }
  if (studentQuiz) {
    return children;
  }
  return <Navigate to={"/"} />;
};
