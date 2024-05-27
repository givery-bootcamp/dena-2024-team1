/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
    colors: {
      black: "#333333",
      white: "#FFFFFF",
      primary: "#007295",
      gray: {
        100: "#535353",
        200: "#818181"
      },
      border: "#CACACA"
    }
  },
  plugins: [],
}
