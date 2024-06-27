import { useCallback, useEffect } from "react";

import { userApi } from "~/shared/services/API";

export const SignOutPage = () => {

  const signOut = useCallback(async () => {
    const response = await userApi.signOut();

    if (response.status === 200) {
        console.log("Success");

        // 画面を再読み込みする
        window.location.href = "/";
    }
  }, []);


  useEffect(() => {
    signOut()
      .then(() => {})
      .catch(() => {});
  }, []);

  return null;
};
