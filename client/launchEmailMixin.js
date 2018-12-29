export default {
	data() {
		return {
			email: "",
			attemptedSubmit: false,
		}
	},

	computed: {
		emailInvalid() {
			return this.attemptedSubmit && !(/^.+@.+\..+$/.test(this.email))
		},
	},

	methods: {
		submitEmail() {
			this.attemptedSubmit = true

			if (this.emailInvalid) console.log('nope')
			else {
				this.attemptedSubmit = false
				console.log(this.email)
			}
		}
	},
}
