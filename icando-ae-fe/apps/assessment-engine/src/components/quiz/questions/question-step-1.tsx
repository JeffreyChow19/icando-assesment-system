import { Button } from "@ui/components/ui/button";
import { DialogFooter } from "@ui/components/ui/dialog";
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@ui/components/ui/form";
import { Input } from "@ui/components/ui/input";
import { RadioGroup, RadioGroupItem } from "@ui/components/ui/radio-group";
import { ScrollArea } from "@ui/components/ui/scroll-area";
import { Textarea } from "@ui/components/ui/textarea";
import { cn } from "@ui/lib/utils";
import { useFormContext } from "react-hook-form";
import { questionFormSchema } from "./question-schema";
import { useMemo } from "react";
import { z } from "zod";
import { XIcon } from "lucide-react";

interface QuestionStep1Props {
  next: () => void;
}

export const QuestionStep1 = ({ next }: QuestionStep1Props) => {
  const form = useFormContext<z.infer<typeof questionFormSchema>>();

  const choicesState = form.watch("choices");
  const answerLength = choicesState.length;

  // choices is always
  const largestId = useMemo(() => {
    return Math.max(...choicesState.map((c) => c.id));
  }, [choicesState]);

  const removeAnswer = (index: number) => {
    const currAnswer = form.getValues("choices");
    const deletedAnswer = currAnswer[index];
    const answerId = form.getValues("answerId");
    currAnswer.splice(index, 1);
    form.setValue("choices", currAnswer);
    if (deletedAnswer.id == answerId) {
      form.setValue("answerId", choicesState[0].id);
    }
  };

  const addAnswer = () => {
    const currAnswer = form.getValues("choices");
    currAnswer.push({
      text: "",
      id: largestId + 1,
    });
    form.setValue("choices", currAnswer);
  };

  return (
    <>
      <ScrollArea className="h-[70vh] px-4">
        <div className="mx-1">
          <FormField
            control={form.control}
            name="text"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Question Text</FormLabel>
                <FormControl>
                  <Textarea placeholder="Question Text" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <div className="flex flex-col gap-2 mt-4">
          <RadioGroup
            onValueChange={(val) => form.setValue("answerId", parseInt(val))}
          >
            {choicesState.map((choice, index) => {
              return (
                <div key={index} className="grid grid-cols-1">
                  <FormField
                    control={form.control}
                    name={`choices.${index}.text`}
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>{`Option ${index + 1}`} </FormLabel>
                        <div
                          className={cn(
                            form.watch("answerId") == choice.id &&
                              "bg-green-100",
                            "flex flex-row justify-between gap-2 items-center py-2 px-3 rounded-md",
                          )}
                        >
                          <RadioGroupItem
                            value={choice.id.toString()}
                            checked={choice.id == form.watch("answerId")}
                            className="mr-2"
                          />
                          <FormControl>
                            <Input
                              placeholder={`Option ${index + 1}`}
                              {...field}
                            />
                          </FormControl>
                          <Button
                            className="h-8 w-8 p-1"
                            variant="ghost"
                            type="button"
                            onClick={() => {
                              removeAnswer(index);
                            }}
                            disabled={answerLength === 1}
                          >
                            <XIcon size="12" />
                          </Button>
                        </div>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>
              );
            })}
          </RadioGroup>
          <Button
            variant="ghost"
            type="button"
            className="text-primary p-2 w-fit"
            onClick={() => {
              addAnswer();
            }}
            disabled={answerLength === 6}
          >
            + Add choice
          </Button>
        </div>
      </ScrollArea>
      <DialogFooter>
        <Button variant="outline" type="button" onClick={() => next()}>
          Next
        </Button>
      </DialogFooter>
    </>
  );
};
