import axios from "axios";

export default {
  getRecipes() {
    return axios.get("/api/recipes").then(response => {
      return response.data.recipes;
    });
  },

  createRecipes(data) {
    return axios.post("/api/recipes", data).then(response => {
      return response.data;
    });
  }
};
