import { valibotResolver } from "@hookform/resolvers/valibot";
import { useCallback } from "react";
import { type SubmitHandler, useForm } from "react-hook-form";
import { Link } from "react-router-dom";

import { type Schema, schema } from "./schema";

type UserAuthFormProps = {
  onSubmit: (username: string, password: string) => void;
  title: string;
  submitText: string;
  linkText: string;
  linkTo: string;
}

export const UserAuthForm = ({ title, submitText, linkText, linkTo, onSubmit }: UserAuthFormProps) => {
  const { register, formState, handleSubmit: handleFormSubmit } = useForm<Schema>({
    resolver: valibotResolver(schema),
  });

  const handleSubmit: SubmitHandler<Schema> = useCallback((data) => {
    const { username, password } = data;
    onSubmit(username, password);
  }, [onSubmit]);

  return (
    <form className="mt-10 flex flex-col gap-7" onSubmit={handleFormSubmit(handleSubmit)}>
      <div className="flex flex-col gap-2">
        <h1 className="text-xl font-bold">{ title }</h1>
      </div>
      <hr className="border-border" />
      <div className="flex flex-col gap-3 text-lg">
        <label className="flex flex-col gap-2 text-lg">
          <span>ユーザー名</span>
          <input className="border border-border p-2" type="text" {...register("username")} />
          {formState.errors.username && <span className="text-sm text-alert">{formState.errors.username.message}</span>}
        </label>
        <label className="flex flex-col gap-2 text-lg">
          <span>パスワード</span>
          <input className="border border-border p-2" type="password" {...register("password")} />
          {formState.errors.password && <span className="text-sm text-alert">{formState.errors.password.message}</span>}
        </label>
        <button className="mt-10 rounded-md bg-primary p-2 text-white" type="submit">{submitText}</button>
      </div>
      <div>
        <Link to={linkTo} className="underline">{linkText}</Link>
      </div>
    </form>
  );
};