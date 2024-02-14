// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({

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
    '@nuxtjs/ionic',
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
  ionic: {
    integrations:
    {
      icons: true,
    },
  },
  generate: {
    routes: [
      '/message',
      '/feed',
      '/event',
      '/login',
      '/register'
    ],
  }
})
