<template>
  <Layout>
    <Header class="header">
      <home-nav-bar></home-nav-bar>
    </Header>
    <Content>Content</Content>
    <Footer>Footer</Footer>
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
  .header {
    background: $light-primary-color;
  }

</style>
