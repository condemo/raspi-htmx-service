/** @type {import('tailwindcss').Config} */
import daisyui from 'daisyui'
module.exports = {
  content: ["services/web/public/views/**/*.{templ,js}"],
  theme: {
    extend: {},
  },
  plugins: [daisyui],
  daisyui: {
    themes: [
      "night",
      "cupcake",
      "halloween",
      "black",
      "sunset"
    ]
  }
}
