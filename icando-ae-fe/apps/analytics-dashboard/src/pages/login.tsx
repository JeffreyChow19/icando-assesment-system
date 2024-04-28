import { useUser } from "../context/user-context";
import { Helmet } from "react-helmet-async";
import { LoginForm } from "../components/login/login-form.tsx";
import { Badge } from "@ui/components/ui/badge.tsx";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

export const Login = () => {
  const { user } = useUser();
  const navigate = useNavigate();

  useEffect(() => {
    if (user) {
      navigate("/");
    }
  }, [navigate, user]);

  return (
    <>
      <Helmet>
        <title>Login - ICANDO Assessment Engine</title>
      </Helmet>
      <main className={"absolute w-full min-h-screen card"}>
        <div className="grid grid-cols-2 h-screen">
          <div className="flex items-center justify-center h-full bg-secondary">
            <img src="/login-img.png" alt="teacher illustration" />
          </div>
          <div className="flex flex-col h-full bg-muted items-center justify-center -mt-2">
            <img src="/logo.png" className="w-1/3" alt="sekolah.mu logo" />
            <Badge className="bg-primary/20 text-primary text-md px-2 mb-2">
              Assessment Engine
            </Badge>
            <LoginForm />
          </div>
        </div>
      </main>
    </>
  );
};
