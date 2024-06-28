/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      height: {
        "screen-minus-h20": "calc(100vh - 80px)",
      },
    },
    colors: {
      black: "#333333",
      white: "#FFFFFF",
      primary: "#007295",
      gray: {
        100: "#535353",
        200: "#818181",
      },
      border: "#CACACA",
      alert: "#DD5757",
    },
    zIndex: {
      modal: 10,
    },
  },
  plugins: [],
};
