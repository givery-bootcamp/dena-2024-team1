import type { FC, ReactNode } from "react";

type Props = {
  children: ReactNode
}

export const ModalFooter: FC<Props> = ({ children }) => {
  return (
    <div className="flex w-full flex-col gap-2">
      {children}
    </div>
  );
};
