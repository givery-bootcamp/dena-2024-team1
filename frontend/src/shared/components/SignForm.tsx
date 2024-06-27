import { valibotResolver } from "@hookform/resolvers/valibot";
import { useCallback } from "react";
import { type SubmitHandler, useForm } from "react-hook-form";
import { Link } from "react-router-dom";
import * as v from "valibot";

const schema = v.object({
  username: v.pipe(
    v.string(),
    v.nonEmpty("ユーザー名を入力してください"),
  ),
  password: v.pipe(
    v.string(),
    v.nonEmpty("パスワードを入力してください"),
    v.minLength(8, "パスワードは8文字以上で入力してください"),
  ),
});
type Schema = v.InferOutput<typeof schema>

type SignFormProps = {
  onSubmit: (username: string, password: string) => void;
  title: string;
  submitText: string;
  linkText: string;
  linkTo: string;
}

export const SignForm = (props: SignFormProps) => {
  const { register, handleSubmit: handleFormSubmit } = useForm<Schema>({
    resolver: valibotResolver(schema),
  });
  const { onSubmit } = props; 

  const handleSubmit: SubmitHandler<Schema> = useCallback((data) => {
    const { username, password } = data;
    onSubmit(username, password);
  }, [onSubmit]);

  const { title, submitText, linkText, linkTo } = props;

  return (
    <form className="mt-10 flex flex-col gap-7" onSubmit={handleFormSubmit(handleSubmit)}>
      <div className="flex flex-col gap-2">
        <h1 className="text-xl font-bold">{ title }</h1>
      </div>
      <hr className="border-border" />
      <div className="flex flex-col gap-3 text-lg">
        <label className="text-lg" htmlFor="username">ユーザー名</label>
        <input className="border border-border p-2" type="text" id="username" {...register("username")} />
        <label className="text-lg" htmlFor="password">パスワード</label>
        <input className="border border-border p-2" type="password" id="password" {...register("password")} />
        <button className="mt-10 rounded-md bg-primary p-2 text-white" type="submit">{submitText}</button>
      </div>
      <div>
        <Link to={linkTo} className="underline">{linkText}</Link>
      </div>
    </form>
  );
};
