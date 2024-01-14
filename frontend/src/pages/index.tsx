import Head from "next/head";
import Navbar from "@/components/Navbar/Navbar";
import styles from "./Home.module.scss";
import { useEffect, useState } from "react";
import { postsRequest } from "@/axios";
import { useAppDispatch, useAppSelector } from "@/stores/hooks";
import { loadPosts } from "@/stores/posts/postsSlice";
import { Post } from "@/stores/posts/types";
import Link from "next/link";
import { FaTrash } from "react-icons/fa";
import { RiPencilFill } from "react-icons/ri";
import DeleteModal from "@/components/DeleteModal";
import { useRouter } from "next/router";

export default function Home() {
   const router = useRouter();
   const dispatch = useAppDispatch();
   const { posts } = useAppSelector((state) => state.posts);
   const { user } = useAppSelector((state) => state.auth);
   const [post, setPost] = useState<Post>();
   const [modalStatus, setModalStatus] = useState<boolean>(false);

   useEffect(() => {
      if (!modalStatus) {
         (async () => {
            try {
               const { data } = await postsRequest();
               dispatch(loadPosts(data));
            } catch (error) {
               console.log(error);
            }
         })();
      }
   }, [modalStatus, dispatch]);

   return (
      <>
         <Head>
            <title>Home</title>
            <meta name='description' content='Generated by create next app' />
            <meta name='viewport' content='width=device-width, initial-scale=1' />
            <link rel='icon' href='/favicon.ico' />
         </Head>
         <Navbar />
         <div className={styles.Container}>
            {posts &&
               posts.map((post: Post, index: number) => (
                  <div key={index}>
                     <div>
                        <Link href={`/post/${post.id}`}>
                           <h2>{post.title}</h2>
                        </Link>
                     </div>
                     <span>{post.subtitle}</span>
                     <p>{post.description}</p>
                     {user.id === post.owner && (
                        <div className={styles.Buttons}>
                           <button
                              onClick={() => {
                                 setModalStatus(true);
                                 setPost(post);
                              }}>
                              <FaTrash />
                           </button>
                           <Link href={`/update/${post.id}`}>
                              <RiPencilFill />
                           </Link>
                        </div>
                     )}
                  </div>
               ))}
         </div>
         <DeleteModal status={modalStatus} setStatus={setModalStatus} post={post} />
      </>
   );
}