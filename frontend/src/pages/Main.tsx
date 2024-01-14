import React, { useEffect } from "react";
import Loading from "@/components/Loading";
import { useAppDispatch, useAppSelector } from "@/stores/hooks";
import { useRouter } from "next/router";
import { login } from "@/stores/auth/authSlice";
import { Toaster } from "react-hot-toast";
import { controlRequest } from "@/axios";

export default function Main({ children }: { children: any }) {
   const { user } = useAppSelector((state) => state.auth);
   const dispatch = useAppDispatch();
   const router = useRouter();

   useEffect(() => {
      if (localStorage.getItem("token")) {
         (async () => {
            try {
               const { data } = await controlRequest();
               dispatch(login({ ...data, status: true }));
            } catch (error) {
               console.log(error);
               router.push("/login");
               dispatch(
                  login({
                     status: false,
                     id: 0,
                     name: "",
                     email: "",
                     token: "",
                     created_at: "",
                  })
               );
            }
         })();
      } else {
         router.push("/login");
         dispatch(
            login({
               status: false,
               id: 0,
               name: "",
               email: "",
               token: "",
               created_at: "",
            })
         );
      }
   }, []);

   return (
      <>
         {children}
         {user.status === null && <Loading />}
         <Toaster position='top-right' reverseOrder={false} />
      </>
   );
}
