import { useCallback } from "react";

import { Container } from "~/shared/components/Container";
import { UserAuthForm } from "~/features/users/UserAuthForm";
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
      <UserAuthForm type="sign-in" onSubmit={handleSubmit} />
    </Container>
  );
};
