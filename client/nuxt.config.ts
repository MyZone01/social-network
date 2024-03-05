import { createResolver } from '@nuxt/kit'
const { resolve } = createResolver(import.meta.url)

const ONE_DAY = 60 * 60 * 24 * 1000;
const ONE_WEEK = ONE_DAY * 7;

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  ssr: false,
  modules: [
    '@nuxt/ui',
    '@vueuse/nuxt',
    '@unocss/nuxt',
    '@nuxtjs/color-mode',
    '@pinia/nuxt',
    '@nuxt/image',
  ],
  colorMode: {
    classSuffix: '',
  },
  typescript: {
    includeWorkspace: true,
  },
  imports: {
    dirs: [
      resolve('./composables'), '~/composables',
      resolve('./stores'), '~/stores',
      resolve('~/types'), '~/types',
      resolve('~/api'), '~/api',
      resolve('~/utils'), '~/utils'
    ],
  },
  image: {
    domains: ['localhost:8081'],
  },
  pinia: {
    storesDirs: ['~/stores/**', '#/stores/**', '@/stores/**'],
  },
  runtimeConfig: {
    cookieName: process.env.COOKIE_NAME || "__social_cookie",
    cookieSecret: process.env.COOKIE_SECRET || "secret",
    cookieExpires: parseInt(process.env.COOKIE_REMEMBER_ME_EXPIRES || ONE_DAY.toString(), 10), // 1 day
    cookieRememberMeExpires: parseInt(process.env.COOKIE_REMEMBER_ME_EXPIRES || ONE_WEEK.toString(), 10), // 7 days
  },
})
