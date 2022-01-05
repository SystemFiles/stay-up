import Vue from "vue";
import Vuex from "vuex";

import add from "./add.module";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    add
  }
});
