<template>
  <div id="categories">
    <b-container fluid class="cat-container">
      <b-row>
        <b-col md="2" offset-md="1">
          <NewRecipe/>
        </b-col>
      </b-row>
      <b-row>
        <b-col md="10" offset-md="1">
          <RecipesTable/>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
import NewRecipe from "./NewRecipeModal.vue";
import RecipesTable from "./RecipesTable.vue";

export default {
  name: "recipes",
  components: {
    NewRecipe,
    RecipesTable
  }
};
</script>

<style>
#recipes {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
