import axios from "axios";

export default {
  getCategories() {
    return axios.get("/api/categories").then(response => {
      return response.data.categories;
    });
  },

  createCategory(data) {
    return axios.post("/api/categories", data).then(response => {
      return response.data;
    });
  }
};
