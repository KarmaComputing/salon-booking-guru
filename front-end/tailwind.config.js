module.exports = {
    purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
    darkMode: false, // or 'media' or 'class'
    theme: {
        extend: {
            'prime-textcolor': {
                DEFAULT: 'var(--text-color)',
                secondary: 'var(--text-color-secondary)',
            },
            surface: {
                b: 'var(--surface-b)',
                e: 'var(--surface-e)',
                0: 'var(--surface-0)',
                50: 'var(--surface-50)',
                100: 'var(--surface-100)',
                200: 'var(--surface-200)',
                300: 'var(--surface-300)',
                400: 'var(--surface-400)',
                500: 'var(--surface-500)',
                600: 'var(--surface-600)',
                700: 'var(--surface-700)',
                800: 'var(--surface-800)',
                900: 'var(--surface-900)',
            },
        },
    },
    variants: {
        extend: {},
    },
    plugins: [],
};
