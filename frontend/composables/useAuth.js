export default () => {
  const useAuthToken = () => useState("accessToken");
  const useAuthUser = () => useState("userAuthenticated");
  const store = useGlobalAuthStore();
  const router = useRouter()
  // const useAuthLoading = () => useState("authLoading", () => true);

  const setToken = (newToken) => {
    const cookie = useCookie('token')
    const authToken = useAuthToken();
    store.login(newToken)
    cookie.value = newToken
    authToken.value = newToken;
  };

  // const setIsAuthLoading = (value) => {
  //   const authLoading = useAuthLoading();
  //   authLoading.value = value;
  // };

  const setUser = (newUser) => {
    const authUser = useAuthUser();
    authUser.value = newUser;
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
    return new Promise.all(async (resolve, reject) => {
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
        if (fetchData) {
          setUser(true)
        }
        setToken(fetchData)
        router.push('/')

        resolve(true);
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
        if (fetchData) {
          console.log(fetchData)
          setUser(true)
        }
        console.log(fetchData)
        setToken(fetchData)
        router.push('/')

        resolve(true);
      } catch (error) {
        reject(error);
      }
    });
  };

  // const refreshToken = () => {
  //   return new Promise(async (resolve, reject) => {
  //     try {
  //       const data = await $fetch("/api/auth/refresh");
  //       setToken(data.access_token);
  //       resolve(true);
  //     } catch (error) {
  //       reject(error);
  //     }
  //   });
  // };

  // const getUser = () => {
  //   return new Promise(async (resolve, reject) => {
  //     try {
  //       const data = await useFetchApi("/api/auth/user");
  //       setUser(data.user);
  //       resolve(true);
  //     } catch (error) {
  //       reject(error);
  //     }
  //   });
  // };

  // const refreshAccessToken = () => {
  //   const authToken = useAuthToken();

  //   if (!authToken) {
  //     return;
  //   }

  //   const jwt = "refreshed"; // jwtDecode(authToken.value)

  //   const newRefreshTime = Date.now(); //jwt.exp - 60000

  //   setTimeout(async () => {
  //     await refreshToken();
  //     refreshAccessToken();
  //   }, newRefreshTime);
  // };

  // const initAuth = () => {
  //   return new Promise(async (resolve, reject) => {
  //     setIsAuthLoading(true);
  //     try {
  //       await refreshToken();
  //       await getUser();

  //       refreshAccessToken();

  //       resolve(true);
  //     } catch (error) {
  //       reject(error);
  //     } finally {
  //       setIsAuthLoading(false);
  //     }
  //   });
  // };
  return {
    login,
    register,
    useAuthToken,
    useAuthUser,
  };
};
