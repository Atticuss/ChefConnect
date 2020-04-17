<template>
  <div id="ingredients">
    <b-container fluid class="cat-container">
        <b-row>
            <b-col md="4" offset-md="1">
                <NewIngredient/>
            </b-col>
        </b-row>
        <b-row>
            <b-col md="10" offset-md="1">
                <IngredientsTable/>
            </b-col>
        </b-row>
    </b-container>
  </div>
</template>

<script>
import NewIngredient from './NewIngredient.vue'
import IngredientsTable from './IngredientsTable.vue'

export default {
  name: 'ingredients',
  components: {
    NewIngredient,
    IngredientsTable
  }
}
</script>

<style>
#ingredients {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
