import { valibotResolver } from "@hookform/resolvers/valibot";
import { useCallback } from "react";
import { type SubmitHandler, useForm } from "react-hook-form";
import { Link } from "react-router-dom";

import { type Schema, schema } from "./schema";

import { Button } from "~/shared/components/Button";

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
      <div className="flex flex-col gap-6 text-lg">
        <label className="flex flex-col gap-2 text-lg">
          <span className="font-bold">ユーザー名</span>
          <div className="flex flex-col gap-1">
            <input className="border border-border p-2" type="text" {...register("username")} />
            {formState.errors.username && <span className="text-sm font-semibold text-alert">{formState.errors.username.message}</span>}
          </div>
        </label>
        <label className="flex flex-col gap-2 text-lg">
          <span className="font-bold">パスワード</span>
          <div className="flex flex-col gap-1">
            <input className="border border-border p-2" type="password" {...register("password")} />
            {formState.errors.password && <span className="text-sm font-semibold text-alert">{formState.errors.password.message}</span>}
          </div>
        </label>
        <Button type="submit">{submitText}</Button>
      </div>
      <div>
        <Link to={linkTo} className="underline">{linkText}</Link>
      </div>
    </form>
  );
};
