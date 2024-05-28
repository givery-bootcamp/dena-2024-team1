import type { ReactNode } from "react";
import classNames from "classnames";

type Props = {
  children: ReactNode,
  className?: string
}

export function Container({ children, className }: Props) {
  const containerClassName = "max-w-[1020px] mx-auto ";
  return <div className={classNames(containerClassName, className)}>{children}</div>;
}
