import { Link } from "react-router-dom";

export function Header() {
  return (
    <header className="bg-primary text-white flex justify-between items-center h-20 px-20 font-bold text-xl">
      <p className="text-white">サンプルアプリケーション</p>
      <Link to="/">Top</Link>
    </header>
  )
}
