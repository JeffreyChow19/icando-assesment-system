import { zodResolver } from '@hookform/resolvers/zod';
import { useMutation } from '@tanstack/react-query';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { Loader2 } from 'lucide-react';
import { useNavigate } from 'react-router-dom';
import { Student } from '../../interfaces/student.ts';
import { useToast } from '@ui/components/ui/use-toast.ts';
import { createStudent, CreateStudentPayload, updateStudent, UpdateStudentPayload } from '../../services/student.ts';
import { onErrorToast } from '../ui/error-toast.tsx';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@ui/components/ui/form.tsx';
import { Input } from '@ui/components/ui/input.tsx';
import { Button } from '@ui/components/ui/button.tsx';

const studentFormSchema = z.object({
  firstName: z.string({ required_error: "First name can't be empty" }).min(1),
  lastName: z.string({ required_error: "Last name can't be empty" }).min(1),
  nisn: z.string({ required_error: "NISN can't be empty" }).min(2),
  email: z.string({ required_error: "Email can't be empty" }).email({ message: 'Email is invalid' }),
});

export const StudentsForm = ({ student }: { student?: Student }) => {
  const navigator = useNavigate();

  const form = useForm<z.infer<typeof studentFormSchema>>({
    resolver: zodResolver(studentFormSchema),
    defaultValues: student
      ? {
        firstName: student.firstName,
        lastName: student.lastName,
        nisn: student.nisn,
        email: student.email,
      }
      : {},
  });
  const { toast } = useToast();

  const mutation = useMutation({
    mutationFn: (payload: CreateStudentPayload) => {
      if (student) {
        const updatePayload: UpdateStudentPayload = {
          firstName: payload.firstName,
          lastName: payload.lastName,
        }
        return updateStudent(updatePayload, student.id);
      }
      return createStudent(payload);
    },
    onSuccess: () => {
      toast({
        description: `Student successfully ${student ? 'saved' : 'created'}!`,
      });
      navigator('/students');
    },
    onError: err => {
      onErrorToast(err);
    },
  });
  return (
    <>
      <Form {...form}>
        <form className="py-4 space-y-4 lg:w-2/3" onSubmit={form.handleSubmit(values => mutation.mutate(values))}>
          <FormField
            control={form.control}
            name="firstName"
            render={({ field }) => (
              <FormItem>
                <FormLabel>First Name*</FormLabel>
                <FormControl>
                  <Input placeholder="First name" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="lastName"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Last Name*</FormLabel>
                <FormControl>
                  <Input placeholder="Last name" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="nisn"
            render={({ field }) => (
              <FormItem>
                <FormLabel>NISN*</FormLabel>
                <FormControl>
                  <Input placeholder="NISN" {...field} disabled={!!student} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email*</FormLabel>
                <FormControl>
                  <Input placeholder="Competency description" {...field} disabled={!!student} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <div className="flex w-full justify-end">
            <Button type="submit" disabled={mutation.isPending}>
              {mutation.isPending ? (
                <>
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  Please wait
                </>
              ) : (
                'Submit'
              )}
            </Button>
          </div>
        </form>
      </Form>
    </>
  );
};
