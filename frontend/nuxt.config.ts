// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: true,
  devtools: { enabled: true },
  typescript: {
    includeWorkspace: true,
  },
  modules: [
    '@nuxtjs/color-mode',
    '@unocss/nuxt',
    '@nuxt/image',
    '@nuxt/ui',
    '@vueuse/nuxt',
    '@nuxt/content',
    'nuxt-og-image',
  ],
  ui: {
    icons: ['heroicons', 'simple-icons']
  },
  colorMode: {
    classSuffix: '',
  },
})
