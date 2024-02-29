import { useGlobalAuthStore } from "@/stores/useGlobalStateAuthStore";

// export type user = {
//     email: string,
//     password: string,
//     firstName: string,
//     lastName: string,
//     dateOfBirth: Date,
//     avatarImage: string,
//     nickname: string,
//     aboutMe: string,
//     isPublic: boolean,
//     createdAt: Date,
//     updatedAt: Date,
//     deletedAt: Date,
// }

// interface Password {
//     current: string,
//     newPassword: string,
//     repeatNewPassword: string,
// }

export const loadUserInfos = async () => {
  const store = useGlobalAuthStore();

  console.log("FROM STORE by LoadUserInfos function", store.user);

  return store;
};

export const editeUser = async (user) => {
  const store = useGlobalAuthStore();

  return new Promise(async (resolve, reject) => {
    const data = {
      email: user.email.trim(),
      password: user.password.trim(),
      firstName: user.firstName.trim(),
      lastName: user.lastName.trim(),
      dateOfBirth: new Date(user.dateOfBirth),
      avatarImage: store.user.avatarImage,
      nickname: user.nickname.trim(),
      aboutMe: user.aboutMe.trim(),
      isPublic: user.isPublic === "public" ? true : false,
    };

    try {
      const response = await $fetch("/api/user", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `${store.token}`,
        },
        body: data,
      });
      if (!response.ok || response.status != 200) {
        reject(response.statusText);  
      }
      console.log(response);
      resolve(true);
    } catch (error) {
      reject(error);
    }
  });
};

export const updatePassword = async (password) => {};
