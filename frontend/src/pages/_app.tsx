import { store } from "@/stores";
import "@/styles/scss/index.scss";
import type { AppProps } from "next/app";
import { Provider } from "react-redux";
import Main from "./Main";

export default function App({ Component, pageProps }: AppProps) {
   return (
      <Provider store={store}>
         <Main>
            <Component {...pageProps} />
         </Main>
      </Provider>
   );
}
