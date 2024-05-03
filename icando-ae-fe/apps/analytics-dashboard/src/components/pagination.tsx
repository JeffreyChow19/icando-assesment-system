import React, { useEffect } from "react";
import { Button } from "@repo/ui/components/ui/button";
import {
  ChevronLeftIcon,
  ChevronRightIcon,
  ChevronsLeft,
  ChevronsRight,
} from "lucide-react";
import { useSearchParams } from "react-router-dom";

interface PaginationProps {
  page: number;
  totalPage: number;
  setPage: React.Dispatch<React.SetStateAction<number>>;
  withSearchParams?: boolean;
}

export const Pagination = ({
  page,
  totalPage,
  setPage,
  withSearchParams,
}: PaginationProps) => {
  const [searchParams, setSearchParams] = useSearchParams();

  useEffect(() => {
    if (withSearchParams) {
      searchParams.set("page", page.toString());
      setSearchParams(searchParams);
    }
  }, [page, searchParams, setSearchParams, withSearchParams]);

  return (
    <div className="flex flex-row justify-end w-full">
      <div className="flex w-[100px] items-center justify-center text-sm font-medium">
        Page {page} of {totalPage}
      </div>
      <div className="flex items-center space-x-2">
        <Button
          variant="outline"
          className="hidden h-8 w-8 p-0 lg:flex"
          onClick={() => setPage(1)}
          disabled={page <= 1}
        >
          <ChevronsLeft className="h-4 w-4" />
        </Button>
        <Button
          variant="outline"
          className="h-8 w-8 p-0"
          onClick={() => setPage(page - 1)}
          disabled={page <= 1}
        >
          <ChevronLeftIcon className="h-4 w-4" />
        </Button>
        <Button
          variant="outline"
          className="h-8 w-8 p-0"
          onClick={() => setPage(page + 1)}
          disabled={page >= totalPage}
        >
          <ChevronRightIcon className="h-4 w-4" />
        </Button>
        <Button
          variant="outline"
          className="hidden h-8 w-8 p-0 lg:flex"
          onClick={() => setPage(totalPage)}
          disabled={page >= totalPage}
        >
          <ChevronsRight className="h-4 w-4" />
        </Button>
      </div>
    </div>
  );
};
