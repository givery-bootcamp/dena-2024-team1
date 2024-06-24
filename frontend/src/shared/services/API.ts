import { createAsyncThunk } from "@reduxjs/toolkit";

import { CreatePostRequest, CreatePostResponse, Post, PostApi } from "~/generated";
import { Hello } from "~/shared/models";

const API_ENDPOINT_PATH =
  import.meta.env.VITE_API_ENDPOINT_PATH ?? "";
const postApi = new PostApi();

export const getHello = createAsyncThunk<Hello>("getHello", async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/hello`);
  return await response.json();
});

export const getPosts = createAsyncThunk<Post[]>("getPosts", async () => {
  const response = await postApi.getAllPosts();
  return response.data;
});

export const getPost = createAsyncThunk<Post, number>("getPost", async (postId) => {
  const response = await postApi.getPostById(postId);
  return response.data;
});

export const createPost = createAsyncThunk<CreatePostResponse, CreatePostRequest>("createPost", async (createPostRequest) => {
  const response = await postApi.postPost(createPostRequest);
  return response.data;
});
