<template>
  <mdb-modal v-if="modal" @close="handleClose()">
    <mdb-modal-header>
      <mdb-modal-title tag="h4" class="w-100 text-center font-weight-bold"
        >Add New Ingredient</mdb-modal-title
      >
    </mdb-modal-header>
    <mdb-modal-body>
      <form class="mx-3 grey-text" @keyup.enter.native="saveIngredient">
        <mdb-input
          name="name"
          label="Name"
          icon="stream"
          placeholder="Avacado"
          type="text"
          @input="handleInput($event, 'name')"
          @keyup.enter.native="saveIngredient"
        />
      </form>
    </mdb-modal-body>
    <mdb-modal-footer class="justify-content-center">
      <mdb-btn color="info" @click.native="saveIngredient">Add</mdb-btn>
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
  name: "NewIngredient",
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
      data: {}
    };
  },
  methods: {
    handleInput(val, type) {
      this.data[type] = val;
    },
    saveIngredient() {
      IngredientAPI.createIngredient(this.data).then(data => {
        this.$emit("update:modal", false);
        this.$root.$emit("new-ingredient", data);
      });
    },
    handleClose() {
      this.$emit("update:modal", false);
    }
  }
};
</script>

<style></style>
