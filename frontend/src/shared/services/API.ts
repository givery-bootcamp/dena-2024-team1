import axios from "axios";
import { mapKeys,camelCase,mapValues,isArray,isObject } from "lodash";
import { createAsyncThunk } from "@reduxjs/toolkit";

import { Hello, Post } from "~/shared/models";

/**
 * 1. とりあえず今の状態でコミットする
 * 2. myAxiosの定義と設定部分を別ファイルに切り出す
 * 3. anyは仕方ないので、2のファイル全体でeslintを無視する
 */

const API_ENDPOINT_PATH =
  import.meta.env.VITE_API_ENDPOINT_PATH ?? "";

const myAxios = axios.create({ baseURL: API_ENDPOINT_PATH });

// eslint-disable-next-line
const mapKeysDeep = (data: any, callback: any) => {
  if (isArray(data)) {
    return data.map(innerData => mapKeysDeep(innerData, callback));
  } else if (isObject(data)) {
    return mapValues(mapKeys(data, callback), val =>
      mapKeysDeep(val, callback),
    );
  } else {
    return data;
  }
};

const mapKeysCamelCase = data =>
  mapKeysDeep(data, (_, key: string) => camelCase(key));

/*
myAxios.interceptors.request.use(
  request => {
    return request;
    return mapKeys(request.data, (_: string, key: string) => snakeCase(key));
});
*/

myAxios.interceptors.response.use(
  response => {
    const { data } = response;
    const convertedData = mapKeysCamelCase(data);
    return { ...response, data: convertedData };
  },
  error => {
    console.log(error);
    return Promise.reject(error);
  },
);

export const getHello = createAsyncThunk<Hello>("getHello", async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/hello`);
  return await response.json();
});

export const getPosts = createAsyncThunk<Post[]>("getPosts", async () => {
  const response = await myAxios.get("/posts");
  //console.log(response.data);
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


