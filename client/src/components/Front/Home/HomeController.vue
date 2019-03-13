<template>
	<v-app>
		<home-nav-bar></home-nav-bar>
		<v-content>
			<v-container>

			</v-container>
		</v-content>
		<home-footer-bar></home-footer-bar>
	</v-app>
</template>

<script>
	import HomeNavBar from '../utils/HomeNavBar'
	import HomeFooterBar from "../utils/HomeFooterBar";
	import {
		VApp,
		VFooter,
		VContent,
		VContainer
	} from 'vuetify/lib'

	export default {
  name: 'HomeController',
  components: {
	  HomeFooterBar,
	  HomeNavBar,
	  VApp,
	  VContent,
	  VContainer,
	  VFooter
  },

}
</script>

<style scoped lang="scss">
  .layout {
    .header-bg {
      @include header-bg;
    }
    .content-bg {
      @include content-bg;
    }
    .footer-bg {
      @include footer-bg;

      color: $title-color;
    }
  }
</style>
