import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { loginFormSchema } from "./login-form-schema.tsx";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@ui/components/ui/form.tsx";
import { Input } from "@ui/components/ui/input.tsx";
import { Button } from "@ui/components/ui/button.tsx";
import { useMutation } from "@tanstack/react-query";
import { login } from "../../services/auth.ts";
import { toast } from "@ui/components/ui/use-toast.ts";
import { AxiosError } from "axios";
import { useState } from "react";
import { onErrorToast, serverErrorToast } from "../ui/error-toast.tsx";
import { useUser } from "../../context/user-context.tsx";

export const LoginForm = () => {
  const [authError, setAuthError] = useState<boolean>(false);
  const { refresh } = useUser();
  const form = useForm<z.infer<typeof loginFormSchema>>({
    resolver: zodResolver(loginFormSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const mutation = useMutation({
    mutationFn: async (payload: z.infer<typeof loginFormSchema>) => {
      setAuthError(false);
      await login(payload);
    },
    onSuccess: () => {
      toast({
        description: "Login successful",
      });
      refresh();
    },
    onError: (e: Error) => {
      if (e instanceof AxiosError) {
        if (e.response?.status === 401 || e.response?.status == 404) {
          setAuthError(true);
        } else {
          onErrorToast(e);
        }
      } else {
        serverErrorToast();
      }
    },
  });

  return (
    <Form {...form}>
      <form
        className="w-1/2 px-4 flex flex-col gap-2"
        onSubmit={form.handleSubmit((values) => mutation.mutate(values))}
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
                <Input placeholder={"Password"} type={"password"} {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        {authError && (
          <p className={"mt-2 text-destructive text-sm"}>
            Invalid email/password
          </p>
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
  );
};
