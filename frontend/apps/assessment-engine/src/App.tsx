import "@repo/ui/globals.css";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { Toaster } from "@repo/ui/components/ui/toaster";
import { UserProvider } from "./context/user-context";
import localizedFormat from "dayjs/plugin/localizedFormat";
import dayjs from "dayjs";
import { AlertDialogProvider } from "./context/alert-dialog";
import { HelmetProvider } from "react-helmet-async";
import { Hello } from "./pages/hello";
import "dayjs/locale/id";
import { TooltipProvider } from "@repo/ui/components/ui/tooltip";

dayjs.locale("id");
dayjs.extend(localizedFormat);

function App() {
  const queryClient = new QueryClient();

  const router = createBrowserRouter([
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
