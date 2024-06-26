import { useCallback, useEffect, useState } from "react";

import { AppRoute, AppRouteNoAuth } from "./AppRoute";

import { sessionUserSlice  } from "~/shared/store/SessionUserSlice";
import { UserApi } from "~/generated";
import { useAppDispatch, useAppSelector } from "~/shared/hooks";
import { Header, HeaderNoAuth } from "~/shared/components/Header";

export const AuthProvider = () => {
  const [loading, setLoading] = useState(true);

  const { sessionUser } = useAppSelector((state) => state.sessionUser);
  const dispatch = useAppDispatch();  

  const getSessionUser = useCallback(async () => {
    setLoading(true);

    const userApi = new UserApi();
    try {
      const response = await userApi.getSessionUser({
        withCredentials: true,
      });
  
      // 取得に成功した場合
      if (response.status === 200) {
          const user = response.data;
  
          // ユーザー情報をセット
          dispatch(sessionUserSlice.actions.setSessionUser(user));
      }
    } catch (error) {
      return;
    }
    
  }, []);

  useEffect(() => {
    if (sessionUser) {
      return;
    }

    getSessionUser()
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
