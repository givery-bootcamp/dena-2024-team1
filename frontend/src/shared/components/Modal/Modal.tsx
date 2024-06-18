import { Children, type FC, type ReactNode } from "react";

type Props = {
  isOpen: boolean;
  children: ReactNode;
}

export const Modal: FC<Props> = ({
  isOpen,
  children,
}) => {
  const parsedChildren = Children.map(children, (child) => child);
  const [header, content, footer] = parsedChildren as [ReactNode, ReactNode, ReactNode];

  if (!isOpen) return null;
  return (
    <div className="absolute left-0 top-0 flex h-screen w-screen items-center justify-center bg-[black]/40">
      <div className="flex flex-col items-center justify-center gap-12 rounded-lg bg-white p-16 text-center">
        {header}
        {content}
        {footer}
      </div>
    </div>
  );
};
