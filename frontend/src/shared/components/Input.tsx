import classNames from "classnames";
import { forwardRef, type InputHTMLAttributes } from "react";

type Props = InputHTMLAttributes<HTMLInputElement>

export const Input = forwardRef<HTMLInputElement, Props>(({ className, ...props }, ref) => {
    return <input ref={ref} className={classNames("border rounded border-border p-2 w-full", className)} {...props} />;
  },
);

Input.displayName = "Input";
