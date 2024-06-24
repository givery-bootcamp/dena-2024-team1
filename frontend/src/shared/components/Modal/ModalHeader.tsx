import type { FC, ReactNode } from "react";

type Props = {
  children: ReactNode
}

export const ModalHeader: FC<Props> = ({ children }) => {
  return <p className="text-2xl font-bold">{children}</p>;
};
