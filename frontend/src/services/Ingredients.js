import axios from "axios";

export default {
  getIngredients() {
    return axios
      .get(`${process.env.VUE_APP_ROOT_DOMAIN}/ingredients/`)
      .catch(error => {
        return Promise.reject(error);
      })
      .then(response => {
        return Promise.resolve(response.data);
      });
  },

  createIngredient(data) {
    return axios
      .post(`${process.env.VUE_APP_ROOT_DOMAIN}/ingredients/`, data)
      .catch(error => {
        return Promise.reject(error);
      })
      .then(response => {
        return Promise.resolve(response.data);
      });
  },

  searchIngredients(searchTerm) {
    return axios
      .get(
        `${process.env.VUE_APP_ROOT_DOMAIN}/search/ingredients/${searchTerm}`
      )
      .catch(error => {
        return Promise.reject(error);
      })
      .then(response => {
        return Promise.resolve(response.data);
      });
  }
};
