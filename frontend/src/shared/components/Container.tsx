import type { ReactNode } from "react"

type Props = {
  children: ReactNode
}

export function Container({ children }: Props) {
  return <div className="max-w-[1020px] mx-auto">{children}</div>
}
