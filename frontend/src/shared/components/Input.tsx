import { forwardRef, type InputHTMLAttributes } from "react";

type Props = InputHTMLAttributes<HTMLInputElement>

export const Input = forwardRef<HTMLInputElement, Props>((props, ref) => {
    return <input ref={ref} className="border border-border p-2" {...props} />;
  },
);

Input.displayName = "Input";
