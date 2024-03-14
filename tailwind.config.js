/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/templates/*.html"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}

