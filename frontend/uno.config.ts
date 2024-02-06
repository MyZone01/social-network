import {
  defineConfig,
  presetAttributify,
  presetIcons,
  presetTypography,
  presetUno,
  transformerDirectives,
} from 'unocss'

export default defineConfig({
  shortcuts: { },
  theme: {
    extend: {
      colors: {
        primary: '#4767ff',
        secondary: '#d8ddfb',
        secondaryAlt: '#cfd5fc',
        secondaryAlt2: '#c9d1fa',
        white: '#fff',
        whiteTransparent: '#ffffffbf',
        whiteTransparent2: '#ffffff9e',
        black: '#000',
        blackTransparent: '#00000009',
        gray: '#e6e6ea',
        gray1: '#e6e6ea',
        gray2: '#c2c2d0',
        yellow: 'rgb(251, 255, 0)',
        cyan: '#00d6e7',
        alert: 'rgb(254, 34, 34)',
        warning: 'rgb(243, 203, 81)',
        success: '#00e788',
      },
    },
  },
  presets: [
    presetUno(),
    presetIcons(),
    presetAttributify(),
    presetTypography(),
  ],
  content: {
    filesystem: [
      './content/**/*.md',
    ],
  },
  transformers: [
    transformerDirectives(),
  ],
})