import type { ReactNode } from "react"

type Props = {
  children: ReactNode,
  className?: string
}

export function Container({ children, className }: Props) {
  const containerClassName = "max-w-[1020px] mx-auto ";
  return <div className={containerClassName + (className === undefined ? "" : className)}>{children}</div>
}
