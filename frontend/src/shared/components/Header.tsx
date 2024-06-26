import { Link } from "react-router-dom";

export function Header() {
  return (
    <header className="flex h-20 items-center justify-between bg-primary px-12 text-xl font-bold text-white">
      <p className="text-white">サンプルアプリケーション</p>
      <div className="flex space-x-4">
        <Link to="/" className="text-white">Top</Link>
        <Link to="/posts/create" className="text-white">新規投稿</Link>
        <Link to="/sketches/create" className="text-white">Sketch</Link>
        <Link to="/signout" className="text-white">Signout</Link>
      </div>
    </header>
  );
}

export function HeaderNoAuth() {
  return (
    <header className="flex h-20 items-center justify-between bg-primary px-12 text-xl font-bold text-white">
      <p className="text-white">サンプルアプリケーション</p>
    </header>
  );
}
