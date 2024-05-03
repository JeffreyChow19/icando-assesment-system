import { useQuery } from "@tanstack/react-query";
import { Button } from "@ui/components/ui/button";
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@ui/components/ui/collapsible";
import {
  Table,
  TableCaption,
  TableCell,
  TableRow,
} from "@ui/components/ui/table";
import { ChevronDown, ChevronRight, SearchIcon } from "lucide-react";
import { useEffect, useState } from "react";
import { getQuizHistory } from "../../services/quiz";
import { formatDate, formatHour } from "../../utils/format-date";
import { Badge } from "@ui/components/ui/badge";
import { Link } from "react-router-dom";

export const HistoryCollapsible = ({ quizId }: { quizId: string }) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);

  const { data, refetch } = useQuery({
    queryFn: () =>
      getQuizHistory({
        id: quizId,
        page: 1,
        limit: 10,
      }),
    queryKey: ["quizHistory", quizId],
    enabled: false,
  });

  useEffect(() => {
    if (isOpen) {
      refetch();
    }
  }, [isOpen, refetch]);

  return (
    <Collapsible
      open={isOpen}
      onOpenChange={setIsOpen}
      className="w-full space-y-2"
      asChild
    >
      <>
        <CollapsibleTrigger asChild>
          <div className="flex flex-row items-center hover:bg-accent hover:text-accent-foreground">
            <Button variant="ghost" size="sm" className="w-9 p-0">
              {isOpen ? (
                <ChevronDown className="h-4 w-4" />
              ) : (
                <ChevronRight className="h-4 w-4" />
              )}
              <span className="sr-only">Toggle</span>
            </Button>
            <p className="cursor-pointer underline text-gray-500 hover:text-gray-400">
              Show Versions
            </p>
          </div>
        </CollapsibleTrigger>
        <CollapsibleContent asChild>
          <Table>
            <TableCaption>
              {data && data.data.length === 0 && (
                <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
                  <SearchIcon className="w-10 h-10" />
                  No versions.
                </div>
              )}
            </TableCaption>
            {data?.data &&
              data.data.length > 0 &&
              data.data.map((version, index, array) => {
                const idx = array.length - 1 - index;
                return (
                  <TableRow>
                    <TableCell>
                      <strong>Version {idx + 1}</strong>
                    </TableCell>
                    <TableCell>
                      {version.lastPublishedAt ? (
                        <div className="flex flex-row">
                          <Badge
                            key={formatDate(new Date(version.lastPublishedAt))}
                            className="mr-2"
                            variant={"outline"}
                          >
                            {formatDate(new Date(version.lastPublishedAt))}
                          </Badge>
                          <Badge
                            key={formatHour(new Date(version.lastPublishedAt))}
                            variant={"outline"}
                          >
                            {formatHour(new Date(version.lastPublishedAt))}
                          </Badge>
                        </div>
                      ) : (
                        ""
                      )}
                    </TableCell>
                    <TableCell>
                      <Button size={"sm"}>
                        <Link to={`/quiz/${quizId}`}>View Details</Link>
                      </Button>
                    </TableCell>
                  </TableRow>
                );
              })}
          </Table>
        </CollapsibleContent>
      </>
    </Collapsible>
  );
};
