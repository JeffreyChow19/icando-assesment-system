import { toast } from '@ui/components/ui/use-toast.ts';
import { AxiosError } from 'axios';

export const serverErrorToast = () => {
  toast({
    variant: 'destructive',
    description: '[500] Unexpected Error Occured',
  });
};

export const onErrorToast = (err: Error) => {
  if (err instanceof AxiosError) {
    toast({
      variant: 'destructive',
      description: err.response?.data.error,
    });
    return;
  }
  serverErrorToast();
};
