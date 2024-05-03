import { ReactNode } from "react";
import { Card } from "@ui/components/ui/card.tsx";

interface StatsCardProps {
  children: ReactNode;
  className?: string;
}

export const StatsCard: React.FC<StatsCardProps> = ({
  children,
  className,
}) => {
  return (
    <Card
      className={`flex flex-col items-center gap-4 rounded-3xl p-4 ${className}`}
    >
      {children}
    </Card>
  );
};
