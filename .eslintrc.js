module.exports = {
	root: true,
	env: {
		node: true
	},
	extends: [
		'plugin:vue/essential',
		'eslint:recommended'
	],
	rules: {
		'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
		'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
		'no-inner-declarations': 'off',
		'no-unused-vars': ['error', { vars: 'all', args: 'after-used', argsIgnorePattern: '^_', caughtErrorsIgnorePattern: '^_', ignoreRestSiblings: false }],
	},
	overrides: [{
		files: ['./client/libs/*.js']
	}],
	parserOptions: {
		parser: 'babel-eslint'
	}
}

/* eslint-disable */
