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
    <div>
      <div>{icon}</div>
      <p>{title}</p>
      <p>{subTitlte}</p>
      <Button onClick={onPrimaryButtonClick}>{primaryButtonLabel}</Button>
      <Button variant="secondary" onClick={onSecondaryButtonClick}>{secondaryButtonLabel}</Button>
    </div>
  );
};
