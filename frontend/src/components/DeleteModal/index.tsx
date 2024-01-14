import { Post } from "@/stores/posts/types";
import { useEffect } from "react";
import styles from "./DeleteModal.module.scss";
import { deletePostRequest } from "@/axios";
import { useRouter } from "next/router";
import { toast } from "react-hot-toast";

export default function index({
   status,
   post,
   setStatus,
   router,
}: {
   post: Post | undefined;
   status: boolean;
   setStatus: any;
   router?: any;
}) {
   const postDeleteHandle = async () => {
      if (post) {
         try {
            await deletePostRequest(post.id);
            setStatus(false);
            router.push("/");
         } catch (error) {
            console.log(error);
            toast.error("Post delete error");
            if (router) {
               router.push("/");
            }
         }
      }
   };

   return (
      <div
         style={
            status
               ? { pointerEvents: "auto", opacity: 1 }
               : { pointerEvents: "none", opacity: 0 }
         }
         className={styles.Modal}>
         <div>
            <h2>Are you sure to delete the post?</h2>
            <div>
               <button onClick={postDeleteHandle}>Delete</button>
               <button onClick={() => setStatus(false)}>Cencel</button>
            </div>
         </div>
      </div>
   );
}
