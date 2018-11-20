<script>

export default {
	props: {
		starting: {
			type: Boolean,
			default: false,
		},
	},

	data() {
		return { open: false }
	},

	created() {
		this.open = this.starting
	},

	methods: {
		toggle() {
			this.open = !this.open
		},
		turnOff() {
			this.open = false
		},
		turnOn() {
			this.open = true
		},
	},
	render(createElement) {
		return this.$scopedSlots.default({
			open: this.open,
			toggle: this.toggle,
			turnOn: this.turnOn,
			turnOff: this.turnOff,
		})
	},
}

</script>
