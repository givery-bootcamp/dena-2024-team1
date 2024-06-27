import { createAsyncThunk } from "@reduxjs/toolkit";

import { Post, PostApi, Sketch, SketchApi, UserApi } from "~/generated";
import { Hello } from "~/shared/models";
import { API_ENDPOINT_PATH, axiosInstance, config } from "~/config/api";

export const postApi = new PostApi(config, API_ENDPOINT_PATH, axiosInstance);
export const sketchApi = new SketchApi(config, API_ENDPOINT_PATH, axiosInstance);
export const userApi = new UserApi(config, API_ENDPOINT_PATH, axiosInstance);

export const getHello = createAsyncThunk<Hello>("getHello", async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/hello`, { credentials: "include" });
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

export const getSketches = createAsyncThunk<Sketch[]>("getSketches", async () => {
  const response = await sketchApi.getAllSketches();
  return response.data;
});
