<template>
  <div id="app">
    <top-nav/>
    <router></router>
  </div>
</template>

<script>
import TopNav from './components/TopNav.vue'
import Router from "./components/Router.vue";

export default {
  name: "app",
  components: {
      Router,
      TopNav
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
