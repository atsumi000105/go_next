export interface AuthTypes {
   user: User;
}

export type User = {
   status: boolean | null;
   id: number;
   name: string;
   email: string;
   token: string;
   created_at: string;
};
