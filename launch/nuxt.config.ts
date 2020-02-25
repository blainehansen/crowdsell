import pkg from '../package.json'

import { Configuration as NuxtConfiguration } from '@nuxt/types'
// import path from 'path'

// const purgecss = require('@fullhuman/postcss-purgecss')({

// 	// Specify the paths to all of the template files in your project
// 	content: [
// 		'./launch/**/*.vue',
// 		'./launch/**/*.js',
// 		'./launch/**/*.sss',
// 		'./.nuxt/**/*.vue',
// 		'./.nuxt/**/*.html',
// 		'./.nuxt/**/*.js',
// 	],

// 	// Include any special characters you're using in this regular expression
// 	defaultExtractor: (content: any) => content.match(/[A-Za-z0-9-_]+/g) || []
// })



const plugins = [
	// // require('precss'),
	// require('postcss-advanced-variables')({
	// 	disable: ['@import'],
	// }),
	// require('postcss-nested')({ bubble: ['screen'] }),

	// require('postcss-strip-inline-comments'),

	// require('postcss-easy-import')({
	// 	extensions: ['.css', '.sss']
	// }),
	// require('postcss-define-property')({
	// 	syntax: {
	// 		atrule: true,
	// 		parameter: '$',
	// 		variable: '$',
	// 		property: '+',
	// 		separator: '',
	// 	}
	// }),
	// require('tailwindcss')('./launch/tailwind.config.js'),
	// require('postcss-color-function'),
	// require('autoprefixer'),

	// purgecss,

	// // ...process.env.NODE_ENV === 'production'
	// // 	? [purgecss]
	// // 	: [],
]

export default {
	buildModules: ['@nuxt/typescript-build'],
	mode: 'spa',

	server: {
		port: 8080,
	},

	router: {
		linkActiveClass: 'active',
		linkExactActiveClass: 'active-exact',
	},

	head: {
		title: pkg.name,
		meta: [
			{ charset: 'utf-8' },
			{ name: 'viewport', content: 'width=device-width, initial-scale=1' },
			{ hid: 'description', name: 'description', content: pkg.description }
		],
		link: [
			{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
		]
	},

	// plugins: ['@/plugins/main.js'],
	// modules: ['vue-scrollto/nuxt'],

	// css: [
	// 	'@/assets/css/main.sss',
	// ],

	build: {
		// postcss: { plugins },

		extend(config) {
			// config.module = config.module || { rules: [] }

			// const vueStyle = { loader: 'vue-style-loader', options: { sourceMap: true } }
			// const css = {
			// 	loader: 'css-loader',
			// 	options: {
			// 		sourceMap: true,
			// 		importLoaders: 2,
			// 		onlyLocals: false
			// 	}
			// }
			// const postcss = {
			// 	loader: 'postcss-loader',
			// 	options: {
			// 		parser: 'sugarss',
			// 		sourceMap: true,
			// 		plugins,
			// 		order: 'presetEnvAndCssnanoLast'
			// 	}
			// }
			// const cssModule = {
			// 	...css,
			// 	options: {
			// 		...css.options,
			// 		localIdentName: '[local]_[hash:base64:5]',
			// 		modules: true,
			// 	},
			// }

			// config.module.rules.push({
			// 	test: /\.s(a)?ss$/,
			// 	oneOf: [
			// 		{ resourceQuery: /module/, use: [vueStyle, cssModule, postcss] },
			// 		{ use: [vueStyle, css, postcss] },
			// 	],
			// })

			config.resolve = config.resolve || {}
			config.resolve.extensions = config.resolve.extensions || []
			config.resolve.extensions.push('.ts')
		},

		// extractCSS: true,
	},

} as NuxtConfiguration

// TODO https://alligator.io/vuejs/vue-router-modify-head/
