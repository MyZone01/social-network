import { useGlobalAuthStore } from "../stores/useGlobalStateAuthStore";
import FormData from "form-data";

export default () => {
  const store = useGlobalAuthStore()

  const setUser = (access) => {
    store.login(access.session, access.user)

    const cookie = useCookie('token')
    // cookie.value = newToken;
  }

  const register = async ({ avatarImage, data }) => {
    return new Promise(async (resolve, reject) => {
      try {
        const body = new FormData()
        if (avatarImage) {
          body.append('file', avatarImage)
        }
        body.append('data', data)

        const response = await $fetch('/api/auth/register', {
          method: 'POST',
          body: body,
        })

        if (response.ok === false && response.status == 200) {
          // alert the user that the avatar does not upload correctly
          //
          //
          // and redirect to the index page
          // resolve(true);
        }
        response.user.dateOfBirth = dateFormatter(response.user.dateOfBirth)
        setUser(response);
        resolve(true);
      } catch (err) {
        reject(err)
      }
    })
  }

  const login = ({ email, password }) => {
    return new Promise(async (resolve, reject) => {
      try {
        const response = await $fetch('/api/auth/login', {
          method: 'POST',
          body: JSON.stringify({ data: { email, password } }),
        });
        response.user.dateOfBirth = dateFormatter(response.user.dateOfBirth)
        setUser(response);
        resolve(true);
      } catch (err) {
        reject(err);
      }
    })
  }

  return {
    login,
    register,
  }
}

const dateFormatter = (date) => {
  return date.split("T00:00:00Z")[0]
}