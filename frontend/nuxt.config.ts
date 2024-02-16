// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    head : {
      script :[
        { src: 'https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.js' },
        { type: "module", src: "https://unpkg.com/ionicons@5.5.2/dist/ionicons/ionicons.esm.js" },
        { src: "https://unpkg.com/ionicons@5.5.2/dist/ionicons/ionicons.js" },
        { src: 'https://cdnjs.cloudflare.com/ajax/libs/uikit/3.15.14/js/uikit.min.js' },
      ]
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
})
