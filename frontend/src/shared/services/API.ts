import { createAsyncThunk } from "@reduxjs/toolkit";

import { Post, PostApi, Sketch, SketchApi } from "~/generated";
import { Hello } from "~/shared/models";
import { config } from "~/config/api";

const API_ENDPOINT_PATH =
  import.meta.env.VITE_API_ENDPOINT_PATH ?? "";

export const postApi = new PostApi();
export const sketchApi = new SketchApi();

export const getHello = createAsyncThunk<Hello>("getHello", async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/hello`);
  return await response.json();
});

export const getPosts = createAsyncThunk<Post[]>("getPosts", async () => {
  const response = await postApi.getAllPosts(config);
  return response.data;
});

export const getPost = createAsyncThunk<Post, number>("getPost", async (postId) => {
  const response = await postApi.getPostById(postId, config);
  return response.data;
});

export const getSketches = createAsyncThunk<Sketch[]>("getSketches", async () => {
  const response = await sketchApi.getAllSketches(config);
  return response.data;
});
