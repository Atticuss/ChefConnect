<template>
  <mdb-modal v-if="display" @close="handleClose()">
    <mdb-modal-header>
      <mdb-modal-title tag="h4" class="w-100 text-center font-weight-bold"
        >Login</mdb-modal-title
      >
    </mdb-modal-header>
    <mdb-modal-body>
      <form class="mx-3 grey-text" @keyup.enter="login()">
        <ul class="list-group text-center" v-if="error.length > 0">
          <li class="list-group-item">
            {{ error }}
          </li>
        </ul>
        <mdb-input
          name="username"
          label="Username"
          type="text"
          @input="handleInput($event, 'username')"
          @keyup.enter.native="login()"
        />
        <mdb-input
          name="password"
          label="Password"
          type="password"
          @input="handleInput($event, 'password')"
          @keyup.enter.native="login()"
        />
      </form>
    </mdb-modal-body>
    <mdb-modal-footer class="justify-content-center">
      <mdb-btn color="info" @click.native="login()">Login</mdb-btn>
    </mdb-modal-footer>
  </mdb-modal>
</template>

<script>
import {
  mdbBtn,
  mdbModal,
  mdbModalHeader,
  mdbModalTitle,
  mdbModalBody,
  mdbModalFooter,
  mdbInput
} from "mdbvue";
import AuthAPI from "@/services/Auth.js";
export default {
  name: "Auth",
  components: {
    mdbInput,
    mdbBtn,
    mdbModal,
    mdbModalHeader,
    mdbModalTitle,
    mdbModalBody,
    mdbModalFooter
  },
  data() {
    return {
      credentialData: {},
      error: "",
      display: false
    };
  },
  mounted() {
    this.$root.$on(
      "login-modal",
      function() {
        this.display = true;
      }.bind(this)
    );
  },
  methods: {
    handleInput(val, type) {
      this.credentialData[type] = val;
    },
    handleClose() {
      this.display = false;
    },
    login() {
      AuthAPI.authRequest(this.$root, this.credentialData)
        .then(
          function() {
            this.error = "";
            this.display = false;
          }.bind(this)
        )
        .catch(
          function() {
            this.error = "Authentication failed";
          }.bind(this)
        );
    }
  }
};
</script>

<style></style>
