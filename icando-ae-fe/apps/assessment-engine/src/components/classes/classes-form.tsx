import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Loader2 } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { useToast } from "@ui/components/ui/use-toast.ts";
import { onErrorToast } from "../ui/error-toast.tsx";
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
import { Class } from "../../interfaces/classes.ts";
import {
  createClass,
  CreateClassPayload,
  updateClass,
  UpdateClassPayload,
} from "../../services/classes.ts";

const SemicolonSeparatedIDsSchema = z
  .string({ required_error: "Teacher ID can't be empty" })
  .refine(
    (value) => {
      const regex = /^(?:[^,]+(?:,[^,]+)*)?$/; // Regex to match comma-separated strings
      return regex.test(value);
    },
    {
      message: "Teacher IDs must be separable by comma.",
    },
  );

const classFormSchema = z.object({
  name: z.string({ required_error: "Class name can't be empty" }).min(1),
  grade: z.string({ required_error: "Class grade can't be empty" }).min(1),
  teacherIds: SemicolonSeparatedIDsSchema,
});

export const ClassesForm = ({ classes }: { classes?: Class }) => {
  const navigator = useNavigate();

  console.log(classes);

  const form = useForm<z.infer<typeof classFormSchema>>({
    resolver: zodResolver(classFormSchema),
    defaultValues: classes
      ? {
          name: classes.name,
          grade: classes.grade,
          teacherIds: classes.teachers!.map((teacher) => teacher.id).join(";"),
          // todo: refine selection of teacherId
        }
      : {},
  });
  const { toast } = useToast();

  const mutation = useMutation({
    mutationFn: (payload: CreateClassPayload) => {
      if (classes) {
        const updatePayload: UpdateClassPayload = {
          name: payload.name,
          grade: payload.grade,
          teacherIds: payload.teacherIds,
        };
        return updateClass(updatePayload, classes.id);
      }
      return createClass(payload);
    },
    onSuccess: () => {
      toast({
        description: `Class successfully ${classes ? "saved" : "created"}!`,
      });
      navigator("/classes");
    },
    onError: (err) => {
      onErrorToast(err);
    },
  });
  return (
    <>
      <Form {...form}>
        <form
          className="py-4 space-y-4 lg:w-2/3"
          onSubmit={form.handleSubmit((values) =>
            mutation.mutate({
              name: values.name,
              grade: values.grade,
              teacherIds: values.teacherIds.split(";"),
            }),
          )}
        >
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Class name*</FormLabel>
                <FormControl>
                  <Input placeholder="Name" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="grade"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Grade*</FormLabel>
                <FormControl>
                  <Input placeholder="Grade" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="teacherIds"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Teacher ID*</FormLabel>
                <FormControl>
                  <Input placeholder="Teacher IDs" {...field} />
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
                "Submit"
              )}
            </Button>
          </div>
        </form>
      </Form>
    </>
  );
};
