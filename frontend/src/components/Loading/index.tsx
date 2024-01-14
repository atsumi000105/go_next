import { useAppSelector } from "@/stores/hooks";
import styles from "./Loading.module.scss";

export default function Home() {
   const { user } = useAppSelector((state) => state.auth);

   return (
      <>
         {user.status === null && (
            <div className={styles.Main}>
               <div className={styles.Animation}></div>
            </div>
         )}
      </>
   );
}
