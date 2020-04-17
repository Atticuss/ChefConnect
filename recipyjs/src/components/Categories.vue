<template>
  <div id="categories">
    <b-container fluid class="cat-container">
      <b-row>
        <b-col md="4" offset-md="1">
          <NewCategory/>
        </b-col>
      </b-row>
      <b-row>
        <b-col md="10" offset-md="1">
          <CategoriesTable/>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
import NewCategory from "./NewCategory.vue";
import CategoriesTable from "./CategoriesTable.vue";

export default {
  name: "categories",
  components: {
    NewCategory,
    CategoriesTable
  }
};
</script>

<style>
#categories {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
