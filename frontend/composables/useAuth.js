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
        if (fetchData.userSession && !store.isAuthenticated) {
          setUser(fetchData.userSession)
          resolve(true);
        }
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
        if (fetchData.userSession && !store.isAuthenticated) {
          setUser(fetchData.userSession)
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
