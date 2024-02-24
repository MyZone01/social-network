export type user = {
    email: string,
    password: string,
    firstName: string,
    lastName: string,
    dateOfBirth: Date,
    avatarImage: string,
    nickname: string,
    aboutMe: string,
    isPublic: boolean,
}


export const editeUser = async(user: user, token: string) => {
    console.log(user, token);
    
    // return fetch('/api/user', {
    //     method: 'PUT',
    //     headers: {
    //         'Content-Type': 'application/json',
    //         Authorization: `Bearer ${token}`,
    //     },
    //     body: JSON.stringify(user),
    // });
}