import { useGlobalAuthStore } from '../stores/useGobalStateAuthStore';

export default () => {
  const store = useGlobalAuthStore();

  const setUser = (newToken) => {
    store.login(newToken)

    const cookie = useCookie('token')
    cookie.value = newToken
  };

  const register = ({
    firstName,
    lastName,
    email,
    nickname,
    password,
    repeatPassword,
    aboutMe,
    avatarImg,
  }) => {
    console.log(avatarImg)
    return new Promise(async (resolve, reject) => {
      try {
        const fetchData = await $fetch("/api/auth/registerValidator", {
          method: "POST",
          body: {
            firstName,
            lastName,
            email,
            nickname,
            password,
            repeatPassword,
            aboutMe,
            avatarImg,
          },
        });
        if (fetchData.userSession) {
          setUser(fetchData.userSession)
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
        await $fetch("/api/auth/loginValidator", {
          method: "POST",
          body: {
            email,
            password,
          },
        }).then(response => {
          if (response) {
            setUser(response)
            resolve(response);
          }
        })

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
