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
  aboutMe: string;
}

export interface ServerResponse<T> {
  status: string;
  session?: string;
  message: string;
  data?: T;
}

export type UserWithoutPassword = Omit<User, "password">;