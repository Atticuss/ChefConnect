<template>
  <div>
    <b-table striped hover :items="items"/>
  </div>
</template>

<script>
import IngredientsAPI from "@/services/Ingredients.js";
export default {
  name: "IngredientsTable",
  data() {
    return {
      items: []
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