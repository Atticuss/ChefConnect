import axios from "axios";

export default {
  getRecipes() {
    return axios
      .get(`${process.env.VUE_APP_ROOT_DOMAIN}/recipes/`)
      .catch(error => {
        return Promise.reject(error);
      })
      .then(response => {
        return Promise.resolve(response.data);
      });
  },

  createRecipe(data) {
    return axios
      .post(`${process.env.VUE_APP_ROOT_DOMAIN}/recipes/`, data)
      .catch(error => {
        return Promise.reject(error);
      })
      .then(response => {
        return Promise.resolve(response.data);
      });
  }
};
