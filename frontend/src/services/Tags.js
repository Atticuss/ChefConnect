import axios from "axios";

export default {
  getTags() {
    return axios
      .get(`${process.env.VUE_APP_ROOT_DOMAIN}/tags/`)
      .catch(error => {
        return Promise.reject(error);
      })
      .then(response => {
        return Promise.resolve(response.data);
      });
  },

  createTag(data) {
    return axios
      .post(`${process.env.VUE_APP_ROOT_DOMAIN}/tags/`, data)
      .catch(error => {
        return Promise.reject(error);
      })
      .then(response => {
        return Promise.resolve(response.data);
      });
  }
};
