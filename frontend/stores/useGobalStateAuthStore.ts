// stores/useGobalStateAuthStore.ts
import { defineStore } from 'pinia'

// Define interface for state properties
interface AuthState {
  isAuthenticated: boolean
  token: String | null
}
// Define store
export const useGlobalAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    isAuthenticated: false,
    token: null,
  }),
  actions: {
    login(newToken: String) {
      // Perform authentication logic
      this.isAuthenticated = true
      this.token = newToken
    },
    logout() {
      // Perform logout logic
      this.isAuthenticated = false
      this.token = null
    },
  },
})