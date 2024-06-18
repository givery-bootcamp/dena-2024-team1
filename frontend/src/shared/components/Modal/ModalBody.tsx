import type { FC, ReactNode } from "react";

type Props = {
  children: ReactNode
}

export const ModalBody: FC<Props> = ({ children }) => {
  return <p className="text-sm">{children}</p>;
};
