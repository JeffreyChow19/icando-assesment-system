import { QuestionList } from "./question-list.tsx";
import { QuizDetail } from "../../../interfaces/quiz.ts";
import { QuizInfo } from "./quiz-info.tsx";
import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from "@ui/components/ui/tabs.tsx";

export const QuizForm = ({ quiz }: { quiz: QuizDetail }) => {
  return (
    <div>
      <div className="mb-4">
        <h1 className="text-center text-xl font-bold">{quiz.name || "Untitled Quiz"}</h1>
      </div>
      <Tabs defaultValue="information" className="w-full">
        <TabsList>
          <TabsTrigger value="information">Information</TabsTrigger>
          <TabsTrigger value="questions">Questions</TabsTrigger>
        </TabsList>
        <TabsContent value="information">
          <QuizInfo quiz={quiz} />
        </TabsContent>
        <TabsContent value="questions">
          <QuestionList questions={quiz.questions || []} />
        </TabsContent>
      </Tabs>
    </div>
  );
};