import { Post } from "./types";
import { initialState } from "./initialState";
import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";

export const counterSlice = createSlice({
   name: "counter",
   initialState,
   reducers: {
      loadPosts: (state, action: PayloadAction<Post[]>) => {
         state.posts = action.payload;
      },
   },
});

export const { loadPosts } = counterSlice.actions;
export default counterSlice.reducer;
