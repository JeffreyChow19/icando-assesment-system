import { useQuery } from "@tanstack/react-query";
import { Button } from "@ui/components/ui/button";
import { DialogFooter } from "@ui/components/ui/dialog";
import { useState } from "react";
import { getAllCompetency } from "../../../services/competency";
import { z } from "zod";
import { useFormContext } from "react-hook-form";
import { questionFormSchema } from "./question-schema";
import { Competency } from "../../../interfaces/competency";
import { FormLabel } from "@ui/components/ui/form";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@ui/components/ui/table";
import { SearchIcon } from "lucide-react";
import { LoadingComponent } from "../../../components/loading";
import { Checkbox } from "@ui/components/ui/checkbox";
import { CompetencyBadge } from "./competency-badge";
import { Pagination } from "../../../components/pagination";

interface QuestionStep2Props {
  prev: () => void;
}

export const QuestionStep2 = ({ prev }: QuestionStep2Props) => {
  const form = useFormContext<z.infer<typeof questionFormSchema>>();
  const [page, setPage] = useState<number>(1);

  const { data } = useQuery({
    queryKey: ["competency", page],
    queryFn: () => getAllCompetency({ page, limit: 8 }),
    retry: false,
  });

  const competencies = data?.competencies;
  const meta = data?.meta;

  const checkCompetency = (id: string) => {
    return form.watch("competencies").findIndex((competency) => {
      return competency.id == id;
    });
  };

  const onSelectCompetency = (
    { id, name, numbering }: Competency,
    checked: boolean | "indeterminate",
  ) => {
    if (checked === "indeterminate") return;
    const newCompetencies = [...form.getValues("competencies")];

    if (checked) {
      newCompetencies.push({
        id,
        name,
        numbering,
      });
    } else {
      const index = checkCompetency(id);
      if (index > -1) {
        newCompetencies.splice(index, 1);
      }
    }
    form.setValue("competencies", newCompetencies);
  };

  return (
    <>
      <div className="flex flex-col gap-2 px-4 h-[70vh]">
        <FormLabel>Competencies</FormLabel>
        <Table>
          <TableCaption>
            {competencies ? (
              competencies.length <= 0 ? (
                <div className="flex flex-col w-full items-center justify-center gap-2 text-muted-foreground text-md">
                  <SearchIcon className="w-10 h-10" />
                  No questions yet for this quiz
                </div>
              ) : (
                meta &&
                meta.totalPage > 1 && (
                  <div className="flex w-full justify-end">
                    <Pagination
                      page={page}
                      totalPage={meta.totalPage}
                      setPage={setPage}
                    />
                  </div>
                )
              )
            ) : (
              <LoadingComponent />
            )}
          </TableCaption>
          <TableHeader>
            <TableRow>
              <TableHead className="w-[4vw]" />
              <TableHead className="w-[4vw]">#</TableHead>
              <TableHead>Competency</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {competencies?.map((competency) => (
              <TableRow key={competency.id}>
                <TableCell>
                  <Checkbox
                    checked={checkCompetency(competency.id) > -1}
                    onCheckedChange={(checked) => {
                      onSelectCompetency(competency, checked);
                    }}
                  />
                </TableCell>
                <TableCell>{competency.numbering}</TableCell>
                <TableCell>{competency.name}</TableCell>
                <TableCell>{competency.description}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        <div className="flex w-full gap-2 flex-wrap mt-2">
          {form.watch("competencies").map((competency) => (
            <CompetencyBadge competency={competency} key={competency.id} />
          ))}
        </div>
      </div>
      <DialogFooter>
        <div className="w-full flex gap-4 justify-end">
          <Button type="button" variant="outline" onClick={() => prev()}>
            Back
          </Button>
          <Button type="submit" form="question">
            Save
          </Button>
        </div>
      </DialogFooter>
    </>
  );
};
