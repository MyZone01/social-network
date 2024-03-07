import type { User } from "~/types";

export const useAuth = () => {
  const authUser = useAuthUser();

  const setUser = (user: User | null) => {
    authUser.value = user;
  };

  const setCookie = (cookie: any) => {
    cookie.value = cookie;
  };

  const login = async (username: string, password: string, rememberMe: boolean) => {
    const response = await $fetch<{ user: User }>("/api/auth/login", {
      method: "POST",
      body: {
        username,
        password,
        rememberMe,
      },
    });

    const user = response.user;
    setUser({ ...user, isLoggedIn: true });

    return authUser;
  };

  const register = async (avatarImage: any, data: string) => {
    return new Promise(async (resolve, reject) => {
      try {
        const body = new FormData();
        if (avatarImage) {
          body.append("file", avatarImage);
        }
        body.append("data", data);

        const response = await $fetch<{ ok: boolean, status: number, user: User }>("/api/auth/register", {
          method: "POST",
          body: body,
        });

        if (response.ok === false && response.status == 200) {
        }
        const user = response.user;
        setUser({ ...user, isLoggedIn: true });
        resolve(true);
      } catch (err) {
        reject(err);
      }
    });
  };

  const logout = async () => {
    const data = await $fetch<{ user: User }>("/api/auth/logout", {
      method: "DELETE",
    });

    setUser(null);
  };

  const me = async () => {
    if (!authUser.value || !authUser.value?.firstName) {
      try {
        const response = await $fetch<{ user: User }>("/api/auth/me", {
          headers: useRequestHeaders(["cookie"]) as HeadersInit,
        });

        if (!response) {
          setUser(null);
          setCookie(null);
        } else {
          setUser({ ...authUser.value, ...response.user });
        }
      } catch (error) {
        setCookie(null);
      }
    }

    return authUser;
  };

  return {
    login,
    register,
    logout,
    me,
    setUser,
  };
};