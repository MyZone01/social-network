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
  // pseudo: string;
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

export interface GroupMessage {
  ID: string,
  GroupID: string,
  SenderID: string,
  Sender: UserWithoutPassword,
  Content: string,
  CreatedAt: string,
  UpdatedAt: string,
  DeletedAt: string
}

export interface Event {
  ID: string,
  GroupID: string,
  title: string,
  description: string,
  CreatorID: string
}

export type UserWithoutPassword = Omit<User, "password">;