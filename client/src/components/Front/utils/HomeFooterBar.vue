<template>
	<v-footer
			height="auto"
			color="primary lighten-1"
	>
		<v-layout
				justify-center
				row
				wrap
		>
			<v-flex
					primary
					lighten-2
					py-3
					text-xs-center
					white--text
					xs12
			>
				&copy;2019 — <strong>Insomnia</strong>
			</v-flex>
		</v-layout>
	</v-footer>
</template>

<script>
	import {VApp,
			VFooter,
			VLayout,
			VFlex
	} from 'vuetify/lib'
	export default {
		name: "HomeFooterBar",
		components: {
			VApp,
			VFooter,
			VLayout,
			VFlex
		},
	}
</script>

<style scoped>

</style>
