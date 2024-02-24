import { useGlobalAuthStore } from "../stores/useGobalStateAuthStore";
// import sendFile from "./uploadfile.js"
import FormData from "form-data";
import axios from "axios";
import { editeUser } from "./editeUser";
// import fs from 'fs';

export default () => {
  const store = useGlobalAuthStore();

  const setUser = (newToken) => {
    store.login(newToken, {});

    const cookie = useCookie("token");
    cookie.value = newToken;
  };

  const register = async ({ avatarImage, data }) => {
    return new Promise(async (resolve, reject) => {
      const body = new FormData();
      if (avatarImage) {
        body.append('file', avatarImage);
      }
      body.append('data', data);

      const response = await $fetch('/api/auth/register', {
        method: 'POST',
        body: body,
      });

      if (response.ok !== true) {
        reject(response);
      }
      if (response.ok === true && response.status !== 200) {
        // alert the user that the avatar does not upload correctly
        // 
        // 
        // and redirect to the index page
      }
      if (response.session && !store.isAuthenticated) {
        setUser(response.session);
        resolve(true);
      }
    });
  };

  const login = ({ email, password }) => {
    return new Promise(async (resolve, reject) => {
      try {
        const fetchData = await $fetch("/api/auth/login", {
          method: "POST",
          body: JSON.stringify({ data: { email, password } })
        });

        if (fetchData.userSession && !store.isAuthenticated) {
          setUser(fetchData.userSession);
          resolve(true);
        }
      } catch (error) {
        reject(error);
      }
    });
  };

  return {
    login,
    register,
  };
};
