import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { Toaster } from "@repo/ui/components/ui/toaster";
import { StudentProvider, QuizProvider } from "./context/user-context.tsx";
import localizedFormat from "dayjs/plugin/localizedFormat.js";
import dayjs from "dayjs";
import { AlertDialogProvider } from "./context/alert-dialog.tsx";
import { HelmetProvider } from "react-helmet-async";
import "dayjs/locale/id";
import { TooltipProvider } from "@repo/ui/components/ui/tooltip";
import { Home } from "./pages/home/index.tsx";
import { Quiz } from "./pages/quiz.tsx";
import Join from "./pages/join.tsx";

dayjs.locale("id");
dayjs.extend(localizedFormat);

function App() {
  const queryClient = new QueryClient();

  const router = createBrowserRouter([
    {
      path: "/quiz",
      element: <Quiz />,
    },
    {
      path: "/",
      element: <Home />,
    }, {
      path: "/join",
      element: <Join />,
    }
  ]);
  return (
    <QueryClientProvider client={queryClient}>
      <AlertDialogProvider>
        <QuizProvider>
          <StudentProvider>
            <HelmetProvider>
              <TooltipProvider>
                <RouterProvider router={router} />
              </TooltipProvider>
            </HelmetProvider>
          </StudentProvider>
        </QuizProvider>
        <Toaster />
      </AlertDialogProvider>
    </QueryClientProvider>
  );
}

export default App;
