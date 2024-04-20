import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { Toaster } from "@repo/ui/components/ui/toaster";
import { UserProvider } from "./context/user-context.tsx";
import localizedFormat from "dayjs/plugin/localizedFormat.js";
import dayjs from "dayjs";
import { AlertDialogProvider } from "./context/alert-dialog.tsx";
import { HelmetProvider } from "react-helmet-async";
import "dayjs/locale/id";
import { TooltipProvider } from "@repo/ui/components/ui/tooltip";
import { Home } from "./pages";
import { LoginPage } from "./pages/login.tsx";
import { Classes } from "./pages/classes";
import { ClassNew } from "./pages/classes/class-new.tsx";
import { ClassEdit } from "./pages/classes/class-edit.tsx";
import { ClassDetailPage } from "./pages/classes/class-detail.tsx";
import { UpdateQuizPage } from "./pages/quiz/update.tsx";
import { StudentsEdit } from "./pages/students/students-edit.tsx";
import { StudentsNew } from "./pages/students/students-new.tsx";
import { Students } from "./pages/students";
import { Quizzes } from "./pages/quizzes/index.tsx";
import { PublishQuizPage } from "./pages/quiz/publish.tsx";
import { ViewQuizPage } from "./pages/quiz/view.tsx";

dayjs.locale("id");
dayjs.extend(localizedFormat);

function App() {
  const queryClient = new QueryClient();

  const router = createBrowserRouter([
    {
      path: "/",
      element: <Home />,
    },
    {
      path: "/login",
      element: <LoginPage />,
    },
    {
      path: "/students",
      element: <Students />,
    },
    {
      path: "/students/new",
      element: <StudentsNew />,
    },
    {
      path: "/students/edit/:id",
      element: <StudentsEdit />,
    },
    {
      path: "/classes",
      element: <Classes />,
    },
    {
      path: "/classes/new",
      element: <ClassNew />,
    },
    {
      path: "/classes/edit/:id",
      element: <ClassEdit />,
    },
    {
      path: "/classes/:id",
      element: <ClassDetailPage />,
    },
    {
      path: "/quiz",
      element: <Quizzes />,
    },
    {
      path: "/quiz/:id/edit",
      element: <UpdateQuizPage />,
    },
    {
      path: "/quiz/:id/publish",
      element: <PublishQuizPage />,
    },
    {
      path: "/history/:id",
      element: <ViewQuizPage />,
    }
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
