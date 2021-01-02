import Vue from "vue";
import Router from "vue-router";
import RecipeIndex from "@/components/RecipeIndex";
import IngredientIndex from "@/components/IngredientIndex";

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/",
      name: "Root",
      component: RecipeIndex
    },
    {
      path: "/recipes",
      name: "RecipeIndex",
      component: RecipeIndex
    },
    {
      path: "/ingredients",
      name: "IngredientIndex",
      component: IngredientIndex
    }
  ]
});
