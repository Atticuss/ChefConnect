<template>
  <mdb-modal container fluid v-if="modal" @close="handleClose()" size="lg">
    <mdb-modal-header>
      <mdb-modal-title tag="h4" class="w-100 text-center font-weight-bold"
        >Add new recipe</mdb-modal-title
      >
    </mdb-modal-header>
    <mdb-modal-body>
      <form class="mx-3 grey-text">
        <ul class="list-group text-center" v-if="error.length > 0">
          <li class="list-group-item">
            {{ error }}
          </li>
        </ul>
        <mdb-input
          name="name"
          label="Name"
          placeholder="Artisinally hand-crafted PB&J"
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

        <mdb-row class="justify-content-between">
          <mdb-col col="8">
            <mdb-row>
              <mdb-col col="1">
                <mdb-badge
                  @click.native="numIngredients += 1"
                  tag="a"
                  color="info-color"
                  class="ml-2 mt-2 justify-content-left"
                >
                  +
                </mdb-badge>
              </mdb-col>

              <mdb-col>
                <p class="h5 mt-2">
                  Add Ingredients
                </p>
              </mdb-col>
            </mdb-row>
          </mdb-col>

          <mdb-col col="3">
            <NewIngredientDropRight />
          </mdb-col>
        </mdb-row>

        <mdb-tbl>
          <mdb-row
            v-for="(item, idx) in numIngredients"
            :index="idx"
            :key="idx"
            class="justify-content-center"
          >
            <mdb-col col="1" class="my-lg-4">
              <mdb-badge
                @click.native="onDelete"
                tag="a"
                color="danger-color"
                class="ml-2 mt-3"
              >
                -
              </mdb-badge>
            </mdb-col>

            <mdb-col col="7">
              <mdb-select
                v-model="ingredients[idx]"
                search
                @search="handleIngredientSearch($event, idx)"
                @change="handleArrayInput($event, 'ingredients', 'uid', idx)"
                disableFilter
                placeholder="Select an ingredient"
                label=""
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
                @change="handleMultiSelectInput($event, 'tags')"
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
                @change="handleMultiSelectInput($event, 'related_recipes')"
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

import NewIngredientDropRight from "@/components/ingredients/NewIngredientDropRight";

import IngredientAPI from "@/services/Ingredients.js";
import TagAPI from "@/services/Tags.js";
import RecipeAPI from "@/services/Recipes.js";

var intTypeList = ["prep_time", "cook_time", "total_servings"];

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
    mdbSelect,
    NewIngredientDropRight
  },
  props: {
    modal: {
      type: Boolean
    }
  },
  data() {
    return {
      error: "",
      newValues: {},
      render: false,
      ingredients: [],
      numIngredients: 1,
      tags: [],
      recipes: [],
      showNewIngredient: true
    };
  },
  created() {
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
      this.newValues[type] = val;

      if (intTypeList.includes(type)) {
        this.newValues[type] = parseInt(val);
      }
    },
    handleArrayInput(val, type, prop, idx) {
      if (val == null) {
        return;
      }

      if (!Object.prototype.hasOwnProperty.call(this.newValues, type)) {
        this.newValues[type] = [];
      }

      if (typeof this.newValues[type][idx] === "undefined") {
        this.newValues[type][idx] = {};
      }

      this.newValues[type][idx][prop] = val;
    },
    handleIngredientSearch(text, idx) {
      if (text.length >= 3) {
        IngredientAPI.searchIngredients(text).then(data => {
          let options = data.ingredients.map(ing => {
            return { text: ing.name, value: ing.uid, index: 0 };
          });

          // required for a state update to trigger. more details:
          // https://vuejs.org/v2/guide/reactivity.html#For-Arrays
          this.$set(this.ingredients, idx, options);
        });
      }
    },
    handleMultiSelectInput(val, type) {
      if (val.length == 0) {
        delete this.newValues[type];
      } else {
        this.newValues[type] = val;
      }
    },
    saveRecipe() {
      RecipeAPI.createRecipe(this.newValues)
        .then(data => {
          this.error = "";
          this.$emit("update:modal", false);
          this.$root.$emit("new-recipe", data);
        })
        .catch(
          function() {
            this.error = "An error ocurred";
          }.bind(this)
        );
    },
    handleClose() {
      this.$emit("update:modal", false);
    }
  }
};
</script>

<style></style>
