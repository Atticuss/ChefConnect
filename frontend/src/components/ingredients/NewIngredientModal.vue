<template>
  <mdb-modal v-if="modal" @close="handleClose()">
    <mdb-modal-header>
      <mdb-modal-title tag="h4" class="w-100 text-center font-weight-bold"
        >Add New Ingredient</mdb-modal-title
      >
    </mdb-modal-header>
    <mdb-modal-body>
      <ul class="list-group text-center" v-if="error.length > 0">
        <li class="list-group-item">
          {{ error }}
        </li>
      </ul>
      <form class="mx-3 grey-text" @keyup.enter="saveIngredient">
        <mdb-input
          name="name"
          label="Name"
          icon="stream"
          placeholder="Avacado"
          type="text"
          @input="handleInput($event, 'name')"
          @keyup.enter="saveIngredient"
        />
      </form>
    </mdb-modal-body>
    <mdb-modal-footer class="justify-content-center">
      <mdb-btn color="info" @click="saveIngredient">
        Add
      </mdb-btn>
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
import IngredientAPI from "@/services/Ingredients.js";
export default {
  name: "NewIngredientModal",
  components: {
    mdbBtn,
    mdbModal,
    mdbModalHeader,
    mdbModalTitle,
    mdbModalBody,
    mdbModalFooter,
    mdbInput
  },
  props: {
    modal: {
      type: Boolean
    }
  },
  data() {
    return {
      data: {},
      error: ""
    };
  },
  methods: {
    handleInput(val, type) {
      this.data[type] = val;
    },
    saveIngredient() {
      IngredientAPI.createIngredient(this.data)
        .then(data => {
          this.$emit("update:modal", false);
          this.$root.$emit("new-ingredient", data);
        })
        .catch(
          function() {
            this.error = "An error ocurred";
          }.bind(this)
        );
    },
    handleClose() {
      this.$emit("update:modal", false);
    }
  }
};
</script>

<style></style>
