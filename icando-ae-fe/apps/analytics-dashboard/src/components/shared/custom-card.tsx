import { ReactNode } from "react";
import { Card } from "@ui/components/ui/card.tsx";

export const CustomCard = ({ children }: { children: ReactNode }) => {
  return <Card className="border-none shadow-sm rounded-lg">{children}</Card>;
};
