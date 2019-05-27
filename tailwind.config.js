const baseFontSize = 20
function rem(px) {
	return `${px / baseFontSize}rem`
}

const spacings =
	[...Array(105 / 5).keys()].map(i => i * 5)
	.concat([...Array(50 / 2).keys()].map(i => i * 2))
	.concat([...Array(400 / 25).keys()].map(i => i * 25))

const rems = spacings.reduce((acc, cur) => { acc[`${cur}`] = rem(cur); return acc}, {})
const percentages = spacings.reduce((acc, cur) => { acc[`${cur}p`] = `${cur}%`; return acc }, {})

module.exports = {
	prefix: '',
	important: false,
	separator: '_',
	theme: {
		screens: {
			sm: '640px',
			md: '800px',
			lg: '1024px',
			xl: '1280px',
		},
		colors: {
			transparent: 'transparent',

			white: '#ffffff',
			black: '#000000',

			'greenish-teal': '#3ecf90',
			azure: '#0ab3e4',
			'blue-blue': '#224bc5',
			'bright-blue': '#0d6fff',
			'violet-blue': '#370ccc',
			'ice-blue': '#f5f9ff',
			'light-teal': '#7ccfdd',
			'navy-blue': '#001e4b',
			'dark-blue-grey': '#162944',
			'pale-grey': '#f3f2fb',
			'blue-grey': '#596a89',
			'pale-grey-two': '#f5f6f8',
			'bluey-grey': '#929db1',
			'white-two': '#fdfdfd',
			'blue-blue-two': '#214ac2',
			'sky-blue': '#6b9cff',
			cornflower: '#6a75ff',
			steel: '#848e9e',
			silver: '#d3d8df',

			'reddish-grey': '#e9e3ff',

			'pastel-red': '#e86060',
		},
		spacing: {
			...rems,
		},
		backgroundColor: theme => theme('colors'),
		backgroundPosition: {
			bottom: 'bottom',
			center: 'center',
			left: 'left',
			'left-bottom': 'left bottom',
			'left-top': 'left top',
			right: 'right',
			'right-bottom': 'right bottom',
			'right-top': 'right top',
			top: 'top',
		},
		backgroundSize: {
			auto: 'auto',
			cover: 'cover',
			contain: 'contain',
		},
		borderColor: theme => ({
			...theme('colors'),
			default: theme('colors.gray.300', 'currentColor'),
		}),
		borderRadius: {
			none: '0',
			// default: '0.25rem',
			normal: rem(6),
			more: rem(10),
			very: rem(25),
			full: '9999px',
		},
		borderWidth: {
			default: '2px',
			// '0': '0',
			// '2': '2px',
			// '4': '4px',
			// '8': '8px',
		},
		boxShadow: {
			default: '0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)',
			md: '0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)',
			lg: '0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)',
			xl: '0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04)',
			'2xl': '0 25px 50px -12px rgba(0, 0, 0, 0.25)',
			inner: 'inset 0 2px 4px 0 rgba(0, 0, 0, 0.06)',
			outline: '0 0 0 3px rgba(66, 153, 225, 0.5)',
			none: 'none',
		},
		container: {
      center: true,
    },
		cursor: {
			auto: 'auto',
			default: 'default',
			pointer: 'pointer',
			wait: 'wait',
			text: 'text',
			move: 'move',
			'not-allowed': 'not-allowed',
		},
		fill: {
			current: 'currentColor',
		},
		flex: {
			'1': '1 1 0%',
			auto: '1 1 auto',
			initial: '0 1 auto',
			none: 'none',
		},
		flexGrow: {
			'0': '0',
			default: '1',
		},
		flexShrink: {
			'0': '0',
			default: '1',
		},
		// fontFamily: {
		//   sans: [
		//     '-apple-system',
		//     'BlinkMacSystemFont',
		//     '"Segoe UI"',
		//     'Roboto',
		//     '"Helvetica Neue"',
		//     'Arial',
		//     '"Noto Sans"',
		//     'sans-serif',
		//     '"Apple Color Emoji"',
		//     '"Segoe UI Emoji"',
		//     '"Segoe UI Symbol"',
		//     '"Noto Color Emoji"',
		//   ],
		//   serif: [
		//     'Georgia',
		//     'Cambria',
		//     '"Times New Roman"',
		//     'Times',
		//     'serif',
		//   ],
		//   mono: [
		//     'Menlo',
		//     'Monaco',
		//     'Consolas',
		//     '"Liberation Mono"',
		//     '"Courier New"',
		//     'monospace',
		//   ],
		// },
		fontSize: {
			hero: rem(70),
			header: rem(40),
			'tricky-logo': rem(31),
			'tricky-big-logo': rem(49),
			heavy: rem(30),
			lead: rem(25),
			base: rem(20),
			small: rem(18),
			tiny: rem(16),
		},
		fontWeight: {
			light: '400',
			normal: '500',
			heavy: '700',
			bold: '900',
		},
		height: theme => ({
			auto: 'auto',
			...theme('spacing'),
			...percentages,
			full: '100%',
			'small-tile': '40vh',
			'essay-tile': '65vh',
			tile: '80vh',
			banner: '110vh',
			'banner-bottom': '10vh',
			screen: '100vh',
		}),
		inset: {
			'0': '0',
			...percentages,
			auto: 'auto',
		},
		letterSpacing: {
			tighter: '-0.05em',
			tight: '-0.025em',
			normal: '0',
			wide: '0.025em',
			wider: '0.05em',
			widest: '0.1em',
		},
		lineHeight: {
			none: '1',
			tight: '1.2',
			'tight-kinda': '1.25',
			normal: '1.4',
			relaxed: '1.43',
			loose: '1.5',
			airy: '1.88',
		},
		listStyleType: {
			none: 'none',
			disc: 'disc',
			decimal: 'decimal',
		},
		margin: (theme, { negative }) => ({
			auto: 'auto',
			...theme('spacing'),
			...percentages,
			...negative(theme('spacing')),
		}),
		maxHeight: {
			full: '100%',
			screen: '100vh',
		},
		maxWidth: {
			xs: '20rem',
			sm: '24rem',
			md: '28rem',
			lg: '32rem',
			xl: '36rem',
			'2xl': '42rem',
			'3xl': '48rem',
			'4xl': '56rem',
			'5xl': '64rem',
			'6xl': '72rem',
			full: '100%',
		},
		minHeight: {
			'0': '0',
			full: '100%',
			screen: '100vh',
		},
		minWidth: {
			'0': '0',
			full: '100%',
		},
		objectPosition: {
			bottom: 'bottom',
			center: 'center',
			left: 'left',
			'left-bottom': 'left bottom',
			'left-top': 'left top',
			right: 'right',
			'right-bottom': 'right bottom',
			'right-top': 'right top',
			top: 'top',
		},
		opacity: {
			'0': '0',
			'25': '0.25',
			'50': '0.5',
			'75': '0.75',
			'100': '1',
		},
		order: {
			first: '-9999',
			last: '9999',
			none: '0',
			'1': '1',
			'2': '2',
			'3': '3',
			'4': '4',
			'5': '5',
			'6': '6',
			'7': '7',
			'8': '8',
			'9': '9',
			'10': '10',
			'11': '11',
			'12': '12',
		},
		padding: theme => theme('spacing'),
		stroke: {
			current: 'currentColor',
		},
		textColor: theme => theme('colors'),
		width: theme => ({
			auto: 'auto',
			...theme('spacing'),
			...percentages,
			'tricky-big-logo': rem(53.1),
			full: '100%',
			screen: '100vw',
		}),
		zIndex: {
			auto: 'auto',
			'0': '0',
			'10': '10',
			'20': '20',
			'30': '30',
			'40': '40',
			'50': '50',
		},
	},
	variants: {
		alignContent: ['responsive'],
		alignItems: ['responsive'],
		alignSelf: ['responsive'],
		appearance: ['responsive'],
		backgroundAttachment: ['responsive'],
		backgroundColor: ['responsive', 'hover', 'focus'],
		backgroundPosition: ['responsive'],
		backgroundRepeat: ['responsive'],
		backgroundSize: ['responsive'],
		borderCollapse: ['responsive'],
		borderColor: ['responsive', 'hover', 'focus'],
		borderRadius: ['responsive'],
		borderStyle: ['responsive'],
		borderWidth: ['responsive'],
		boxShadow: ['responsive', 'hover', 'focus'],
		cursor: ['responsive'],
		display: ['responsive'],
		fill: ['responsive'],
		flex: ['responsive'],
		flexDirection: ['responsive'],
		flexGrow: ['responsive'],
		flexShrink: ['responsive'],
		flexWrap: ['responsive'],
		float: ['responsive'],
		// fontFamily: ['responsive'],
		fontSize: ['responsive'],
		fontSmoothing: ['responsive'],
		fontStyle: ['responsive'],
		fontWeight: ['responsive', 'hover', 'focus'],
		height: ['responsive'],
		inset: ['responsive'],
		justifyContent: ['responsive'],
		letterSpacing: ['responsive'],
		lineHeight: ['responsive'],
		listStylePosition: ['responsive'],
		listStyleType: ['responsive'],
		margin: ['responsive'],
		maxHeight: ['responsive'],
		maxWidth: ['responsive'],
		minHeight: ['responsive'],
		minWidth: ['responsive'],
		objectFit: ['responsive'],
		objectPosition: ['responsive'],
		opacity: ['responsive'],
		order: ['responsive'],
		outline: ['responsive', 'focus'],
		overflow: ['responsive'],
		padding: ['responsive'],
		pointerEvents: ['responsive'],
		position: ['responsive'],
		resize: ['responsive'],
		stroke: ['responsive'],
		tableLayout: ['responsive'],
		textAlign: ['responsive'],
		textColor: ['responsive', 'hover', 'focus'],
		textDecoration: ['responsive', 'hover', 'focus'],
		textTransform: ['responsive'],
		userSelect: ['responsive'],
		verticalAlign: ['responsive'],
		visibility: ['responsive'],
		whitespace: ['responsive'],
		width: ['responsive'],
		wordBreak: ['responsive'],
		zIndex: ['responsive'],
	},
	corePlugins: {
		fontFamily: false,
	},
	plugins: [],
}
