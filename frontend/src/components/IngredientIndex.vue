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
        <mdb-row>
          <mdb-col xl="3" md="6" class="mx-auto text-center">
            <mdb-btn color="info" @click.native="modal = true"
              >Add Ingredient</mdb-btn
            >
          </mdb-col>
        </mdb-row>
      </mdb-col>
    </mdb-row>

    <NewIngredient v-bind:modal.sync="modal"> </NewIngredient>
  </mdb-container>
</template>

<script>
import {
  mdbContainer,
  mdbRow,
  mdbCol,
  mdbIcon,
  mdbBtn,
  mdbModal,
  mdbModalHeader,
  mdbModalTitle,
  mdbModalBody,
  mdbModalFooter,
  mdbInput,
  mdbTextarea,
} from "mdbvue";
import Ingredient from "@/components/ingredients/Ingredient";
import NewIngredient from "@/components/ingredients/NewIngredient";
import IngredientAPI from "@/services/Ingredients.js";
export default {
  name: "IngredientIndex",
  components: {
    mdbContainer,
    mdbRow,
    mdbCol,
    mdbIcon,
    mdbBtn,
    mdbModal,
    mdbModalHeader,
    mdbModalTitle,
    mdbModalBody,
    mdbModalFooter,
    mdbInput,
    mdbTextarea,
    Ingredient,
    NewIngredient,
  },
  data() {
    return {
      ingredients: [],
      modal: false,
    };
  },
  created() {
    IngredientAPI.getIngredients().then((data) => {
      this.ingredients = data.ingredients;
    });
  },
  mounted() {
    this.$root.$on("new-ingredient", (data) => {
      this.ingredients.push(data);
    });
  },
  methods: {
    handleDelete(eventIndex) {
      this.ingredients.splice(eventIndex, 1);
    },
  },
};
</script>

<style></style>
