<template>
  <div class="mt-4">
    <b-table striped hover :items="items"/>
  </div>
</template>

<script>
import IngredientsAPI from "@/services/Ingredients.js";
export default {
  name: "RecipeIngredientsTable",
  data() {
    return {
      fields: ["Ingredient", "Amount"],
      items: [{ ingredient: "", amount: "" }]
    };
  },
  created() {
    IngredientsAPI.getIngredients().then(ingredients => {
      this.items = ingredients;
    });
  },
  mounted() {
    this.$root.$on("new-ingredient", data => {
      this.items.push(data);
    });
  }
};
</script>