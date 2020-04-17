<template>
  <div id="newcategory">
    <b-form @submit="onSubmit" @reset="onReset" v-if="show" inline>
      <b-button type="submit" variant="primary">Create</b-button>
      <b-form-group id="newCatInputGroup" label-for="newCatInput">
        <b-form-input
          id="newCatInput"
          type="text"
          v-model="form.name"
          required
          placeholder="Enter new category"
        />
      </b-form-group>
    </b-form>
  </div>
</template>

<script>
import CategoriesAPI from "@/services/Categories.js";
export default {
  data() {
    return {
      form: {
        name: ""
      },
      show: true
    };
  },
  methods: {
    onSubmit(evt) {
      evt.preventDefault();
      //alert(JSON.stringify(this.form));
      CategoriesAPI.createCategory(this.form).then(resp => {
        this.$root.$emit("new-cat", resp);
        this.reset();
      });
    },
    reset() {
      this.form.name = "";
      /* Trick to reset/clear native browser form validation state */
      this.show = false;
      this.$nextTick(() => {
        this.show = true;
      });
    }
  }
};
</script>