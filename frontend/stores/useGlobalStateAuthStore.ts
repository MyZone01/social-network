import { defineStore } from 'pinia'

// Define interface for state properties
interface AuthState {
  isAuthenticated: Boolean
  token: String
  user: Object
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
    async login(newToken: String, userInfos: Object) {
      // Perform authentication logic
      this.isAuthenticated = true
      this.token = newToken
      this.user = userInfos
      await navigateTo('/')
      return
    },
    async logout() {
      // Perform logout logic
      await $fetch('/api/auth/clear', {
        method: 'DELETE',
      })
      this.isAuthenticated = false
      this.token = ''
      this.user = {}
      
      await navigateTo('/auth')
      return
    },
    async update(newToken: string, userInfos: Object) {
      // Perform update user logic
      this.isAuthenticated = true
      this.token = newToken
      this.user = userInfos
      return
    },
  },
})