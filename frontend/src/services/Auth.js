import axios from "axios";

export default {
  authRequest(rootVueInstance, data) {
    return axios
      .post(`${process.env.VUE_APP_ROOT_DOMAIN}/auth/login`, data)
      .then(response => {
        let eventData = {
          authToken: response.data.authToken,
          refreshToken: response.data.refreshToken
        };

        rootVueInstance.$emit("successful-auth", eventData);

        window.localStorage.setItem("jwt", response.data.authToken);
        window.localStorage.setItem("refresh", response.data.refreshToken);

        axios.defaults.headers.common.Authorization = response.data.authToken;
      })
      .catch(err => {
        return Promise.reject(err);
      });
  },

  logout() {
    window.localStorage.removeItem("refresh");
    window.localStorage.removeItem("jwt");
    axios.defaults.headers.common.Authorization = null;
  }
};
