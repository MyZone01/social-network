import { useGlobalAuthStore } from "../stores/useGobalStateAuthStore";
import FormData from "form-data";
import axios from "axios";
// import fs from 'fs';

export default () => {
  const store = useGlobalAuthStore();

  const setUser = (newToken) => {
    store.login(newToken);

    const cookie = useCookie("token");
    cookie.value = newToken;
  };

  const register = ({
    firstName,
    lastName,
    email,
    nickname,
    password,
    repeatPassword,
    dateOfBirth,
    aboutMe,
    avatarImg,
    avatarUrl,
  }) => {
    const form = new FormData();
    form.append("file", avatarImg);
    
    return new Promise(async (resolve, reject) => {
      try {
        const fetchData = await $fetch("/api/auth/register", {
          method: "POST",
          body: {
            firstName,
            lastName,
            email,
            nickname,
            password,
            repeatPassword,
            dateOfBirth,
            aboutMe,
            form,
            avatarUrl,
          },
        });

        // if (avatarUrl != "") {
        //   await axios.post("http://localhost:8081/avatarupload", form, {
        //       headers: {
        //         'Authorization': `Bearer ${fetchData.userSession}`,
        //         'Content-Type': 'multipart/form-data'
        //       },
        //     })
        //     .then((res) => {
        //       console.log(res)
        //     })
        //     .catch((result) => {
        //       console.log(result)
        //     })
        // }
        if (fetchData.userSession && !store.isAuthenticated) {
          setUser(fetchData.userSession);
          resolve(true);
        }
      } catch (error) {
        reject(error);
      }
    });
  };

  const login = ({ email, password }) => {
    return new Promise(async (resolve, reject) => {
      try {
        const fetchData = await $fetch("/api/auth/loginValidator", {
          method: "POST",
          body: {
            email,
            password,
          },
        });
        // console.log(fetchData)
        // console.log(store.user)
        if (fetchData.userSession && !store.isAuthenticated) {
          setUser(fetchData.userSession, fetchData);
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
