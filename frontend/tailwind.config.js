/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        orange: {
          DEFAULT: '#FF7F00', // Cam chủ đạo
          dark: '#E66A00', // Cam đậm
          light: '#FFD4A3', // Cam nhạt
        },
        navy: {
          DEFAULT: '#003366', // Xanh than chủ đạo
          dark: '#00224D', // Xanh than đậm
          light: '#4C6FA7', // Xanh than nhạt
        },
        white: '#FFFFFF', // Trắng
        gray: {
          light: '#F8F9FA', // Nền sáng
          dark: '#333333', // Chữ đậm
        }
      }
    },
  },
  plugins: [],
}

