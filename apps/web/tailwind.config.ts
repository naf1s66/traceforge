import type { Config } from "tailwindcss";

export default {
  content: ["./app/**/*.{ts,tsx,js,jsx}"],
  theme: { extend: {} },
  plugins: [],
  darkMode: "class",
} satisfies Config;
