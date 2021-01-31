<template>
  <div id="app">
    <TopNav :state="state" />
    <Router :state="state" />
  </div>
</template>

<script>
import TopNav from "./components/TopNav.vue";
import Router from "./components/Router.vue";
import Api from "@/services/Api.js";

var state = {
  isAuthd: false,
  name: "",
  setIsAuthd(val) {
    this.isAuthd = val;
  },
  setName(val) {
    this.name = val;
  }
};

export default {
  name: "App",
  components: {
    TopNav,
    Router
  },
  // a stupid hack to get around es-lint-prettier enforcing
  // unused imports. but the only thing within Api.js are
  // interceptors for handling re-auth on token expiry.
  created() {
    Api.nop();
  },
  data() {
    return {
      state: state
    };
  }
};
</script>

<!-- 
<style lang="scss">
$image-path: "~@/../mdb/mdbvue/img";
@import "~@/../mdb/mdbvue/scss/mdb-pro.scss";

@import url("https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap");
</style>
-->

<style></style>
