export default defineAppConfig({
  API_BASE_URL: process.env.NUXT_PUBLIC_API_BASE || 'http://127.0.0.1:8081/',
  ui: {
    strategy: 'override',
    icon: 'mdi',
      primary: 'green',

  },
  })