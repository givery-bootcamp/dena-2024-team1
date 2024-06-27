import { useCallback } from "react";

import { Container } from "~/shared/components/Container";
import { UserApi } from "~/generated";
import { config } from "~/config/api";
import { UserAuthForm } from "~/features/users/UserAuthForm";

export const SignUpPage = () => {
  const handleSubmit = useCallback(async (username: string, password: string) => {
    const userApi = new UserApi();
    
    try {
      const response = await userApi.signUp({
        username,
        password,
      }, config);
  
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
      <UserAuthForm
        onSubmit={handleSubmit}
        title="新規登録"
        submitText="アカウントを作成"
        linkText="アカウントをお持ちの方はこちら"
        linkTo="/signin"
      />
    </Container>
  );
};
