import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { Toaster } from "@repo/ui/components/ui/toaster";
import { UserProvider } from "./context/user-context.tsx";
import localizedFormat from "dayjs/plugin/localizedFormat.js";
import dayjs from "dayjs";
import { AlertDialogProvider } from "./context/alert-dialog.tsx";
import { HelmetProvider } from "react-helmet-async";
import { Dashboard } from "./pages/dashboard.tsx";
import "dayjs/locale/id";
import { TooltipProvider } from "@repo/ui/components/ui/tooltip";
import { Login } from "./pages/login.tsx";
import { StudentQuizReview } from './pages/quizzes/review';
import { StudentStatistics } from "./pages/students/student";
import { QuizDetail } from "./pages/quizzes/quiz-detail.tsx";
import { Quizzes } from './pages/quizzes';
import { Students } from './pages/students';

dayjs.locale("id");
dayjs.extend(localizedFormat);

function App() {
  const queryClient = new QueryClient();

  const router = createBrowserRouter([
    {
      path: "/",
      element: <Dashboard />,
    },
    {
      path: "/login",
      element: <Login />,
    },
    {
      path: "/quiz",
      element: <Quizzes />,
    },
    {
      path: "/quiz/:quizid/review/:studentquizid",
      element: <StudentQuizReview />,
    },
    {
      path: "/student",
      element: <Students />,
    },
    {
      path: "/student/:studentId",
      element: <StudentStatistics />,
    },
    {
      path: "/quiz/:id",
      element: <QuizDetail />,
    },
  ]);
  return (
    <QueryClientProvider client={queryClient}>
      <AlertDialogProvider>
        <UserProvider>
          <HelmetProvider>
            <TooltipProvider>
              <RouterProvider router={router} />
            </TooltipProvider>
          </HelmetProvider>
        </UserProvider>
        <Toaster />
      </AlertDialogProvider>
    </QueryClientProvider>
  );
}

export default App;
