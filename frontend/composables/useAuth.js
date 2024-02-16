// import { jwtDecode } from 'jwt-decode'

export default () => {
  const useAuthToken = () => useState("accessToken");
  const useAuthUser = () => useState("userAuthenticated");
  const useAuthLoading = () => useState("authLoading", () => true);

  const setToken = (newToken) => {
    const authToken = useAuthToken();
    authToken.value = newToken;
  };

  const setIsAuthLoading = (value) => {
    const authLoading = useAuthLoading();
    authLoading.value = value;
  };

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
        // setToken(data.access_token)
        // setUser(data.user)
        console.log(fetchData);

        resolve(true);
      } catch (error) {
        reject(error);
      }
    });
  };

<<<<<<< HEAD
    const login = ({ email, password }) => {
        return new Promise(async (resolve, reject) => {
            try {
                const data = await $fetch('/api/auth/loginValidator', {
                    method: 'POST',
                    body: {
                        email,
                        password
                    }
                })

                // directing data to backend
                // const loginAccess = await $fetch('/api/auth/server', {
                //     method: 'POST',
                //     body: JSON.stringify({registration: null, loginAccess: data})
                // })

                // FROM BACKEND RESPONSE
                // logic based on response from backend
                // setToken(data.access_token)
                // setUser(data.user)
                console.log(data)
=======
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
          // logic based on response from backend token and user settled
          // console.log(useAuthToken())
          // setToken(fetchData.accessToken);
          // console.log(useAuthToken())
>>>>>>> origin/master

          // console.log(fetchData.token)
          console.log(fetchData);
          // setUser(data.user)
        }

<<<<<<< HEAD
        const jwt = 'refreshed'// jwtDecode(authToken.value)
=======
        resolve(true);
      } catch (error) {
        reject(error);
      }
    });
  };
>>>>>>> origin/master

  const refreshToken = () => {
    return new Promise(async (resolve, reject) => {
      try {
        const data = await $fetch("/api/auth/refresh");
        setToken(data.access_token);
        resolve(true);
      } catch (error) {
        reject(error);
      }
    });
  };

  const getUser = () => {
    return new Promise(async (resolve, reject) => {
      try {
        const data = await useFetchApi("/api/auth/user");
        setUser(data.user);
        resolve(true);
      } catch (error) {
        reject(error);
      }
    });
  };

  const refreshAccessToken = () => {
    const authToken = useAuthToken();

    if (!authToken) {
      return;
    }

    const jwt = "refreshed"; // jwtDecode(authToken.value)

    const newRefreshTime = Date.now(); //jwt.exp - 60000

    setTimeout(async () => {
      await refreshToken();
      refreshAccessToken();
    }, newRefreshTime);
  };

  const initAuth = () => {
    return new Promise(async (resolve, reject) => {
      setIsAuthLoading(true);
      try {
        await refreshToken();
        await getUser();

        refreshAccessToken();

        resolve(true);
      } catch (error) {
        reject(error);
      } finally {
        setIsAuthLoading(false);
      }
    });
  };
  return {
    login,
    register,
    useAuthUser,
    useAuthToken,
    initAuth,
    useAuthLoading,
  };
};
