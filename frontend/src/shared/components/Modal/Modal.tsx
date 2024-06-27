import { Children, type FC, type ReactNode } from "react";
import { X } from "react-feather";

type Props = {
  isOpen: boolean;
  children: ReactNode;
  onClose: () => void;
}

export const Modal: FC<Props> = ({
  isOpen,
  children,
  onClose,
}) => {
  const parsedChildren = Children.map(children, (child) => child);
  const [header, content, footer] = parsedChildren as [ReactNode, ReactNode, ReactNode];

  if (!isOpen) return null;
  return (
    <div className="absolute left-0 top-0 z-modal flex h-screen w-screen items-center justify-center bg-[black]/40">
      <div className="relative flex flex-col items-center justify-center gap-12 rounded-lg bg-white p-16 text-center">
        <button className="absolute right-2 top-2" onClick={onClose}>
          <X size={28} />
        </button>
        {header}
        {content}
        {footer}
      </div>
    </div>
  );
};
