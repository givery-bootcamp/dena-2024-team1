import { configureStore } from "@reduxjs/toolkit";


import helloReducer, { helloSlice } from './HelloSlice';
import postsReducer, { postsSlice } from './PostsSlice';


export const store = configureStore({
  reducer: {
    hello: helloReducer,
    posts: postsReducer,
  },
});

export const actions = {
  ...helloSlice.actions,
  ...postsSlice.actions,
};

export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
