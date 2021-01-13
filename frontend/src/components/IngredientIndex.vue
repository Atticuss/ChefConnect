<template>
  <mdb-container>
    <mdb-row>
      <mdb-col col="11">
        <Ingredient
          v-for="(ingredient, index) in ingredients"
          :index="index"
          :key="index"
          :uid="ingredient.uid"
          :name="ingredient.name"
          :tags="ingredient.tags"
          @delete="handleDelete"
        />
        <mdb-row v-if="state.isAuthd">
          <mdb-col xl="3" md="6" class="mx-auto text-center">
            <mdb-btn color="info" @click.native="modal = true">
              Add Ingredient
            </mdb-btn>
          </mdb-col>
        </mdb-row>
      </mdb-col>
    </mdb-row>

    <NewIngredientModal v-bind:modal.sync="modal" />
  </mdb-container>
</template>

<script>
import { mdbContainer, mdbRow, mdbCol, mdbBtn } from "mdbvue";
import Ingredient from "@/components/ingredients/Ingredient";
import NewIngredientModal from "@/components/ingredients/NewIngredientModal";
import IngredientAPI from "@/services/Ingredients.js";
export default {
  name: "IngredientIndex",
  components: {
    mdbContainer,
    mdbRow,
    mdbCol,
    mdbBtn,
    Ingredient,
    NewIngredientModal
  },
  props: {
    state: {
      type: Object
    }
  },
  data() {
    return {
      ingredients: [],
      modal: false
    };
  },
  created() {
    IngredientAPI.getIngredients().then(data => {
      this.ingredients = data.ingredients;
    });
  },
  mounted() {
    this.$root.$on("new-ingredient", data => {
      this.ingredients.push(data);
    });
  },
  methods: {
    handleDelete(eventIndex) {
      this.ingredients.splice(eventIndex, 1);
    }
  }
};
</script>

<style></style>
