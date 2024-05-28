import { Link } from "react-router-dom";

export function Header() {
  return (
    <header className="flex h-20 items-center justify-between bg-primary px-12 text-xl font-bold text-white">
      <p className="text-white">サンプルアプリケーション</p>
      <Link to="/">Top</Link>
    </header>
  );
}
