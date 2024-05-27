import { Link } from "react-router-dom";

export function Header() {
  return (
    <header className="bg-primary">
      <p>サンプルアプリケーション</p>
      <Link to="/">Top</Link>
    </header>
  )
}
