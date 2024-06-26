import { useCallback, useEffect, useState } from "react";

import { AppRoute, AppRouteNoAuth } from "./AppRoute";

import { sessionUserSlice  } from "~/shared/store/SessionUserSlice";
import { useAppDispatch, useAppSelector } from "~/shared/hooks";
import { Header, HeaderNoAuth } from "~/shared/components/Header";
import { userApi } from "~/shared/services/API";

export const AuthProvider = () => {
  const [loading, setLoading] = useState(true);

  const { sessionUser } = useAppSelector((state) => state.sessionUser);
  const dispatch = useAppDispatch();  

  const getSessionUser = useCallback(async () => {
    setLoading(true);

    try {
      const response = await userApi.getSessionUser();
  
      // 取得に成功した場合
      if (response.status === 200) {
          const user = response.data.user;
  
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
