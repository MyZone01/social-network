import { useGlobalAuthStore } from "../stores/useGobalStateAuthStore";
// import sendFile from "./uploadfile.js"
import FormData from "form-data";
import axios from "axios";
import { editeUser } from "./editeUser";
// import fs from 'fs';

export default () => {
  const store = useGlobalAuthStore();

  const setUser = (newToken: string) => {
    store.login(newToken, {});

    const cookie = useCookie("token");
    cookie.value = newToken;
  };

  const register = async ({
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
  }: any) => {
    let avatarImage: File = avatarImg
    // const form = new FormData();
    // form.append("file", avatarImg);
    // console.log(avatarImg);


    //  console.log(reponse);
    return new Promise(async (resolve, reject) => {
      try {
        const body = new FormData();
        body.append('file', avatarImage, avatarImage.name);
        body.append('data', JSON.stringify({
          firstName,
          lastName,
          email,
          nickname,
          password,
          repeatPassword,
          dateOfBirth,
          aboutMe,
          avatarUrl,
          }));
        
        console.log(body);
        
        const fetchData: { userSession: string } = await $fetch('/api/auth/register', {
          method: 'POST',
          body: body,
        });
        // const fetchData: { userSession: string } = await $fetch("/api/auth/register", {
        //   method: "POST",
        //   body: {
        //     firstName,
        //     lastName,
        //     email,
        //     nickname,
        //     password,
        //     repeatPassword,
        //     dateOfBirth,
        //     aboutMe,
        //     // form,
        //     avatarUrl,
        //   },
        // });

        if (fetchData.userSession && !store.isAuthenticated) {
          console.log("test 1");
          const reponse = await sendFile(avatarImage, fetchData.userSession).then(async (res) => {
            return await res.json();
          });
          console.log("test 12");
          if (reponse.status === 200) {
            // u.avatarImg = reponse.data
            const u: user = {
              email,
              password: "",
              nickname,
              firstName,
              lastName,
              dateOfBirth,
              aboutMe,
              isPublic: true,
              avatarImage: reponse.data
            }
            console.log("test 123");

            const reponse2 = await editeUser(u, fetchData.userSession)
            console.log(reponse2);
          }
          setUser(fetchData.userSession);
          resolve(true);
        }
      } catch (error) {
        reject(error);
      }
    });
  };

  const login = ({ email, password }: any) => {
    return new Promise(async (resolve, reject) => {
      try {
        const fetchData: { userSession: string } = await $fetch("/api/auth/loginValidator", {
          method: "POST",
          body: {
            email,
            password,
          },
        });
        // console.log(fetchData)
        // console.log(store.user)
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
