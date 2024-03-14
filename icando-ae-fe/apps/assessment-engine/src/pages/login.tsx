import React, { useEffect, useState } from 'react';
import { cn } from "@ui/lib/utils";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@ui/components/ui/form";
import { Input } from "@ui/components/ui/input";
import { Button } from "@ui/components/ui/button";
import { AxiosError } from "axios";
import { useUser } from "../context/user-context";
import { useNavigate } from "react-router-dom";
import { toast } from "@ui/components/ui/use-toast";
import { Helmet } from "react-helmet-async";
import { login } from "../services/auth";
import { useMutation } from "@tanstack/react-query";

interface UserAuthFormProps extends React.HTMLAttributes<HTMLDivElement> {}

const formSchema = z.object({
  email: z.string().min(5).max(320),
  password: z.string().min(8).max(127),
});

export function LoginPage({ className, ...props }: UserAuthFormProps) {
  const [fieldError, setFieldError] = useState<string | null>(null);
  const { refresh, user } = useUser();
  const navigate = useNavigate();

  useEffect(() => {
    if (user) {
      navigate('/');
    }
  }, [navigate, user]);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const mutation = useMutation({
    mutationFn: async (payload: z.infer<typeof formSchema>) => {
      await login(payload);
    },
    onSuccess: () => {
      refresh();
      toast({
        description: "Login berhasil",
      });
      navigate("/");
    },
    onError: (e: Error) => {
      if (e instanceof AxiosError) {
        if (e.response && e.response.status === 400) {
          const message = e.response.data.message;

          if (message) {
            setFieldError(message);
          }
        }

        if (e.response) {
          const message = e.response.data.error;
          if (message) {
            toast({
              description: message,
            });
          }
        }
      } else {
        toast({
          variant: "destructive",
          description: e.message,
        });
      }
    },
  });

  return (
    <>
      <Helmet>Login</Helmet>
      <main className={"absolute w-full min-h-screen card"}>
        <div className={"mx-auto mt-10 max-w-md"}>
          <div className={"primary p-8 shadow-2xl"}>
            <h1 className={"text-xl font-bold text-center mb-4"}>
              ICANDO Assessment Engine
            </h1>
            <div className={cn("grid gap-6", className)} {...props}>
              <Form {...form}>
                <form
                  onSubmit={form.handleSubmit((values) =>
                    mutation.mutate(values),
                  )}
                >
                  <FormField
                    control={form.control}
                    name={"email"}
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Email</FormLabel>
                        <FormControl>
                          <Input placeholder={"Email"} {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name={"password"}
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Password</FormLabel>
                        <FormControl>
                          <Input
                            placeholder={"Password"}
                            type={"password"}
                            {...field}
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  {fieldError && (
                    <p className={"mt-2 text-red-600 text-sm"}>{fieldError}</p>
                  )}
                  <Button
                    variant={"default"}
                    type={"submit"}
                    disabled={mutation.isPending}
                    className={"mt-6 w-full"}
                  >
                    Login
                  </Button>
                </form>
              </Form>
            </div>
          </div>
        </div>
      </main>
    </>
  );
}
