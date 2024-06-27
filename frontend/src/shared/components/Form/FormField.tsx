import{ Children, type FC, type ReactNode } from "react";

type Props = {
  children: ReactNode
}

export const FormField: FC<Props> = ({ children }) => {
  const parsedChildren = Children.map(children, (child) => child);
  const [label, input, errorMessage] = parsedChildren as [ReactNode, ReactNode, ReactNode];

  return (
    <label className="flex flex-col gap-2 text-lg">
      {label}
      <div className="flex flex-col gap-1">
        {input}
        {errorMessage}
      </div>
    </label>
  );
};
