import { createSlice } from "@reduxjs/toolkit";

import { Post } from "~/generated";
import { APIService } from "~/shared/services";

export type PostState = {
  post?: Post;
};

export const initialState: PostState = {};

export const postSlice = createSlice({
  name: "post",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(APIService.getPost.fulfilled, (state, action) => {
      state.post = action.payload;
    });
  },
});

export default postSlice.reducer;
