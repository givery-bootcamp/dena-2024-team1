import type { FC, ReactNode } from "react";

type Props = {
  children: ReactNode
}

export const FormErrorMessage: FC<Props> = ({ children }) => {
  return <p className="text-sm font-semibold text-alert">{children}</p>;
};
