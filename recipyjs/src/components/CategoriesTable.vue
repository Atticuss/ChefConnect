<template>
  <div>
    <b-table striped hover :items="items"/>
  </div>
</template>

<script>
import CategoriesApi from "@/services/Categories.js";
export default {
  name: "CategoriesTable",
  data() {
    return {
      items: []
    };
  },
  created() {
    CategoriesApi.getCategories().then(cats => {
      this.items = cats;
    });
  },
  mounted() {
    this.$root.$on("new-cat", data => {
      this.items.push(data);
    });
  }
};
</script>