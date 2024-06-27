import type { FC, ReactNode } from "react";

type Props = {
  children: ReactNode
}

export const FormLabel: FC<Props> = ({ children }) => {
  return <span className="font-bold">{children}</span>;
};
