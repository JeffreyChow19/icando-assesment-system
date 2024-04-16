import { useForm } from "react-hook-form";
import { quizPublishFormSchema } from "./quiz-schema";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useToast } from "@ui/components/ui/use-toast";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { getAllClasses } from "../../services/classes";
import { publishQuiz, PublishQuizPayload } from "../../services/quiz";
import { onErrorToast } from "../ui/error-toast";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@ui/components/ui/form.tsx";
import { Input } from "@ui/components/ui/input";
import { Multiselect } from "@ui/components/ui/multiselect";
import { Button } from "@ui/components/ui/button";
import { Loader2 } from "lucide-react";
import { useNavigate } from "react-router-dom";

export const QuizPublishForm = ({ id }: { id: string }) => {
  const form = useForm<z.infer<typeof quizPublishFormSchema>>({
    resolver: zodResolver(quizPublishFormSchema),
    defaultValues: {
      quizId: id,
    },
  });
  const { toast } = useToast();
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  const { data: classData } = useQuery({
    queryKey: ["classes"],
    queryFn: () => getAllClasses(),
  });

  const mutation = useMutation({
    mutationFn: async (payload: PublishQuizPayload) => {
      console.log(payload);
      return await publishQuiz(payload);
    },
    onSuccess: (response) => {
      console.log(response);
      queryClient.invalidateQueries({ queryKey: ["quiz", id] });
      toast({
        description: "Quiz published successfully",
      });
      navigate(-1);
    },
    onError: (err: Error) => {
      onErrorToast(err);
    },
  });

  return (
    <Form {...form}>
      <form
        className="py-4 space-y-4 lg:w-2/3"
        onSubmit={form.handleSubmit((values) =>
          mutation.mutate({
            quizId: id,
            quizDuration: values.quizDuration,
            startDate:
              values.startAt instanceof Date
                ? values.startAt.toISOString()
                : values.startAt,
            endDate:
              values.endAt instanceof Date
                ? values.endAt.toISOString()
                : values.endAt,
            assignedClasses: values.assignedClasses,
          }),
        )}
      >
        <FormField
          control={form.control}
          name="quizDuration"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Duration*</FormLabel>
              <FormControl>
                <Input
                  type="number"
                  min="0"
                  placeholder="Duration"
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="startAt"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Quiz Start Date*</FormLabel>
              <FormControl>
                <Input
                  type="datetime-local"
                  placeholder="Start Date"
                  {...field}
                  value={
                    field.value instanceof Date
                      ? field.value.toISOString()
                      : field.value
                  }
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="endAt"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Quiz End Date*</FormLabel>
              <FormControl>
                <Input
                  type="datetime-local"
                  placeholder="End Date"
                  {...field}
                  value={
                    field.value instanceof Date
                      ? field.value.toISOString()
                      : field.value
                  }
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="assignedClasses"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Publish to Classes*</FormLabel>
              <FormControl>
                <Multiselect
                  onChange={(val) => {
                    console.log(
                      val.map((item: { value: string }) => item.value),
                    );

                    field.onChange(
                      val.map((item: { value: string }) => item.value),
                    );
                  }}
                  options={
                    classData
                      ? classData.data.map((item) => {
                          return {
                            label: `${item.name} - ${item.grade}`,
                            value: item.id,
                          };
                        })
                      : []
                  }
                />
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
  );
};
