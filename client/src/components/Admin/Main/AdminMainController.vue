<template>
	<v-app>
		<admin-nav-bar app></admin-nav-bar>
		<v-content>
			<v-container fill-height>

			</v-container>
		</v-content>
	</v-app>
</template>

<script>
	import {
		VApp,
		VContent,
		VContainer
	} from 'vuetify/lib'
	import AdminNavBar from '../utils/AdminNavBar';
	export default {
		name: 'AdminMainController',
		components: {
			VApp,
			VContent,
			VContainer,
			AdminNavBar
		},
		data() {
			return {
				isDrawer: false
			}
		}
	}
</script>

<style scoped>

</style>
