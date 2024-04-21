import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@ui/components/ui/card.tsx";
import { Helmet } from "react-helmet-async";
import { useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { setToken } from "../utils/local-storage.ts";
import { useStudentProfile, useStudentQuiz } from "../context/user-context.tsx";

export const Join = () => {
  const pageTitle = "ICANDO";
  const [searchParams] = useSearchParams();
  const studentQuizContext = useStudentQuiz();
  const studentProfileContext = useStudentProfile();
  const navigate = useNavigate();

  useEffect(() => {
    if (searchParams && searchParams.get("token")) {
      setToken(searchParams.get("token") as string);
      studentProfileContext.refresh();
      studentQuizContext.refresh();
      navigate("/");
    }
  }, [searchParams, studentQuizContext, studentProfileContext, navigate]);

  return (
    <>
      <Helmet>
        <title>{pageTitle}</title>
      </Helmet>
      <div className="flex flex-col items-center w-full max-w-md mx-auto min-h-screen bg-primary overflow-hidden">
        <header className="relative w-full flex flex-row justify-between px-2 lg:px-6 py-4 z-20 items-center">
          <h1 className="text-center flex-grow text-white font-semibold text-2xl">
            {pageTitle}
          </h1>
        </header>
        <div
          className={`relative w-full flex-grow bg-[#EDF3FF] overflow-hidden p-5 rounded-t-3xl`}
        >
          <main className="w-full p-5 rounded-t-3xl">
            <Card className="w-[350px]">
              <CardHeader>
                <CardTitle>Join Quiz</CardTitle>
              </CardHeader>
              <CardContent>
                {studentProfileContext.loading || studentQuizContext.loading ? (
                  <p>Tunggu sebentar ...</p>
                ) : (
                  <p>Silakan mengikuti quiz melalui tautan yang diberikan.</p>
                )}
              </CardContent>
            </Card>
          </main>
        </div>
      </div>
    </>
  );
};

export default Join;
