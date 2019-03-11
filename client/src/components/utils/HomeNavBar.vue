<template>
  <div class="layout">
  <Layout>
    <Row class="nav-bg">
      <Col span="12">
        <div class="nav-title-label">Insomnia's World</div>
      </Col>
      <Col span="12">
        <Menu mode="horizontal"  active-name="1" style="background: #2d8cf0" class="menu">
          <MenuItem name="1" class="menu-item">
            <Icon type="ios-paper" />
            内容管理
          </MenuItem>
          <MenuItem name="2" class="menu-item">
            <Icon type="ios-people" />
            用户管理
          </MenuItem>
          <Submenu name="3"  class="menu-item">
            <template slot="title">
              <Icon type="ios-stats" />
              统计分析
            </template>
            <MenuGroup title="使用">
              <MenuItem name="3-1">新增和启动</MenuItem>
              <MenuItem name="3-2">活跃分析</MenuItem>
              <MenuItem name="3-3">时段分析</MenuItem>
            </MenuGroup>
            <MenuGroup title="留存">
              <MenuItem name="3-4">用户留存</MenuItem>
              <MenuItem name="3-5">流失用户</MenuItem>
            </MenuGroup>
          </Submenu>
          <MenuItem name="4" class="menu-item">
            <Icon type="ios-construct" />
            综合设置
          </MenuItem>
        </Menu>
      </Col>
    </Row>
  </Layout>
  </div>
</template>

<script>
export default {
  name: 'HomeNavBar'
}
</script>

<style scoped lang="scss">
  .nav-bg {
    background: $primary-color;
    min-width:1200px;
    .nav-title-label {
      @include flex-row-center;
      @include title-font;
      color: white;
      font-size: 52px;
      font-weight: bold;
      overflow: hidden;
    }
    .menu {
      background: $primary-color;
      .menu-item {
        color: white;
      }
    }
  }
</style>
