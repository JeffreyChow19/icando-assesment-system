import { Loader2 } from "lucide-react";

export function LoadingComponent() {
  return (
    <div className="flex gap-x-2 min-h-full w-full items-center justify-center text-muted-foreground">
      <Loader2 className="h-20 w-20 animate-spin" />
      <h2>Loading...</h2>
    </div>
  );
}
