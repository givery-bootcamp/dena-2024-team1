import { useCallback, useEffect } from "react";

import { UserApi } from "~/generated";
import { config } from "~/config/api";

export const SignOutPage = () => {

  const signOut = useCallback(async () => {
    const userApi = new UserApi();
    const response = await userApi.signOut(config);

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
