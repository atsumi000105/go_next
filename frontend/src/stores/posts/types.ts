export interface AuthTypes {
   posts: Post[];
}

export type Post = {
   id: number;
   title: string;
   subtitle: string;
   description: string;
   owner: number;
   user: {
      id: number;
      name: string;
      email: string;
      created_at: string;
   };
   created_at: string;
   updated_at: string;
   deleted_at: string | null;
};
