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
import { Check, ChevronsUpDown, Loader2 } from "lucide-react";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@ui/components/ui/popover.tsx";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "@ui/components/ui/command.tsx";
import { SUBJECTS } from "../../utils/constants.ts";
import { cn } from "@ui/lib/utils.ts";
import { Badge } from "@ui/components/ui/badge.tsx";

export const QuizInfo = ({ isPending }: { isPending: boolean }) => {
  const form = useFormContext<z.infer<typeof quizFormSchema>>();
  const orderedSubjects = SUBJECTS.sort((a, b) => a.localeCompare(b));
  const isSubjectSelected = (subject: string) => {
    const selectedSubjects = form.getValues("subject");
    return selectedSubjects.includes(subject);
  };

  const unselectSubject = (subject: string) => {
    form.setValue(
      "subject",
      form.getValues("subject").filter((s) => s !== subject),
    );
  };

  const selectSubject = (subject: string) => {
    const newSubjects = [...form.getValues("subject"), subject];
    form.setValue("subject", newSubjects);
  };

  const onSubjectSelected = (subject: string) => {
    if (isSubjectSelected(subject)) {
      unselectSubject(subject);
      return;
    }
    selectSubject(subject);
  };

  return (
    <div className={"flex flex-col gap-6 max-w-[500px]"}>
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
          <FormItem className="flex flex-col gap-2">
            <FormLabel>Quiz Subject</FormLabel>
            <Popover>
              <PopoverTrigger asChild>
                <FormControl>
                  <Button
                    variant="outline"
                    role="combobox"
                    className={cn(
                      "w-[300px] justify-between",
                      !field.value && "text-muted-foreground",
                    )}
                  >
                    Select subject(s)
                    <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                  </Button>
                </FormControl>
              </PopoverTrigger>
              <PopoverContent className="w-[300px] p-0">
                <Command>
                  <CommandInput placeholder="Search language..." />
                  <CommandEmpty>No language found.</CommandEmpty>
                  <CommandGroup>
                    {orderedSubjects.map((subject) => (
                      <CommandItem
                        value={subject}
                        key={subject}
                        onSelect={() => onSubjectSelected(subject)}
                      >
                        <Check
                          className={cn(
                            "mr-2 h-4 w-4",
                            isSubjectSelected(subject)
                              ? "opacity-100"
                              : "opacity-0",
                          )}
                        />
                        {subject}
                      </CommandItem>
                    ))}
                  </CommandGroup>
                </Command>
              </PopoverContent>
            </Popover>
            <FormMessage />
            {form.watch("subject") && form.watch("subject").length > 0 && (
              <div className="flex flex-wrap gap-2">
                {form.watch("subject").map((subject) => (
                  <Badge key={subject}>{subject}</Badge>
                ))}
              </div>
            )}
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
        <Button type="submit" size={"lg"} disabled={isPending}>
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
