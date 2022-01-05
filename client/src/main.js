import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import "./registerServiceWorker";
import ApiService from "./common/api.service";

Vue.config.productionTip = false;

ApiService.init();

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
