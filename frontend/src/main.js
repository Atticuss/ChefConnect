import "@fortawesome/fontawesome-free/css/all.min.css";
import "bootstrap-css-only/css/bootstrap.min.css";
import "mdbvue/lib/mdbvue.css";
import Vue from "vue";
import App from "./App";
import router from "./router";
import store from "./store";

Vue.config.productionTip = false;

import Notify from "mdbvue/lib/components/Notify";
Vue.use(Notify);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
