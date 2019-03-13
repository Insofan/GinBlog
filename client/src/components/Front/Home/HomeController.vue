<template>
	<v-app>
		<v-navigation-drawer
				v-model="drawer"
				fixed
				app
		>
		</v-navigation-drawer>
	</v-app>
  <!--<Layout>-->
    <!--<div class="layout">-->
      <!--<Header class="header-bg">-->
        <!--<home-nav-bar></home-nav-bar>-->
      <!--</Header>-->
      <!--<Content class="content-bg">Content</Content>-->
      <!--<Footer class="footer-bg">-->
        <!--<div>-->
          <!--Footer-->
        <!--</div>-->
      <!--</Footer>-->
    <!--</div>-->
  <!--</Layout>-->
</template>

<script>
import HomeNavBar from '../utils/HomeNavBar'
import {VApp,
	VIcon,
	VCard,
	VNavigationDrawer,
} from 'vuetify/lib'
export default {
  name: 'HomeController',
  components: {
	  HomeNavBar,
	  VApp,
	  VIcon,
	  VCard,
	  VNavigationDrawer
  }
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
