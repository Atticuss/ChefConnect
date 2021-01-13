<template>
  <mdb-navbar color="elegant" dark>
    <mdb-navbar-brand href="https://mdbootstrap.com/">
      Voracious
    </mdb-navbar-brand>

    <mdb-navbar-toggler>
      <mdb-navbar-nav>
        <mdb-nav-item href="/" active>Home</mdb-nav-item>
        <mdb-nav-item href="/recipes/">Recipes</mdb-nav-item>
        <mdb-nav-item href="/ingredients/">Ingredients</mdb-nav-item>
        <mdb-nav-item href="/tags/">Tags</mdb-nav-item>
        <mdb-nav-item href="/search/">Search</mdb-nav-item>
      </mdb-navbar-nav>

      <mdb-btn
        v-if="!this.state.isAuthd"
        color="elegant"
        @click.native="$root.$emit('login-modal')"
      >
        Login
      </mdb-btn>

      <div v-if="this.state.isAuthd" class="text-light">
        Hello, {{ this.state.name }}
      </div>

      <mdb-dropdown v-if="this.state.isAuthd">
        <mdb-dropdown-toggle class="text-light" slot="toggle">
          <mdb-icon fa icon="bars" className="ml-2" />
        </mdb-dropdown-toggle>
        <mdb-dropdown-menu class="dark">
          <mdb-dropdown-item>Profile</mdb-dropdown-item>
          <mdb-dropdown-item>Favorites</mdb-dropdown-item>
          <mdb-dropdown-item @click.native="handleLogout">
            Logout
          </mdb-dropdown-item>
        </mdb-dropdown-menu>
      </mdb-dropdown>
    </mdb-navbar-toggler>
  </mdb-navbar>
</template>

<script>
import {
  mdbNavbar,
  mdbNavbarBrand,
  mdbNavbarToggler,
  mdbNavbarNav,
  mdbNavItem,
  mdbDropdown,
  mdbDropdownMenu,
  mdbDropdownToggle,
  mdbDropdownItem,
  mdbBtn,
  mdbIcon
} from "mdbvue";

import axios from "axios";
import jwt_decode from "jwt-decode";
import AuthAPI from "@/services/Auth.js";

export default {
  name: "NavbarPage",
  components: {
    mdbNavbar,
    mdbNavbarBrand,
    mdbNavbarToggler,
    mdbNavbarNav,
    mdbNavItem,
    mdbDropdown,
    mdbDropdownMenu,
    mdbDropdownToggle,
    mdbDropdownItem,
    mdbBtn,
    mdbIcon
  },
  props: {
    state: {
      type: Object
    }
  },
  data() {
    return {};
  },
  mounted() {
    this.$root.$on("successful-auth", data => {
      var decoded = jwt_decode(data.authToken);

      this.state.setIsAuthd(true);
      this.state.setName(decoded.name);
    });

    var token = window.localStorage.getItem("jwt");
    if (token !== "undefined") {
      var decoded = jwt_decode(token);
      this.state.setIsAuthd(true);
      this.state.setName(decoded.name);
      axios.defaults.headers.common.Authorization = token;
    }
  },
  methods: {
    handleLogout() {
      AuthAPI.logout();
      this.state.setIsAuthd(false);
    }
  }
};
</script>

<style></style>
