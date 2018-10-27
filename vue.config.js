const Dotenv = require('dotenv-webpack')
const resolve = require('path').resolve

// https://github.com/vuejs/vue-cli/issues/1134
// https://stackoverflow.com/questions/50898675/how-can-i-change-main-folders-in-vue

module.exports = {
	configureWebpack: {
		resolve: {
			extensions: ['.gql'],
			alias: {
				'@': __dirname + '/client'
			},
		},

		entry: {
			app: './client/main.js'
		},

		module: {
			rules: [
				{
					test: /\.gql$/,
					use: [
						'graphql-tag/loader',
					]
				},
			]
		},

		externals: {
			promisepay: 'promisepay',
		},

		plugins: [
			new Dotenv({
				path: process.env.NODE_ENV === 'production'
					? resolve('.env.prod.sh')
					: resolve('.env.dev.sh'),
				safe: resolve('.env.template'),
			})
		],
	},
}
