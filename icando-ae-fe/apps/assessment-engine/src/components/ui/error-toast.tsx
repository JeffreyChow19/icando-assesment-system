import { toast } from "@ui/components/ui/use-toast.ts";
import { AxiosError } from "axios";

export const serverErrorToast = () => {
  toast({
    variant: "destructive",
    description: "[500] Unexpected Error Occured",
  });
};

export const onErrorToast = (err: Error) => {
  if (err instanceof AxiosError && err.response) {
    const response = err.response!;

    let message;

    if (response.data.error) {
      message = response.data.error;
    } else if (response.data.errors) {
      message = response.data.errors;
    } else if (response.data.message) {
      message = response.data.message;
    } else if (response.data.messages) {
      message = response.data.messages;
    } else {
      message = response.data;
    }

    toast({
      variant: "destructive",
      description: message,
    });
    return;
  }
  serverErrorToast();
};
