<!-- 
The official vue-router is for SPAs, relying on anchor tags instead of full paths. 
Fuck. That. And fuck anyone who uses that. SPAs are garbage and break standards
we've trained users. Using a my own router, as defined here:
https://vueschool.io/articles/vuejs-tutorials/creating-your-own-router/ 
-->
<template>
  <div>
    <component :is="routedComponent" :state="state" />
    <Auth />
  </div>
</template>

<script>
import RecipeIndex from "./RecipeIndex";
import IngredientIndex from "./IngredientIndex";
import Auth from "./Auth";

const routes = {
  "/": RecipeIndex,
  "/recipes/": RecipeIndex,
  "/ingredients/": IngredientIndex
};

export default {
  components: {
    Auth
  },
  props: {
    state: {
      type: Object
    }
  },
  data() {
    return {
      current: window.location.pathname
    };
  },
  computed: {
    routedComponent() {
      return routes[this.current];
    }
  }
};
</script>

<style></style>
