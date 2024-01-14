import { AuthTypes } from "./types";

export const initialState: AuthTypes = {
   user: {
      status: null,
      id: 0,
      name: "",
      email: "",
      token: "",
      created_at: "",
   },
};
