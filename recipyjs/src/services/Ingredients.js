import axios from "axios";

export default {
  getIngredients() {
    return axios.get("/api/ingredients").then(response => {
      return response.data.ingredients;
    });
  },

  createIngredient(data) {
    return axios.post("/api/ingredients", data).then(response => {
      return response.data;
    });
  },

  getIngredientTypes() {
    return axios.get("/api/ingredients/types").then(response => {
      return response.data.ingredienttypes;
    });
  },

  createIngredientType(data) {
    return axios.post("/api/ingredients/types", data).then(response => {
      return response.data;
    });
  }
};
