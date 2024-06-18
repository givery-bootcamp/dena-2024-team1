import { configureStore } from "@reduxjs/toolkit";

import helloReducer, { helloSlice } from "./HelloSlice";
import postsReducer, { postsSlice } from "./PostsSlice";
import sessionUserReducer, { sessionUserSlice } from "./SessionUserSlice";


export const store = configureStore({
  reducer: {
    hello: helloReducer,
    posts: postsReducer,
    sessionUser: sessionUserReducer,
  },
});

export const actions = {
  ...helloSlice.actions,
  ...postsSlice.actions,
  ...sessionUserSlice.actions,
};

export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
