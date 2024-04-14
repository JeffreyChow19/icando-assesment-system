import { cn } from "@ui/lib/utils";
import React from "react";
import Select from "react-select";

const Multiselect = React.forwardRef<
  React.ElementRef<typeof Select>,
  React.ComponentPropsWithoutRef<typeof Select>
>(({ className, ...props }, ref) => (
  <Select
    ref={ref}
    isMulti
    className={cn(
      "h-10 rounded-md border border-input",
      className,
    )}
    {...props}
  />
));

export { Multiselect };
