import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	safelist: [
        'text-resultError',
        'text-resultSuccess'
    ],
	theme: {
		extend: {
			colors: {
				background: '#262329',
				topbar: '#463F61',
				topbarHover: '#2c283d',
				item: '#2A243F',
				itemHover: `#473e69`,
				interactable: '#B87EFF',
				interactableHover: '#6e41a6',
				author: '#d3d3d3',
				description: '#cacaca',

				resultError: '#ff4646',
				resultSuccess: '#6dff5a'
			},
			fontFamily: {
				Itim: ['Itim', 'sans-serif']
			},
			screens: {
				'2cols' : '1600px'
			}
		}
	},

	plugins: [
	]
} satisfies Config;
