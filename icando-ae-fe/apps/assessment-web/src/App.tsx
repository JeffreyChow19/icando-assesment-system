import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { Toaster } from "@repo/ui/components/ui/toaster";
import { UserProvider } from "./context/user-context.tsx";
import localizedFormat from "dayjs/plugin/localizedFormat.js";
import dayjs from "dayjs";
import { AlertDialogProvider } from "./context/alert-dialog.tsx";
import { HelmetProvider } from "react-helmet-async";
import { Hello } from "./pages/hello.tsx";
import "dayjs/locale/id";
import { TooltipProvider } from "@repo/ui/components/ui/tooltip";
import { TestPage } from "./pages/test.tsx";

dayjs.locale("id");
dayjs.extend(localizedFormat);

function App() {
  const queryClient = new QueryClient();

  const router = createBrowserRouter([
    {
      path: "/hello",
      element: <Hello />,
    },
    {
      path: "/",
      element: <Hello />,
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
