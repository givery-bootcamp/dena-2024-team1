import { createAsyncThunk } from '@reduxjs/toolkit';

import { Hello, Post } from '../models';

const API_ENDPOINT_PATH =
  import.meta.env.VITE_API_ENDPOINT_PATH ?? '';

export const getHello = createAsyncThunk<Hello>('getHello', async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/hello`);
  return await response.json();
});

export const getPosts = createAsyncThunk<Post[]>('getPosts', async () => {
  //const response = await fetch(`${API_ENDPOINT_PATH}/hello`);
  //return await response.json();
  const posts: Post[] = [{
    title: 'バナナはおやつに含まれますか？',
    userName: 'Tadashi',
    updatedAt: '5/27 14:23'
  },
  {
    title: 'りんごはおやつに含まれますか？',
    userName: 'Tadashi',
    updatedAt: '5/27 14:23'
  },
];
  return posts;
});
