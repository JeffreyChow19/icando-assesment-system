import { Link, useParams } from "react-router-dom";
import SidebarIcon from "../../public/ic_sidebar.svg";
import { useStudentQuiz } from "../context/user-context.tsx";
import { Sheet, SheetContent } from "@ui/components/ui/sheet.tsx";
import { Dispatch, FC, SetStateAction } from "react";

export const Navigation = ({
  toggleSidebar,
  showNavigation,
}: {
  pageTitle: string;
  toggleSidebar: () => void;
  showNavigation: boolean;
}) => {
  return (
    <header className="relative w-full flex flex-row justify-between px-2 lg:px-6 py-4 z-20 items-center">
      <img src={"/logo.png"} alt={"logo"} className="bg-white w-48 m-auto rounded-md" />

      {showNavigation &&
        <button
          onClick={toggleSidebar}
          className="absolute right-5 top-0 bottom-0"
        >
          <img src={SidebarIcon} alt="Sidebar Icon" className="h-6 w-6 " />
        </button>
      }
    </header>
  );
};

interface SideBarProps {
  sidebarOpen: boolean;
  setSidebarOpen: Dispatch<SetStateAction<boolean>>;
}

export const SideBar: FC<SideBarProps> = ({ sidebarOpen, setSidebarOpen }) => {
  const { questions, studentAnswers } = useStudentQuiz();
  const { number } = useParams();
  const currNumber = number ? parseInt(number) : 1;

  const CURRENT_NUMBER_STYLE = "bg-primary text-white";
  const ANSWERED_NUMBER_STYLE = "bg-[#facc13] text-black";
  const DEFAULT_NUMBER_STYLE = "bg-[#D9D9D9] text-back";

  const questionNumbers = questions?.map((question, index) => {
    const answered = studentAnswers?.some(
      (answer) => answer.questionId === question.id,
    );
    const num = index + 1;

    return {
      number: num,
      style:
        num == currNumber
          ? CURRENT_NUMBER_STYLE
          : answered
            ? ANSWERED_NUMBER_STYLE
            : DEFAULT_NUMBER_STYLE,
    };
  });

  return (
    <Sheet open={sidebarOpen} onOpenChange={() => setSidebarOpen(!sidebarOpen)}>
      <SheetContent>
        <div className="h-auto w-fit py-6 px-3 overflow-auto grid grid-cols-2 gap-6">
          {questionNumbers?.map((question) => (
            <Link
              to={`/quiz/${question.number}`}
              key={question.number}
              className={`rounded-lg text-center w-12 h-12 flex items-center justify-center ${question.style}`}
            >
              {question.number}
            </Link>
          ))}
        </div>
      </SheetContent>
    </Sheet>
  );
};
