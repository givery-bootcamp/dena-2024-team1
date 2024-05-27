import type { ReactNode } from "react"

type Props = {
  children: ReactNode
}

export function Container({ children }: Props) {
  return <p className="max-w-[1220px] mx-auto">{children}</p>
}
