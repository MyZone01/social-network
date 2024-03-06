import { defineStore } from 'pinia'
// import axios from 'axios'

// Define interface for state properties
interface AuthState {
  isAuthenticated: Boolean
  token: String
  user: any
}
// Define store
export const useGlobalAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    isAuthenticated: false,
    token: '',
    user: {}
  }),
  persist: true,
  actions: {
    async login(newToken: String, userInfos: any) {
      // Perform authentication logic
      this.isAuthenticated = true
      this.token = newToken
      this.user = userInfos
      await navigateTo('/')
      return
    },
    async logout() {
      // Perform logout logic
      await useFetch('/api/auth/clear', {
        method: 'DELETE',
      })
      this.isAuthenticated = false
      this.token = ''
      this.user = {}

      await navigateTo('/auth')
      return
    },
  },
})