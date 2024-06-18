import { useCallback, useEffect, useState } from "react";

import { AppRoute, AppRouteNoAuth } from "./AppRoute";

import { sessionUserSlice  } from "~/shared/store/SessionUserSlice";
import { UserApi } from "~/generated";
import { useAppDispatch, useAppSelector } from "~/shared/hooks";
import { Header, HeaderNoAuth } from "~/shared/components/Header";

export const AuthProvider = () => {
  const [loading, setLoading] = useState(false);

  const { sessionUser } = useAppSelector((state) => state.sessionUser);
  const dispatch = useAppDispatch();  

  const getSessionUser = useCallback(async () => {
    setLoading(true);

    const userApi = new UserApi();
    const response = await userApi.getSessionUser({
      withCredentials: true,
    });

    // 取得に成功した場合
    if (response.status === 200) {
        const user = response.data;

        console.log(user);

        // ユーザー情報をセット
        dispatch(sessionUserSlice.actions.setSessionUser(user));
    }
  }, []);

  useEffect(() => {
    if (sessionUser) {
      return;
    }

    getSessionUser()
    .then(() => {})
    .catch(() => {})
    .finally(() => {
      setLoading(false);
    });
  }, [sessionUser]);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (sessionUser) {
    return (
      <>
        <Header/>
        <main>
          <AppRoute />
        </main>
      </>
    );
  }

  return (
    <>
      <HeaderNoAuth/>
      <main>
        <AppRouteNoAuth />
      </main>
    </>
  );
};
