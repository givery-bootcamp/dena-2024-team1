import { useCallback, FormEvent } from "react";
import { Link } from "react-router-dom";

type SignFormProps = {
  onSubmit: (username: string, password: string) => void;
  title: string;
  submitText: string;
  linkText: string;
  linkTo: string;
}


export const SignForm = (props: SignFormProps) => {
  const { onSubmit } = props; 

  const handleSubmit = useCallback(async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const form = event.currentTarget;
    const formData = new FormData(form);
    const username = formData.get("username") as string;
    const password = formData.get("password") as string;

    onSubmit(username, password);
  }, [onSubmit]);

  const { title, submitText, linkText, linkTo } = props;

  return (
    <form className="mt-10 flex flex-col gap-7" onSubmit={handleSubmit}>
      <div className="flex flex-col gap-2">
        <h1 className="text-xl font-bold">{ title }</h1>
      </div>
      <hr className="border-border" />
      <div className="flex flex-col gap-3 text-lg">
        <label className="text-lg" htmlFor="username">ユーザー名</label>
        <input className="border border-border p-2" type="text" id="username" name="username" />
        <label className="text-lg" htmlFor="password">パスワード</label>
        <input className="border border-border p-2" type="password" id="password" name="password" />
        <button className="mt-10 rounded-md bg-primary p-2 text-white" type="submit">{submitText}</button>
      </div>
      <div>
        <Link to={linkTo} className="underline">{linkText}</Link>
      </div>
    </form>
  );
};
