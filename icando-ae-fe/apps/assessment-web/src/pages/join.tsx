import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@ui/components/ui/card.tsx"
import { Input } from "@ui/components/ui/input"
import { Button } from "@ui/components/ui/button"
import { Helmet } from "react-helmet-async";
import { useState } from "react";
import { Link } from "react-router-dom";
import { setToken } from "../utils/local-storage.ts";



export const Join = () => {
  const pageTitle = "ICANDO"
  const [quizToken, setQuizToken] = useState('')
  // const navigate = useNavigate();
  const handleJoin = () => { 
    setToken(quizToken)
    // navigate(`/`);
    // navigate(`/?token=${quizToken}`);
  }
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
                <CardTitle>Masukkan Token</CardTitle>
              </CardHeader>
              <CardContent>
                <form>
                  <div className="grid w-full items-center gap-4">
                    <div className="flex flex-col space-y-1.5">
                      <Input
                        id="token"
                        placeholder="Token Kuis"
                        value={quizToken}
                        onChange={(e) => setQuizToken(e.target.value)} />
                    </div>
                  </div>
                </form>
              </CardContent>
              <CardFooter className="flex justify-end">
                <Button disabled={quizToken == '' ? true : false} onClick={handleJoin}>
                  <Link to={`/?token=${quizToken}`}>Masuk ke Kuis</Link>
                </Button>
              </CardFooter>
            </Card>
          </main>
        </div>
      </div>
    </>
  );
};

export default Join;
