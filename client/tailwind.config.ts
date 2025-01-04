import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			colors: {
				background: '#262329',
				topbar: '#463F61',
				topbarHover: '#2c283d',
				item: '#2A243F',
				interactable: '#B87EFF',
				author: '#d3d3d3',
				description: '#cacaca'
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
