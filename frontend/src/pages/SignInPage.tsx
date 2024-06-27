import { useCallback } from "react";

import { Container } from "~/shared/components/Container";
import { SignForm } from "~/shared/components/SignForm";
import { userApi } from "~/shared/services/API";


export const SignInPage = () => {
  const handleSubmit = useCallback(async (username: string, password: string) => {
    try {
      const responnse = await userApi.signIn({
        username,
        password,
      });
  
      if (responnse.status === 200) {
        console.log("Success");
  
        // 画面を再読み込みする
        window.location.href = "/";
      }
    } catch (error) {
      alert("サインインに失敗しました。ユーザー名かパスワードが間違っています");
    }

  }, []);

  return (
    <Container>
      <SignForm
        onSubmit={handleSubmit}
        title="ログイン"
        submitText="ログイン"
        linkText="アカウントをお持ちでない方はこちら"
        linkTo="/signup"
      />
    </Container>
  );
};
