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

export interface Group {
  ID: string,
  Title: string,
  Description: string,
  BannerURL: string,
  CreatorID: string,
  CreatedAt: string,
  UpdatedAt: string,
  DeletedAt: string,
  GroupMembers: GroupMember[]
}

export interface GroupMember {
  ID: string,
  GroupID: string,
  MemberID: string,
  Status: string,
  Role: string,
  User: UserWithoutPassword,
  CreatedAt: string,
  UpdatedAt: string,
  DeletedAt: string
}

export type UserWithoutPassword = Omit<User, "password">;