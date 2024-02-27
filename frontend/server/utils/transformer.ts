interface UserInfos {
    id: String,
    email: String,
    password: String,
    firstName: String,
    lastName: String,
    dateOfBirth: String,
    avatarImage: String,
    nickname: String,
    aboutMe: String,
    isPublic: true,
    createdAt: String,
    updatedAt: String,
    deletedAt: { Time: '0001-01-01T00:00:00Z', Valid: false }
}

export const secure = (user: UserInfos) => {
    return {
        email: user.email,
        firstName: user.firstName, 
        lastName: user.lastName, 
        dateOfBirth: (user.dateOfBirth).split('T')[0], 
        avatarImage: user.avatarImage, 
        nickname: user.nickname, 
        aboutMe: user.aboutMe, 
        isPublic: user.isPublic, 
        createdAt: user.createdAt, 
        updatedAt: user.updatedAt, 
    }
}

export const encoder = (token: string) => {
    return btoa(token)
}

export const decoder = (token: string) => {
    return atob(token)
}