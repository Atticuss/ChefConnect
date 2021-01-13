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

    <mdb-tbl v-else>
      <mdb-row>
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
          <mdb-row v-if="state.isAuthd">
            <mdb-col xl="3" md="6" class="mx-auto text-center">
              <mdb-btn color="info" @click.native="modal = true"
                >Add Recipe</mdb-btn
              >
            </mdb-col>
          </mdb-row>
        </mdb-col>
      </mdb-row>
    </mdb-tbl>

    <NewRecipe v-bind:modal.sync="modal"> </NewRecipe>
    <Auth> </Auth>
  </mdb-container>
</template>

<script>
import {
  mdbContainer,
  mdbTbl,
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
    mdbTbl,
    mdbRow,
    mdbCol,
    mdbBtn,
    mdbModal,
    mdbModalBody,
    Recipe,
    NewRecipe,
    Auth
  },
  props: {
    state: {
      type: Object
    }
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
  mounted() {
    this.$root.$on("new-recipe", data => {
      this.recipes.push(data);
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
