import { useStudentQuiz } from "../context/user-context.tsx";
import { Navigate, useSearchParams } from "react-router-dom";
import { ReactNode } from "react";
import { LoadingPage } from "../pages/loading.tsx";
import { setToken } from "../utils/local-storage.ts";

export const ProtectedRoute = ({ children }: { children: ReactNode }) => {
  const [searchParams, setSearchParams] = useSearchParams();

  // fetch token from param
  const token = searchParams.get("token");

  // save token
  if (token !== null) {
    setToken(token);

    searchParams.delete("token");
    setSearchParams(searchParams);
    // console.log(token);
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
