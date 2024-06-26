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
import { Input } from "~/shared/components/Input";
import {
  FormErrorMessage,
  FormField,
  FormLabel,
} from "~/shared/components/Form";

type UserAuthFormProps = {
  type: PageType;
  onSubmit: (username: string, password: string) => void;
}

export const UserAuthForm = ({ type, onSubmit }: UserAuthFormProps) => {
  const {
    register,
    formState: { errors },
    handleSubmit: handleFormSubmit,
  } = useForm<Schema>({
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
        <FormField>
          <FormLabel>ユーザー名</FormLabel>
          <Input {...register("username")} />
          {errors.username && (
            <FormErrorMessage>{errors.username.message}</FormErrorMessage>
          )}
        </FormField>
        <FormField>
          <FormLabel>パスワード</FormLabel>
          <Input type="password" {...register("password")} />
          {errors.password && (
            <FormErrorMessage>{errors.password.message}</FormErrorMessage>
          )}
        </FormField>
        <Button type="submit">{buttonLabel}</Button>
      </div>
      <div>
        <Link to={link} className="underline">{navigationText}</Link>
      </div>
    </form>
  );
};
