// import { jwtDecode } from 'jwt-decode'

export default () => {
    const useAuthToken = () => useState('auth_token')
    const useAuthUser = () => useState('auth_user')
    const useAuthLoading = () => useState('auth_loading', () => true)

    const setToken = (newToken) => {
        const authToken = useAuthToken()
        authToken.value = newToken
    }

    const setIsAuthLoading = (value) => {
        const authLoading = useAuthLoading()
        authLoading.value = value
    }

    const setUser = (newUser) => {
        const authUser = useAuthUser()
        authUser.value = newUser
    }

    const register = ({ email, password }) => {
        return new Promise(async (resolve, reject) => {
            try {
                const data = await $fetch('/api/auth/login', {
                    method: 'POST',
                    body: {
                        email,
                        password
                    }
                })
                setToken(data.access_token)
                setUser(data.user)

                resolve(true)
            } catch (error) {
                reject(error)
            }
        })
    }

    const login = ({ email, password }) => {
        return new Promise(async (resolve, reject) => {
            try {
                const data = await $fetch('/api/auth/login', {
                    method: 'POST',
                    body: {
                        email,
                        password
                    }
                })
                setToken(data.access_token)
                setUser(data.user)

                resolve(true)
            } catch (error) {
                reject(error)
            }
        })
    }

    const refreshToken = () => {
        return new Promise(async (resolve, reject) => {
            try {
                const data = await $fetch('/api/auth/refresh')
                setToken(data.access_token)
                resolve(true)
            } catch (error) {
                reject(error)
            }
        })
    }

    const getUser = () => {
        return new Promise(async (resolve, reject) => {
            try {
                const data = await useFetchApi('/api/auth/user')
                setUser(data.user)
                resolve(true)
            } catch (error) {
                reject(error)
            }
        })
    }

    const refreshAccessToken = () => {
        const authToken = useAuthToken()

        if (!authToken) {
            return
        }

        const jwt = "refreshed"// jwtDecode(authToken.value)

        const newRefreshTime = Date.now() //jwt.exp - 60000

        setTimeout(async () => {
            await refreshToken()
            refreshAccessToken()
        }, newRefreshTime)
    }

    const initAuth = () => {
        return new Promise(async (resolve, reject) => {
            setIsAuthLoading(true)
            try {
                await refreshToken()
                await getUser()

                refreshAccessToken()

                resolve(true)
            } catch (error) {
                reject(error)
            } finally {
                setIsAuthLoading(false)
            }
        })
    }
    return {
        login,
        useAuthUser,
        useAuthToken,
        initAuth,
        useAuthLoading
    }
}