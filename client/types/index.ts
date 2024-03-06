export interface User {
  id: string;
  isLoggedIn: boolean;
  avatarImage?: string;
  isPublic: boolean;
  firstName: string;
  lastName: string;
  nickname: string;
  email: string;
  password: string;
  dateOfBirth: string;
}

export interface ServerResponse<T> {
  status: string;
  session: string;
  message: string;
  user: T;
}


export type UserWithoutPassword = Omit<User, "password">;