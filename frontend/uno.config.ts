import {
  defineConfig,
  presetIcons,
  transformerDirectives,
} from 'unocss'

export default defineConfig({
  shortcuts: {
    'border-base': 'border-gray-200 dark:border-gray-800',
    'bg-active': 'bg-gray:10',
    'bg-faded': 'bg-gray:5',
    'bg-base': 'bg-white p-2 dark:bg-gray-900',
    // 'text-base': 'text-black dark:text-white',
    'input-base': 'rounded-[50px] min-w-[300px] border-gray-200',
    'ui-base': 'bg-base text-base border-base',
  },
  theme: {
    extend: {
      spacing: {
        gap: '2.5rem',
        mdGap: '1.5rem',
        smGap: '.5rem',
        headerH: '75px',
        footerH: '2.5rem',
        filterH: '75px',
        logoGap: '7rem',
        mainCol: '3.3',
      },
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
      fontFamily: {
        body: ['Helvetica'],
        bodyAlt: ['OpenSans'],
        bodyHero: ['SuezOne'],
      },
      fontSize: {
        hero: '3rem',
        title: '2rem',
        subtitle: '1.75rem',
        medium: '1.25rem',
        normal: '.975rem',
        small: '.875rem',
        heroH: '60vh',
        logo: '4.5rem',
      },
      zIndex: {
        above: '10',
        below: '-1',
      },
    },
  },
  presets: [
    presetIcons(),
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