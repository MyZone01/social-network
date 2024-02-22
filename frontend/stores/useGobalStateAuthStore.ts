// stores/useGobalStateAuthStore.ts
import { defineStore } from 'pinia'
import axios from 'axios'

// Define interface for state properties
interface AuthState {
  isAuthenticated: boolean
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
      return
    },
    logout() {
      // Perform logout logic
      this.isAuthenticated = false
      this.token = ''
      this.user = {}
      return
    },
  },
})