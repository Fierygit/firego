module.exports = {
  purge: [
    './src/**/*.html',
    './src/**/*.vue',
    './src/**/*.jsx',
  ],
  darkMode: 'class',
  theme: {
    extend: {},
  },
  variants: {
    extend: {
      backgroundColor: ['checked', 'active'],
      borderColor: ['checked'],
      display: ['group-hover'],
      opacity: ['disabled'],
      animation: ['hover'],
    },
  },
  plugins: [],
}
