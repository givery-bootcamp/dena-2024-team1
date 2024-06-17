import { createAsyncThunk } from "@reduxjs/toolkit";

import { client } from "~/shared/services/client";
import { Hello, Post } from "~/shared/models";


const API_ENDPOINT_PATH =
  import.meta.env.VITE_API_ENDPOINT_PATH ?? "";

export const getHello = createAsyncThunk<Hello>("getHello", async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/hello`);
  return await response.json();
});

export const getPosts = createAsyncThunk<Post[]>("getPosts", async () => {
  const response = await client.get("/posts");
  return response.data;
});
