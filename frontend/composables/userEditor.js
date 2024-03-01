import { useGlobalAuthStore } from "@/stores/useGlobalStateAuthStore";

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
        body: JSON.stringify(data),
      });

      await store.update(response.session, response.user);
      resolve(response.message);
    } catch (error) {
      reject(error);
    }
  });
};

export const updatePassword = async (password) => {
  const store = useGlobalAuthStore();

  return new Promise(async (resolve, reject) => {
    const data = {
      password: password.currentPassword.trim(),
      newPassword: password.newPassword.trim(),
      repeatNewPassword: password.repeatNewPassword.trim(),
    };

    try {
      const response = await $fetch("/api/updatepassword", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `${store.token}badtoken`,
        },
        body: JSON.stringify(data),
      });
      console.log(response);
      await store.update(response.session, response.user);
      resolve(response.message);
    } catch (error) {
      reject(error);
    }
  });
};
