// const Dotenv = require('dotenv-webpack')
const resolve = require('path').resolve

// https://github.com/vuejs/vue-cli/issues/1134
// https://stackoverflow.com/questions/50898675/how-can-i-change-main-folders-in-vue

module.exports = {
	configureWebpack: {
		resolve: {
			// extensions: ['.gql'],
			alias: {
				'@': __dirname + '/client',
				'styles': __dirname + '/client/styles',
			},
		},

		entry: {
			app: './client/main.js'
		},

		// module: {
		// 	rules: [
		// 		{
		// 			test: /\.gql$/,
		// 			use: './client/queries/gql-loader.js',
		// 		},
		// 	]
		// },

		// plugins: [
		// 	new Dotenv({
		// 		path: process.env.NODE_ENV === 'production'
		// 			? resolve('.env.prod.sh')
		// 			: resolve('.env.dev.sh'),
		// 		safe: resolve('.env.template'),
		// 	})
		// ],
	},

	chainWebpack(config) {
		// Only convert .svg files that are imported by these files as Vue component
		const FILE_RE = /\.(vue|js|svg)$/

		// Use vue-cli's default rule for svg in non .vue .js .ts files
		config.module.rule('svg').issuer(file => !FILE_RE.test(file))

		// Use our loader to handle svg imported by other files
		config.module
			.rule('svg-component')
			.test(/\.svg$/)
			.issuer(file => FILE_RE.test(file))
			.use('vue')
			.loader('vue-loader')
			.end()
			.use('svg-to-vue-component')
			.loader('svg-to-vue-component/loader')
	}
}
