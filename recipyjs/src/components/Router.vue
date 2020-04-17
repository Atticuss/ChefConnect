<!-- credit: https://vueschool.io/articles/vuejs-tutorials/creating-your-own-router/ -->
<template>
  <component :is="routedComponent"></component>
</template>

<script>
import Recipes from "./Recipes";
import Ingredients from "./Ingredients";
import Categories from "./Categories";

const routes = {
  "/": Recipes,
  "/recipes": Recipes,
  "/ingredients": Ingredients,
  "/categories": Categories
};

export default {
  data() {
    return { current: window.location.pathname };
  },
  computed: {
    routedComponent() {
      return routes[this.current.toLowerCase()];
    }
  }
};
</script>