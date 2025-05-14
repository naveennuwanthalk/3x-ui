<template>
  <a-config-provider :theme="{
    token: {
      colorPrimary: '#1890ff',
    },
  }">
    <div id="app">
      <a-layout class="layout">
        <a-layout-header>
          <div class="logo" />
          <a-menu
            theme="dark"
            mode="horizontal"
            :default-selected-keys="['1']"
            :style="{ lineHeight: '64px' }"
          >
            <a-menu-item key="1">Dashboard</a-menu-item>
            <a-menu-item key="2">Inbounds</a-menu-item>
            <a-menu-item key="3">Settings</a-menu-item>
          </a-menu>
        </a-layout-header>
        <a-layout-content style="padding: 0 50px">
          <div :style="{ background: '#fff', padding: '24px', minHeight: '280px' }">
            <!-- Router view will go here -->
          </div>
        </a-layout-content>
        <a-layout-footer style="text-align: center">
          X-UI Panel ©2023
        </a-layout-footer>
      </a-layout>
    </div>
  </a-config-provider>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial,
    'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol',
    'Noto Color Emoji';
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.layout {
  min-height: 100vh;
}

.logo {
  width: 120px;
  height: 31px;
  background: rgba(255, 255, 255, 0.2);
  margin: 16px 24px 16px 0;
  float: left;
}
</style>