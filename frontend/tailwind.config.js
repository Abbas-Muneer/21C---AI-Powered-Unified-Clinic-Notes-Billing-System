/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{ts,tsx}"],
  theme: {
    extend: {
      fontFamily: {
        display: ["'Outfit'", "sans-serif"],
        body: ["'Source Sans 3'", "sans-serif"]
      },
      colors: {
        ink: "#102033",
        mist: "#eff4f8",
        accent: "#0d9488",
        amber: "#f59e0b",
        coral: "#dc6b52"
      },
      boxShadow: {
        soft: "0 24px 64px rgba(16,32,51,0.12)"
      }
    }
  },
  plugins: []
};
