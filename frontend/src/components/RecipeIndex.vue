<template>
  <mdb-container>
    <mdb-modal v-if="loading" size="sm" centered>
      <mdb-modal-body>
        <div class="text-center">
          <div
            class="spinner-border"
            style="width: 3rem; height: 3rem"
            role="status"
          >
            <span class="sr-only">Loading...</span>
          </div>
        </div>
      </mdb-modal-body>
    </mdb-modal>

    <mdb-row v-else>
      <mdb-col col="11">
        <Recipe
          v-for="(recipe, index) in recipes"
          :index="index"
          :key="index"
          :uid="recipe.uid"
          :name="recipe.name"
          :tags="recipe.tags"
          @delete="handleDelete"
        />
        <mdb-row>
          <mdb-col xl="3" md="6" class="mx-auto text-center">
            <mdb-btn color="info" @click.native="modal = true"
              >Add Recipe</mdb-btn
            >
          </mdb-col>
        </mdb-row>
      </mdb-col>
    </mdb-row>

    <NewRecipe v-bind:modal.sync="modal"> </NewRecipe>
    <Auth> </Auth>
  </mdb-container>
</template>

<script>
import {
  mdbContainer,
  mdbRow,
  mdbCol,
  mdbBtn,
  mdbModal,
  mdbModalBody
} from "mdbvue";
import Recipe from "@/components/recipes/Recipe";
import NewRecipe from "@/components/recipes/NewRecipe";
import RecipeAPI from "@/services/Recipes.js";
import Auth from "@/components/Auth";

export default {
  name: "RecipeIndex",
  components: {
    mdbContainer,
    mdbRow,
    mdbCol,
    mdbBtn,
    mdbModal,
    mdbModalBody,
    Recipe,
    NewRecipe,
    Auth
  },
  data() {
    return {
      recipes: [],
      modal: false,
      loading: true
    };
  },
  created() {
    RecipeAPI.getRecipes().then(data => {
      this.recipes = data.recipes;
      this.loading = false;
    });
  },
  methods: {
    handleDelete() {
      console.log("deleted! but not really");
    }
  }
};
</script>

<style></style>
