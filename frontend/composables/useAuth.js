import { useGlobalAuthStore } from "../stores/useGobalStateAuthStore";
import FormData from "form-data";

export default () => {
  const store = useGlobalAuthStore();

  const setUser = (access) => {
    store.login(access.session, access.user);

    const cookie = useCookie("token");
    cookie.value = newToken;
  };

  const register = async ({ avatarImage, data }) => {
    return new Promise(async (resolve, reject) => {
      const body = new FormData();
      if (avatarImage) {
        body.append("file", avatarImage);
      }
      body.append("data", data);

      const response = await $fetch("/api/auth/register", {
        method: "POST",
        body: body,
      });

      if (response.ok !== true) {
        reject(response.message);
      }
      if (response.ok === true && response.status !== 200) {
        // alert the user that the avatar does not upload correctly
        //
        //
        // and redirect to the index page
      }
      if (response.session && !store.isAuthenticated) {
        setUser(response);
        resolve(true);
      }
    });
  };

  const login = ({ email, password }) => {
    return new Promise(async (resolve, reject) => {
        const response = await $fetch("/api/auth/login", {
          method: "POST",
          body: JSON.stringify({ data: { email, password } }),
        });

        if (response.ok !== true) {
          reject(response.message);
        }
        if (response.session && !store.isAuthenticated) {
          setUser(response.session);
          resolve(true);
        }
    });
  };

  return {
    login,
    register,
  };
};
