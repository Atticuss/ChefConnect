<template>
  <div id="app">
    <TopNav />
    <router-view />
  </div>
</template>

<script>
import TopNav from "./components/TopNav.vue";

export default {
  name: "App",
  components: {
    TopNav
  }
};
</script>

<style lang="scss">
$image-path: "~@/../mdb/mdbvue/img";
@import "~@/../mdb/mdbvue/scss/mdb-pro.scss";

@import url("https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap");
</style>
