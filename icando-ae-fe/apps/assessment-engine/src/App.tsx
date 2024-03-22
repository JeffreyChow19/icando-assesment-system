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
import { NewQuizPage } from "./pages/quiz/new.tsx";
import { StudentsEdit } from "./pages/students/students-edit.tsx";
import { StudentsNew } from "./pages/students/students-new.tsx";
import { Students } from "./pages/students";
import { Classes } from "./pages/classes";
import { ClassNew } from "./pages/classes/class-new.tsx";
import { ClassEdit } from "./pages/classes/class-edit.tsx";

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
      path: "/quiz/new",
      element: <NewQuizPage />,
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
