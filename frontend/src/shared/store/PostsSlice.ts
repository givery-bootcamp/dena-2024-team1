import { createSlice } from "@reduxjs/toolkit";

import { Post } from "~/generated";
import { APIService } from "~/shared/services";

export type PostsState = {
  posts?: Post[];
};

export const initialState: PostsState = {};

export const postsSlice = createSlice({
  name: "posts",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(APIService.getPosts.fulfilled, (state, action) => {
      state.posts = action.payload;
    });
  },
});

export default postsSlice.reducer;
