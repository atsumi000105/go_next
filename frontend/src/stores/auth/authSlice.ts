import { User } from "./types";
import { initialState } from "./initialState";
import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";

export const counterSlice = createSlice({
   name: "counter",
   initialState,
   reducers: {
      login: (state, action: PayloadAction<User>) => {
         state.user = action.payload;
         if (action.payload.token.length > 0) {
            localStorage.setItem("token", action.payload.token);
         }
      },
   },
});

export const { login } = counterSlice.actions;
export default counterSlice.reducer;
