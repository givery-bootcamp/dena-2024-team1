import { useCallback } from "react";

import { Container } from "~/shared/components/Container";
import { UserApi } from "~/generated";
import { SignForm } from "~/shared/components/SignForm";
import { config } from "~/config/api";


export const SignInPage = () => {
  const handleSubmit = useCallback(async (username: string, password: string) => {
    const userApi = new UserApi();
    
    try {
      const responnse = await userApi.signIn({
        username,
        password,
      }, config);
  
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
