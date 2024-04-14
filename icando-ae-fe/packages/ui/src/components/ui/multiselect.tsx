import { cn } from "@ui/lib/utils";
import React, { ForwardedRef } from "react";
import Select, { SelectInstance } from "react-select";

export type Options = {
  label: string;
  value: string;
};

export type MultiselectProps = {
  options: Options[];
  defaultValue: Options[];
  ref?: ForwardedRef<SelectInstance>;
  // eslint-disable-next-line no-unused-vars
  onChange: (selectedItem: Options[]) => void;
  className?: string | undefined;
};

const Multiselect = React.forwardRef<SelectInstance, MultiselectProps>(
  ({ options, defaultValue, onChange, className, ...props }, ref) => (
    <Select
      options={options}
      ref={ref}
      isMulti
      defaultValue={defaultValue}
      onChange={(value: unknown) => onChange(value as Options[])}
      className={cn("h-10 rounded-md border border-input", className)}
      {...props}
    />
  ),
);

export { Multiselect };
