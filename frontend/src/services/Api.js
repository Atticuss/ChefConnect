import axios from "axios";

var url2path = function(url) {
  return url;
};

/*
axios.interceptors.request.use(error => {
  Promise.reject(error);
});
*/

axios.interceptors.response.use(
  function(response) {
    return Promise.resolve(response);
  },
  async function(error) {
    const originalRequest = error.config;
    var path = url2path(originalRequest.url);

    if (error.response.status === 401 && path === "/auth/refresh-token") {
      return Promise.reject(error);
    }

    const refreshToken = window.localStorage.getItem("refresh");
    if (!refreshToken) {
      return Promise.reject(error);
    }

    if (error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;
      const res = await axios.post(
        `${process.env.VUE_APP_ROOT_DOMAIN}/auth/refresh-token`,
        { refreshToken }
      );
      if (res.status === 200) {
        window.localStorage.setItem("jwt", res.data.authToken);
        window.localStorage.setItem("refresh", res.data.refreshToken);

        axios.defaults.headers.common.Authorization = res.data.authToken;
        originalRequest.headers.Authorization = res.data.authToken;

        return axios(originalRequest);
      }
    }

    return Promise.reject(error);
  }
);
