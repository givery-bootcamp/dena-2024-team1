import { createSlice } from "@reduxjs/toolkit";

import { User } from "~/generated";

export type SessionUserState = {
  sessionUser?: User;
};

export const initialState: SessionUserState = {};

export const sessionUserSlice = createSlice({
  name: "sessionUser",
  initialState,
  reducers: {
    setSessionUser: (state, action) => {
      state.sessionUser = action.payload;
    },
    clearSessionUser: (state) => {
      state.sessionUser = undefined;
    },
  },
});

export default sessionUserSlice.reducer;
