import { Link } from "react-router-dom";

export function Header() {
  return (
    <header className="flex h-20 items-center justify-between bg-primary px-12 text-xl font-bold text-white">
      <p className="text-white">お絵描きの森</p>
      <div className="flex space-x-4">
        <Link to="/" className="text-white">みんなの絵</Link>
        <Link to="/sketches/create" className="text-white">お絵描きする</Link>
        <Link to="/signout" className="text-white">ログアウト</Link>
      </div>
    </header>
  );
}

export function HeaderNoAuth() {
  return (
    <header className="flex h-20 items-center justify-between bg-primary px-12 text-xl font-bold text-white">
      <p className="text-white">お絵描きの森</p>
      <div className="flex space-x-4">
        <Link to="/" className="text-white">みんなの絵</Link>
        <Link to="/sketches/create" className="text-white">お絵描きする</Link>
        <Link to="/signin" className="text-white">ログイン</Link>
        <Link to="/signup" className="text-white">新規登録</Link>
      </div>
    </header>
  );
}
