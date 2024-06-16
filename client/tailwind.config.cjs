/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{ts,tsx}", "./node_modules/solid-daisy/dist/**/*.{js,ts,jsx,tsx}"],
    plugins: [require("daisyui")],
};
