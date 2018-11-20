export function optionsValidator(options) {
	return options.every(
		option => option && typeof option.name === 'string' && (option.value === null || typeof option.value === 'string')
	)
}
