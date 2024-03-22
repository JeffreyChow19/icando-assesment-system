import { Badge } from "@ui/components/ui/badge.tsx";

interface CompetencyBadgeProps {
  competency: {
    id: string;
    name: string;
    numbering: string;
  };
}
export const CompetencyBadge = ({ competency }: CompetencyBadgeProps) => {
  return (
    <Badge key={competency.id} variant="outline">
      {competency.numbering} - {competency.name}
    </Badge>
  );
};
