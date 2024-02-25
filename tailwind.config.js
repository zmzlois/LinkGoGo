/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["internal/templates/**/*.templ"],
  darkMode: "class",
  theme: {
    extend: {},
  },
  plugins: [],
  corePlugins: {
    preflight: true,
  }
}

