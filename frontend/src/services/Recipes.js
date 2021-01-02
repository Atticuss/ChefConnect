import axios from "axios";

export default {
  getRecipes() {
    return axios
      .get(`${process.env.VUE_APP_ROOT_DOMAIN}/recipes/`)
      .then(response => {
        return response.data;
      });
  },

  createRecipes(data) {
    return axios
      .post(`${process.env.VUE_APP_ROOT_DOMAIN}/recipes/`, data)
      .then(response => {
        return response.data;
      });
  }
};
