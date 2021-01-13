<template>
  <mdb-dropdown multiLevel btn-group dropright class="float-right">
    <mdb-dropdown-toggle slot="toggle" color="info">
      Create
    </mdb-dropdown-toggle>
    <mdb-dropdown-menu style="min-width: 400px">
      <form class="px-2 py-2">
        <ul class="list-group text-center" v-if="error.length > 0">
          <li class="list-group-item">
            {{ error }}
          </li>
        </ul>
        <div class="form-group">
          <label for="ingName">Name</label>
          <input
            type="text"
            class="form-control"
            id="ingName"
            placeholder="Avacado"
            @input="handleInput($event.target.value, 'name')"
            @click.stop
          />
        </div>
        <mdb-btn color="info" @click.native="saveIngredient">Add</mdb-btn>
      </form>
    </mdb-dropdown-menu>
  </mdb-dropdown>
</template>

<script>
import {
  mdbBtn,
  mdbDropdown,
  mdbDropdownToggle,
  mdbDropdownMenu
} from "mdbvue";
import IngredientAPI from "@/services/Ingredients.js";
export default {
  name: "NewIngredientDropRight",
  components: {
    mdbBtn,
    mdbDropdown,
    mdbDropdownToggle,
    mdbDropdownMenu
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
            this.error = "An error occurred -- are you logged in?";
          }.bind(this)
        );
    },
    handleClose() {
      this.$emit("update:modal", false);
    }
  }
};
</script>

<style scoped>
.dropdown-toggle:after {
  display: none;
}
</style>
