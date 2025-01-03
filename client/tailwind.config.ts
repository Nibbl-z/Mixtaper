import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			colors: {
				background: '#262329',
				topbar: '#463F61',
				item: '#2A243F',
				interactable: '#B87EFF',
				author: '#d3d3d3',
				description: '#cacaca'
			}
		}
	},

	plugins: []
} satisfies Config;
