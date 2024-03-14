import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger
} from "@ui/components/ui/dialog.tsx";
import {Button} from "@ui/components/ui/button.tsx";
import { z } from "zod"
import {Input} from "@ui/components/ui/input.tsx";
import {Question} from "../../../interfaces/question.ts";
import {EditIcon, XIcon} from "lucide-react";
import {useForm} from "react-hook-form";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from "@ui/components/ui/form.tsx";
import {zodResolver} from "@hookform/resolvers/zod";
import {Textarea} from "@ui/components/ui/textarea.tsx";
import {useState} from "react";
import {RadioGroup, RadioGroupItem} from "@ui/components/ui/radio-group.tsx";
import {cn} from "@ui/lib/utils.ts";

interface QuestionFormProps {
  type: 'edit' | 'new'
  question?: Question
}
export const QuestionForm = ({type, question}: QuestionFormProps) => {
  const [answerLength, setAnswerLength] = useState(question?.choices?.length || 1);
  const [correctAnswer, setCorrectAnswer] = useState(question?.answerId && question?.choices ? question.choices.findIndex((choice) => choice.id == question.answerId) : 0);
  const answerFormSchema = z.object({
    text: z.string({required_error: "Choice is required"}).min(1, {
      message: "Choice must be at least 1 character"
    })
  })

  const formSchema = z.object({
    text: z.string({required_error: "Question text is required"}).min(2, {
      message: "Question text must be at least 2 characters.",
    }),
    choices: z.array(answerFormSchema).min(1)
  })


  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      text: "",
      choices: [{
        text: "Option 1"
      }]
    },
  })

  const removeAnswer = (index: number) => {
    const currAnswer = form.getValues("choices")
     currAnswer.splice(index, 1)
    form.setValue("choices", currAnswer)
    if (index == correctAnswer) setCorrectAnswer(0)
    setAnswerLength(answerLength - 1)
  }

  const addAnswer = () => {
    const currAnswer = form.getValues("choices")
    currAnswer.push({
      text: `Option ${answerLength + 1}`
    })
    form.setValue("choices", currAnswer)
    setAnswerLength(answerLength + 1)
    console.log(form.getValues("choices"))
  }

  const answerElements = Array.from({ length: answerLength }, (_, index) => (
    <div key={index} className="grid grid-cols-1">
      <div>
        <FormField
          control={form.control}
          name={`choices.${index}.text`}
          render={({ field }) => (
            <FormItem>
              <FormLabel>{`Option ${index + 1}`} </FormLabel>
              <div className={cn( correctAnswer == index && "bg-green-100", "flex flex-row justify-between gap-2 items-center py-2 px-3 rounded-md")}>
                <RadioGroupItem value={index.toString()} checked={index == correctAnswer} className="mr-2"/>
                <FormControl>
                  <Input placeholder={`Option ${index + 1}`} {...field} />
                </FormControl>
                <Button
                  className="h-8 w-8 p-1"
                  variant="ghost"
                  onClick={(e) => {
                    e.preventDefault();
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
    </div>
  ));

  const handleSubmit = () => {

  }

  return (
    <Dialog>
      <DialogTrigger asChild>
        {
          type == 'new' ? <Button>New Question</Button> :
            <Button className="h-8 w-8 p-1" variant="secondary"><EditIcon size="12"/></Button>
        }
      </DialogTrigger>
      <DialogContent className="w-[80vw]">
        <Form {...form}>
          <form onSubmit={form.handleSubmit(handleSubmit)} className="space-y-8">
            <DialogHeader>
              <DialogTitle>{type == 'new' ? 'New Question' : 'Edit Question'}</DialogTitle>
            </DialogHeader>
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
            <div className="flex flex-col gap-2">
              <RadioGroup onValueChange={(val) => setCorrectAnswer(parseInt(val))}>
                {answerElements}
              </RadioGroup>
              <Button variant="ghost" className="text-primary p-2 w-fit" onClick={addAnswer}>
               + Add choice
              </Button>
            </div>
            <DialogFooter>
              <Button type="submit">Save</Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  )
}
