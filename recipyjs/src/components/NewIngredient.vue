<template>
  <div id="newingredient">
    <b-form @submit="onSubmit" @reset="onReset" v-if="show" inline>
      <b-button type="submit" variant="primary">Create</b-button>
      <b-form-group id="newIngredientInputGroup" label-for="newIngredientInput">
        <b-form-input
          id="newIngredientInput"
          type="text"
          v-model="form.name"
          required
          placeholder="Enter new ingredient"
        />
      </b-form-group>
    </b-form>
  </div>
</template>

<script>
import IngredientsAPI from "@/services/Ingredients.js";
export default {
  data() {
    return {
      form: {
        name: ""
      },
      show: true
    };
  },
  methods: {
    onSubmit(evt) {
      evt.preventDefault();
      IngredientsAPI.createIngredient(this.form).then(resp => {
        this.$root.$emit("new-ingredient", resp);
        this.reset();
      });
    },
    reset() {
      this.form.name = "";
      /* Trick to reset/clear native browser form validation state */
      this.show = false;
      this.$nextTick(() => {
        this.show = true;
      });
    }
  }
};
</script>