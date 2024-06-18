import type { FC, ReactNode } from "react";

import { Button } from "~/shared/components/Button";

type Props = {
  isOpen: boolean;
  icon: ReactNode;
  title: string;
  subTitlte: string;
  primaryButtonLabel: string;
  secondaryButtonLabel: string;
  onPrimaryButtonClick: () => void;
  onSecondaryButtonClick: () => void;
}

export const Modal: FC<Props> = ({
  isOpen,
  icon,
  title,
  subTitlte,
  primaryButtonLabel,
  secondaryButtonLabel,
  onPrimaryButtonClick,
  onSecondaryButtonClick,
}) => {
  if (!isOpen) return null;
  return (
    <div className="absolute left-0 top-0 flex h-screen w-screen items-center justify-center bg-[black]/40">
      <div className="flex flex-col items-center justify-center gap-12 rounded-lg bg-white p-16 text-center">
        <div className="flex flex-col items-center gap-6">
          <div>{icon}</div>
          <p className="text-2xl font-bold">{title}</p>
        </div>
        <p className="text-sm">{subTitlte}</p>
        <div className="flex w-full flex-col gap-2">
          <Button onClick={onPrimaryButtonClick}>{primaryButtonLabel}</Button>
          <Button variant="secondary" onClick={onSecondaryButtonClick}>{secondaryButtonLabel}</Button>
        </div>
      </div>
    </div>
  );
};
