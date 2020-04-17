<template>
  <div>
    <b-button v-b-modal.modal1>Add New Recipe</b-button>

    <!-- Modal Component -->
    <b-modal id="modal1" title="Add New Recipe" size="xl">
      <b-row>
        <b-col md="4">
          <NewIngredient/>
        </b-col>
        <b-col md="4">
          <NewCategory/>
        </b-col>
      </b-row>
      <b-row>
        <b-col md="10" md-offset="1">
          <IngredientsTable/>
        </b-col>
      </b-row>
    </b-modal>
  </div>
</template>

<script>
import NewCategory from "./NewCategory.vue";
import NewIngredient from "./NewIngredient.vue";
import IngredientsTable from "./NewRecipe/RecipeIngredientsTable.vue"

export default {
  name: "new-recipe",
  components: {
    NewCategory,
    NewIngredient,
    IngredientsTable
  }
};
</script>

<style>
#new-recipe {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>