import { createAsyncThunk } from "@reduxjs/toolkit";

import { Post, PostApi } from "~/generated";
import { Hello } from "~/shared/models";

const API_ENDPOINT_PATH =
  import.meta.env.VITE_API_ENDPOINT_PATH ?? "";

export const getHello = createAsyncThunk<Hello>("getHello", async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/hello`);
  return await response.json();
});

export const getPosts = createAsyncThunk<Post[]>("getPosts", async () => {
  const postApi = new PostApi();
  const response = await postApi.getAllPosts();
  return response.data;
});
