import { useFormContext } from "react-hook-form";
import { z } from "zod";
import { quizFormSchema } from "./quiz-schema";
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@ui/components/ui/form.tsx";
import { Input } from "@ui/components/ui/input.tsx";
import { Button } from "@ui/components/ui/button";
import { Loader2 } from "lucide-react";

export const QuizInfo = ({ isPending }: { isPending: boolean }) => {
  const form = useFormContext<z.infer<typeof quizFormSchema>>();

  return (
    <div className={"flex flex-col gap-4 max-w-[500px]"}>
      <FormField
        control={form.control}
        name={"name"}
        render={({ field }) => (
          <FormItem>
            <FormLabel>Quiz Name</FormLabel>
            <FormControl>
              <Input {...field} placeholder={"Enter quiz name"} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )}
      />
      <FormField
        control={form.control}
        name={"subject"}
        render={({ field }) => (
          <FormItem>
            <FormLabel>Quiz Subject</FormLabel>
            <FormControl>
              <Input {...field} placeholder={"Enter quiz subject"} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )}
      />
      <FormField
        control={form.control}
        name={"passingGrade"}
        render={({ field }) => (
          <FormItem>
            <FormLabel>Passing Grade</FormLabel>
            <FormDescription>
              Passing grade value should be between 0 and 100
            </FormDescription>
            <FormControl>
              <Input
                type="number"
                min="0"
                max="100"
                {...field}
                placeholder={"Enter quiz passing grade"}
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        )}
      />
      <div className="flex w-full justify-end">
        <Button type="submit" form="quiz" size={"lg"} disabled={isPending}>
          {isPending ? (
            <>
              <Loader2 className="mr-2 h-4 w-4 animate-spin" />
              Updating ...
            </>
          ) : (
            <>Update</>
          )}
        </Button>
      </div>
    </div>
  );
};
