<template>
  <Layout>
    <Header class="header-bg">
      <home-nav-bar></home-nav-bar>
    </Header>
    <Content class="content-bg">Content</Content>
    <Footer class="footer-bg">Footer</Footer>
  </Layout>
</template>

<script>
import HomeNavBar from '../utils/HomeNavBar'
export default {
  name: 'HomeController',
  components: {HomeNavBar}
}
</script>

<style scoped lang="scss">
  .header-bg {
    background: $light-primary-color;
  }
  .content-bg {
    padding-left: 64px;
    padding-right: 64px;
    min-height: 600px;
  }
  .footer-bg {
    background: $dark-primary-color;
    color: $title-color;
  }
</style>
