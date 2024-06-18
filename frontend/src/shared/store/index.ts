import { configureStore } from "@reduxjs/toolkit";

import helloReducer, { helloSlice } from "./HelloSlice";
import postsReducer, { postsSlice } from "./PostsSlice";
import postReducer, { postSlice } from "./PostSlice";


export const store = configureStore({
  reducer: {
    hello: helloReducer,
    posts: postsReducer,
    post: postReducer,
  },
});

export const actions = {
  ...helloSlice.actions,
  ...postsSlice.actions,
  ...postSlice.actions,
};

export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
