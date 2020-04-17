<template>
  <div>
    <b-table striped hover :items="items"/>
  </div>
</template>

<script>
import RecipesAPI from "@/services/Recipes.js";
export default {
  name: "RecipesAPI",
  data() {
    return {
      items: []
    };
  },
  created() {
    RecipesAPI.getRecipes().then(recipes => {
      this.items = recipes;
    });
  },
  mounted() {
    this.$root.$on("new-recipe", data => {
      this.items.push(data);
    });
  }
};
</script>