import Head from "next/head";
import styles from "./Register.module.scss";
import Link from "next/link";
import toast from "react-hot-toast";
import { registerRequest } from "@/axios";
import { useAppDispatch, useAppSelector } from "@/stores/hooks";
import { login } from "@/stores/auth/authSlice";
import { useRouter } from "next/router";
import { useEffect } from "react";

const Register = () => {
   const { user } = useAppSelector((state) => state.auth);
   const dispatch = useAppDispatch();
   const router = useRouter();

   useEffect(() => {
      if (user.status) {
         router.push("/404");
      }
   }, [user.status, router]);

   const submitHandle = async (e: any) => {
      e.preventDefault();

      const { username, email, password } = e.target;

      if (username.value.length < 4 || username.value.length > 18) {
         toast.error("Username must be between 6 and 36 values");
         return;
      }
      if (email.value.length < 8 || email.value.length > 120) {
         toast.error("Email ust be between 6 and 36 values");
         return;
      }
      if (password.value.length < 6 || password.value.length > 36) {
         toast.error("Password must be between 6 and 36 values");
         return;
      }

      try {
         const { data } = await registerRequest({
            name: username.value,
            email: email.value,
            password: password.value,
         });
         dispatch(login({ ...data, status: true }));
         setTimeout(() => {
            router.push("/");
         }, 0);
      } catch (error) {
         console.log(error);
      }
   };

   return (
      <>
         <Head>
            <title>Register</title>
            <meta name='description' content='Generated by create next app' />
            <meta name='viewport' content='width=device-width, initial-scale=1' />
            <link rel='icon' href='/favicon.ico' />
         </Head>
         <div className={styles.Container}>
            <div className={styles.RegisterInner}>
               <h1>Register</h1>
               <form onSubmit={submitHandle} className={styles.Form}>
                  <input name='username' placeholder='Username' type='text' />
                  <input name='email' placeholder='E-Mail' type='email' />
                  <input name='password' placeholder='Password' type='password' />
                  <button>Register</button>
               </form>
               <Link href='/login'>Login</Link>
            </div>
         </div>
      </>
   );
};

export default Register;
