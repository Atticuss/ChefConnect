<template>
  <mdb-modal container fluid v-if="modal" @close="handleClose()">
    <mdb-modal-header>
      <mdb-modal-title tag="h4" class="w-100 text-center font-weight-bold"
        >Add new recipe</mdb-modal-title
      >
    </mdb-modal-header>
    <mdb-modal-body>
      <form class="mx-3 grey-text">
        <mdb-input
          name="name"
          label="Name"
          placeholder="Artisinally, hand-crafted PB&J"
          type="text"
          @input="handleInput($event, 'name')"
        />

        <mdb-input
          name="url"
          label="Url"
          placeholder="https://hipst.er/my-moms-pbj"
          type="text"
          @input="handleInput($event, 'url')"
        />

        <mdb-input
          name="prep-time"
          label="Prep Time"
          placeholder="3 hours"
          type="text"
          @input="handleInput($event, 'prep_time')"
        />

        <mdb-input
          name="cook-time"
          label="Cook Time"
          placeholder="30 minutes"
          type="text"
          @input="handleInput($event, 'cook_time')"
        />

        <mdb-input
          name="total-servings"
          label="Total Servings"
          placeholder="3"
          type="text"
          @input="handleInput($event, 'total_servings')"
        />

        <mdb-input
          name="directions"
          label="Directions"
          placeholder="Smear dat 'ish all over. Shove directly into facehole."
          type="textarea"
          @input="handleInput($event, 'directions')"
        />

        <p>
          Ingredients
        </p>

        <mdb-btn
          color="info"
          class="justify-content-left"
          size="sm"
          @click.native="numIngredients += 1"
        >
          Add
        </mdb-btn>

        <mdb-btn
          color="info"
          class="justify-content-left"
          size="sm"
          @click.native="addNewIngredient()"
        >
          New
        </mdb-btn>

        <mdb-tbl>
          <mdb-row
            v-for="idx in numIngredients"
            :index="idx"
            :key="idx"
            class="justify-content-start"
          >
            <mdb-col col="1" class="my-lg-4">
              <mdb-badge
                @click.native="onDelete"
                tag="a"
                color="danger-color"
                class="ml-2 float-center"
              >
                -
              </mdb-badge>
            </mdb-col>

            <mdb-col col="7">
              <mdb-select
                :filter="
                  (text, search) => {
                    return text.includes(search);
                  }
                "
                v-model="ingredients[idx - 1]"
                placeholder="Select an ingredient"
                label=""
                search
              >
              </mdb-select>
            </mdb-col>

            <mdb-col col="3">
              <mdb-input
                name="amount"
                label="Amount"
                placeholder="1 cup"
                type="text"
                @input="handleArrayInput($event, 'ingredients', 'amount', idx)"
              />
            </mdb-col>
          </mdb-row>
        </mdb-tbl>

        <mdb-tbl>
          <mdb-row class="justify-content-start">
            <mdb-col>
              <mdb-select
                :filter="
                  (text, search) => {
                    return text.includes(search);
                  }
                "
                v-model="tags"
                placeholder="Select tags"
                label=""
                search
                multiple
                selectAll
              />
            </mdb-col>

            <mdb-col>
              <mdb-select
                :filter="
                  (text, search) => {
                    return text.includes(search);
                  }
                "
                v-model="recipes"
                placeholder="Select related recipes"
                label=""
                search
                multiple
                selectAll
              />
            </mdb-col>
          </mdb-row>
        </mdb-tbl>
      </form>
    </mdb-modal-body>
    <mdb-modal-footer class="justify-content-center">
      <mdb-btn color="info" @click.native="saveRecipe">Add</mdb-btn>
    </mdb-modal-footer>
  </mdb-modal>
</template>

<script>
import {
  mdbBtn,
  mdbTbl,
  mdbRow,
  mdbCol,
  mdbBadge,
  mdbModal,
  mdbModalHeader,
  mdbModalTitle,
  mdbModalBody,
  mdbModalFooter,
  mdbInput,
  mdbSelect
} from "mdbvue";

import IngredientAPI from "@/services/Ingredients.js";
import TagAPI from "@/services/Tags.js";
import RecipeAPI from "@/services/Recipes.js";

export default {
  name: "NewRecipe",
  components: {
    mdbBtn,
    mdbTbl,
    mdbRow,
    mdbCol,
    mdbBadge,
    mdbModal,
    mdbModalHeader,
    mdbModalTitle,
    mdbModalBody,
    mdbModalFooter,
    mdbInput,
    mdbSelect
  },
  props: {
    modal: {
      type: Boolean
    }
  },
  data() {
    return {
      newValues: [],
      render: false,
      ingredients: [],
      numIngredients: 1,
      tags: [],
      recipes: []
    };
  },
  created() {
    IngredientAPI.getIngredients().then(data => {
      this.ingredients[0] = data.ingredients.map(ing => {
        return { text: ing.name, value: ing.uid };
      });
    });

    TagAPI.getTags().then(data => {
      this.tags = data.tags.map(tag => {
        return { text: tag.name, value: tag.uid };
      });
    });

    RecipeAPI.getRecipes(data => {
      this.recipes = data.recipes.map(recipe => {
        return { text: recipe.name, value: recipe.uid };
      });
    });
  },
  methods: {
    addIngredientRow() {
      this.ingredients[this.numIngredients] = JSON.parse(
        JSON.stringify(this.ingredients[0])
      );
      this.numIngredients += 1;
    },
    addTagRow() {
      this.tags[this.numTags] = JSON.parse(JSON.stringify(this.tags[0]));
      this.numTags += 1;
    },
    addRelatedRecipeRow() {
      this.recipes[this.numRecipes] = JSON.parse(
        JSON.stringify(this.recipes[0])
      );
      this.numRecipes += 1;
    },
    handleDelete(eventIndex) {
      this.recipes.splice(eventIndex, 1);
    },
    handleInput(val, type) {
      console.log(`handleInput(${val}, ${type})`);
      this.newValues[type] = val;
    },
    handleArrayInput(val, type, prop, idx) {
      console.log(`handleArrayInput(${val}, ${type}, ${prop}, ${idx})`);
      console.log(this.ingredients[0]);
      this.selectedIngredients[idx] = {
        selected: true,
        text: val,
        value: val
      };
      console.log(this.selectedIngredients[idx]);

      if (!Object.prototype.hasOwnProperty.call(this.newValues, type)) {
        this.newValues[type] = [];
      }

      if (typeof this.newValues[type][idx] === "undefined") {
        this.newValues[type][idx] = {};
      }

      //this.newValues[type][idx] = { uid: val };
      this.newValues[type][idx][prop] = val;
    },
    saveRecipe() {
      //IngredientAPI.createIngredient(this.newValues).then(function() {
      //  this.$emit("update:modal", false);
      //});
      console.log(this.newValues);
    },
    handleClose() {
      this.$emit("update:modal", false);
    }
  }
};
</script>

<style></style>
