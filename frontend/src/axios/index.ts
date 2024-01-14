import axios from "axios";
const HTTP = axios.create({
   baseURL: process.env.API_URL,
});

export const controlRequest = async () =>
   await HTTP.get("/user/control", {
      headers: {
         Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
   });

export const registerRequest = async (user: {
   name: string;
   email: string;
   password: string;
}) => await HTTP.post("/user/register", user);

export const loginRequest = async (user: { email: string; password: string }) =>
   await HTTP.post("/user/login", user);

export const postsRequest = async () =>
   await HTTP.get("/post/all", {
      headers: {
         Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
   });

export const createPostRequest = async (user: {
   title: string;
   subtitle: string;
   description: string;
}) =>
   await HTTP({
      url: "/post/create",
      method: "POST",
      headers: {
         Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
      data: user,
   });

export const postRequest = async (id: number) =>
   await HTTP.get(`/post/one/${id}`, {
      headers: {
         Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
   });

export const updatePostRequest = async (
   user: {
      title: string;
      subtitle: string;
      description: string;
   },
   id: number
) =>
   await HTTP({
      url: `/post/update/${id}`,
      method: "PUT",
      headers: {
         Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
      data: user,
   });

export const deletePostRequest = async (id: number) =>
   await HTTP({
      url: `/post/delete/${id}`,
      method: "DELETE",
      headers: {
         Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
   });
