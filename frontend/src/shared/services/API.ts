import { createAsyncThunk } from "@reduxjs/toolkit";

import { client } from "~/shared/services/client";
import { Hello, Post } from "~/shared/models";

/**
 * 1. とりあえず今の状態でコミットする
 * 2. myAxiosの定義と設定部分を別ファイルに切り出す
 * 3. anyは仕方ないので、2のファイル全体でeslintを無視する
 */

const API_ENDPOINT_PATH =
  import.meta.env.VITE_API_ENDPOINT_PATH ?? "";

export const getHello = createAsyncThunk<Hello>("getHello", async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/hello`);
  return await response.json();
});

export const getPosts = createAsyncThunk<Post[]>("getPosts", async () => {
  const response = await client.get("/posts");
  return response.data;
  //return await response.json();
  /*
  const posts: Post[] = [{
    title: "バナナはおやつに含まれますか？",
    userName: "Tadashi",
    updatedAt: "5/27 14:23",
    id: "1",
  },
  {
    title: "りんごはおやつに含まれますか？",
    userName: "Tadashi",
    updatedAt: "5/27 14:23",
    id: "2",
  },
];
  return posts;
  */
});


