import { useGlobalAuthStore } from "@/stores/useGlobalStateAuthStore";

export const loadUserInfos = async () => {
  const store = useGlobalAuthStore();

  console.log("FROM STORE by LoadUserInfos function", store.user);

  return store;
};

export const editeUser = async (user) => {
  const store = useGlobalAuthStore();
  console.log(user)

  const error = validateUserInfo(user);
  if (error) {
    return error;
  }

  return new Promise(async (resolve, reject) => {
    // const required = [user.email, ]
    const data = {
      email: user.email.trim(),
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

  const error = validateUpdatePassword(password.currentPassword, password.newPassword, password.repeatNewPassword)
  if (error) {
    return error
  }

  return new Promise(async (resolve, reject) => {
    const data = {
      email: store.user.email,
      password: password.currentPassword.trim(),
      newPassword: password.newPassword.trim(),
    };

    try {
      const response = await $fetch("/api/updatepassword", {
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

function validateUserInfo(userInfo) {

  // Validate nickname
  if (!userInfo.nickname || userInfo.nickname.trim() === "") {
    return "Nickname is required";
  }

  // Validate firstName
  if (!userInfo.firstName || userInfo.firstName.trim() === "") {
    return "First name is required";
  }

  // Validate lastName
  if (!userInfo.lastName || userInfo.lastName.trim() === "") {
    return "Last name is required";
  }

  // Validate email
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!userInfo.email || !emailRegex.test(userInfo.email)) {
    return "Invalid email address";
  }

  // Validate dateOfBirth
  // const currentDate = new Date();
  // const minDateOfBirth = new Date(currentDate.getFullYear() - 10, currentDate.getMonth(), currentDate.getDate());
  // const dateOfBirth = new Date(userInfo.dateOfBirth);
  // if (!userInfo.dateOfBirth || dateOfBirth > minDateOfBirth) {
  //     errors.dateOfBirth = 'You must be at least 10 years old';
  // }

  if (typeof userInfo.aboutMe != "string") {
    return "About You must be text";
  }

  // Validate state
  if (userInfo.isPublic !== 'public' && userInfo.isPublic !== 'private') {
    return "Profile status must be public or private";
  }

  return null;
}

function validateUpdatePassword(
  currentPassword,
  newPassword,
  repeatNewPassword
) {

  // Check if current password is provided
  if (!currentPassword.trim() || !newPassword.trim() || !repeatNewPassword.trim()) {
    return "All fields are required";
  }

  // Check if new password matches the repeat new password
  if (newPassword !== repeatNewPassword) {
    return "New password and repeat new password do not match";
  }

  // Check if new password is the same as the current password
  if (newPassword === currentPassword) {
    return "New password must be different from the current password";
  }

  // Password length should be at least 8 characters
  if (newPassword.length < 8) {
    return "New password must be at least 8 characters long";
  }

  // Password should contain at least one lowercase letter, one uppercase letter, one digit, and one special character
  // const regex =
  //   /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]+$/;
  // if (!regex.test(newPassword)) {
  //   return "New password must contain at least one lowercase letter, one uppercase letter, one digit, and one special character";
  // }
  // Password is valid
  return null;
}
