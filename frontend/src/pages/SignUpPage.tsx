import { useCallback } from "react";

import { Container } from "~/shared/components/Container";
import { UserAuthForm } from "~/features/users/UserAuthForm";
import { userApi } from "~/shared/services/API";

export const SignUpPage = () => {
  const handleSubmit = useCallback(async (username: string, password: string) => {
    try {
      const response = await userApi.signUp({
        user_name: username,
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
      <UserAuthForm type="sign-up" onSubmit={handleSubmit} />
    </Container>
  );
};
