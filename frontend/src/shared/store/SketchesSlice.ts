import { createSlice } from "@reduxjs/toolkit";

import { Sketch } from "~/generated";
import { APIService } from "~/shared/services";

export type SketchesState = {
  sketches?: Sketch[];
};

export const initialState: SketchesState = {};

export const sketchesSlice = createSlice({
  name: "sketches",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(APIService.getSketches.fulfilled, (state, action) => {
      state.sketches = action.payload;
    });
  },
});

export default sketchesSlice.reducer;
