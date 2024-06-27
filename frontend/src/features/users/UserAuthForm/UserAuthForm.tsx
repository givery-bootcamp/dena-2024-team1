import { valibotResolver } from "@hookform/resolvers/valibot";
import { useCallback } from "react";
import { type SubmitHandler, useForm } from "react-hook-form";
import { Link } from "react-router-dom";

import { type Schema, schema } from "./schema";
import {
  type PageType,
  typeToButtonLabel,
  typeToLink,
  typeToNavigationText,
  typeToTitle,
} from "./mapping";

import { Button } from "~/shared/components/Button";

type UserAuthFormProps = {
  type: PageType;
  onSubmit: (username: string, password: string) => void;
}

export const UserAuthForm = ({ type, onSubmit }: UserAuthFormProps) => {
  const { register, formState, handleSubmit: handleFormSubmit } = useForm<Schema>({
    resolver: valibotResolver(schema),
  });

  const title = typeToTitle[type];
  const buttonLabel = typeToButtonLabel[type];
  const link = typeToLink[type];
  const navigationText = typeToNavigationText[type];

  const handleSubmit: SubmitHandler<Schema> = useCallback((data) => {
    const { username, password } = data;
    onSubmit(username, password);
  }, [onSubmit]);

  return (
    <form className="mt-10 flex flex-col gap-10" onSubmit={handleFormSubmit(handleSubmit)}>
      <h1 className="text-center text-2xl font-bold">{title}</h1>
      <div className="flex flex-col gap-6 text-lg">
        <label className="flex flex-col gap-2 text-lg">
          <span className="font-bold">ユーザー名</span>
          <div className="flex flex-col gap-1">
            <input className="border border-border p-2" type="text" {...register("username")} />
            {formState.errors.username && (
              <span className="text-sm font-semibold text-alert">
                {formState.errors.username.message}
              </span>
            )}
          </div>
        </label>
        <label className="flex flex-col gap-2 text-lg">
          <span className="font-bold">パスワード</span>
          <div className="flex flex-col gap-1">
            <input className="border border-border p-2" type="password" {...register("password")} />
            {formState.errors.password && (
              <span className="text-sm font-semibold text-alert">
                {formState.errors.password.message}
              </span>
            )}
          </div>
        </label>
        <Button type="submit">{buttonLabel}</Button>
      </div>
      <div>
        <Link to={link} className="underline">{navigationText}</Link>
      </div>
    </form>
  );
};
