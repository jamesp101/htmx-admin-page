

/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./templates/**/*.templ"],
    theme: {
        extend: {},
    },
    plugins: [
        require('@tailwindcss/typography'),
        require('@tailwindcss/forms'),
        require('@tailwindcss/aspect-ratio'),
        require('@tailwindcss/line-clamp'),
        require("daisyui")
    ],
    daisyui: {
        themes: [ "cupcake", "bumblebee", "forest"],
        darkTheme: false,
    },

}

