import { useCallback } from "react";

import { Container } from "~/shared/components/Container";
import { SignForm } from "~/shared/components/SignForm";
import { userApi } from "~/shared/services/API";

export const SignUpPage = () => {
  const handleSubmit = useCallback(async (username: string, password: string) => {
    try {
      const response = await userApi.signUp({
        username,
        password,
      });
  
      if (response.status === 200) {
        window.location.href = "/sign-in";
        
        return;
      }
    } catch (error) {
      alert("サインアップに失敗しました。ユーザー名が既に使用されている可能性があります");
    }
    
  }, []);

  return (
    <Container>
      <SignForm
        onSubmit={handleSubmit}
        title="新規登録"
        submitText="アカウントを作成"
        linkText="アカウントをお持ちの方はこちら"
        linkTo="/signin"
      />
    </Container>
  );
};
