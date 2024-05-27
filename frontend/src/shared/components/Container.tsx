import type { ReactNode } from "react"

type Props = {
  children: ReactNode
}

export function Container({ children }: Props) {
  return <p className="max-w-[1020px] mx-auto">{children}</p>
}
