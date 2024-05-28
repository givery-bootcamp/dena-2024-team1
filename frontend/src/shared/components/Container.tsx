import type { ReactNode } from "react";

type Props = {
  children: ReactNode
}

export function Container({ children }: Props) {
  return <div className="mx-auto max-w-[1020px]">{children}</div>;
}
