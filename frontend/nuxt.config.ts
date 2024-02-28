// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    head: {
      script: [
        { src: 'https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.js' },
        { src: 'https://cdnjs.cloudflare.com/ajax/libs/uikit/3.15.14/js/uikit.min.js' },
        { src: "https://cdn.jsdelivr.net/npm/select2@4.1.0-rc.0/dist/js/select2.min.js" },
        { type: "module", src: "https://unpkg.com/ionicons@5.5.2/dist/ionicons/ionicons.esm.js" },
      ],
      link: [
        { rel: 'stylesheet', href: 'https://cdn.jsdelivr.net/npm/select2@4.1.0-rc.0/dist/css/select2.min.css' },
      ]
      // style:[
      //   {href : '~/assets/css/style.css',}
      // ]
    }
  },
  ssr: true,
  devtools: { enabled: true },
  typescript: {
    includeWorkspace: true,
  },
  css: [
    '~/assets/css/style.css',
  ],
  modules: [
    '@nuxtjs/color-mode',
    '@unocss/nuxt',
    '@nuxt/image',
    '@nuxt/ui',
    '@vueuse/nuxt',
    '@nuxt/content',
    'nuxt-og-image',
    '@nuxtjs/tailwindcss',
    // '@nuxtjs/ionic',
    '@pinia/nuxt',
    '@pinia-plugin-persistedstate/nuxt',
  ],
  ui: {
    icons: ['heroicons', 'simple-icons']
  },
  nitro: {
    routeRules: {
      '/**': {
        headers: {
          'Cross-Origin-Embedder-Policy': 'require-corp',
          'Cross-Origin-Opener-Policy': 'same-origin',
        },
      },
    },
  },
  colorMode: {
    classSuffix: '',
  },
  // ionic: {
  //   integrations:
  //   {
  //     icons: true,
  //   },
  // },
  piniaPersistedstate: {
    cookieOptions: {
      sameSite: 'strict',
    },
    storage: 'sessionStorage'
  },
  runtimeConfig: {
    apiSecret: "thisisthesecretsauce",
  }
})
