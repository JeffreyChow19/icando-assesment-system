import { Link } from "react-router-dom";
import SidebarIcon from "../../public/ic_sidebar.svg";

export const Navigation = ({
  pageTitle,
  toggleSidebar,
  showNavigation,
}: {
  pageTitle: string;
    toggleSidebar: () => void;
    showNavigation: boolean;
}) => {
  return (
    <header className="relative w-full flex flex-row justify-between px-2 lg:px-6 py-4 z-20 items-center">
      <h1 className="text-center flex-grow text-white font-semibold text-2xl">
        {pageTitle}
      </h1>
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

export const SideBar = ({ sidebarOpen }: { sidebarOpen: boolean }) => {
  // TODO: get actual datas
  const currNumber = 1;
  const numOfQuestions = 5;

  const questionNumbers = Array.from(
    { length: numOfQuestions },
    (_, i) => i + 1,
  );

  return (
    <div
      className={`absolute top-0 right-0 h-full w-fit bg-white rounded-l-xl transition-transform duration-300 ease-in-out ${sidebarOpen ? "transform translate-x-0" : "transform translate-x-full"}`}
    >
      <div className="h-auto p-6 overflow-auto grid grid-cols-2 gap-6">
        {questionNumbers.map((number) => (
          <Link
            to={`/question/${number}`}
            key={number}
            className={`rounded-lg text-center w-12 h-12 flex items-center justify-center ${number === currNumber ? "bg-primary text-white" : "bg-[#D9D9D9] text-black"}`}
          >
            {number}
          </Link>
        ))}
      </div>
    </div>
  );
};
