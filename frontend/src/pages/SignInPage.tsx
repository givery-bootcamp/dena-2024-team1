import { useCallback, FormEvent } from "react";

import { Container } from "~/shared/components/Container";
import { UserApi } from "~/generated";


export const SignInPage = () => {
  const handleSubmit = useCallback(async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const form = event.currentTarget;
    const formData = new FormData(form);
    const username = formData.get("username") as string;
    const password = formData.get("password") as string;

    const userApi = new UserApi();
    const responnse = await userApi.signIn({
      username,
      password,
    }, {
      withCredentials: true,
    });

    if (responnse.status === 200) {
      console.log("Success");

      // 画面を再読み込みする
      window.location.href = "/";
    }

  }, []);

  return (
    <Container>
      <form className="mt-10 flex flex-col gap-7" onSubmit={handleSubmit}>
        <div className="flex flex-col gap-2">
          <h1 className="text-xl font-bold">サインイン</h1>
        </div>
        <hr className="border-border" />
        <div className="flex flex-col gap-3 text-lg">
          <label className="text-lg" htmlFor="username">ユーザーネーム</label>
          <input className="border border-border p-2" type="text" id="username" name="username" />
          <label className="text-lg" htmlFor="password">パスワード</label>
          <input className="border border-border p-2" type="password" id="password" name="password" />
          <button className="rounded-md bg-primary p-2 text-white" type="submit">サインイン</button>
        </div>
      </form>
    </Container>
  );
};
