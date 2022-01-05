import { SvcService } from "../common/api.service";
import { POST_SERVICE } from "./actions.type";
import { ADDSVC } from "./mutations.type";

const state = {
  serviceName: "",
  serviceHost: "",
  serviceDesc: "",
  servicePort: "",
  serviceProtocol: "",
  serviceTimeout: 0
};

const getters = {
  serviceName(state) {
    return state.serviceName;
  },
  serviceHost(state) {
    return state.serviceHost;
  },
  servicePort(state) {
    return state.servicePort;
  },
  serviceDesc(state) {
    return state.serviceDesc;
  },
  serviceProtocol(state) {
    return state.serviceProtocol;
  },
  serviceTimeout(state) {
    return state.serviceTimeout;
  }
};

const actions = {
  [POST_SERVICE](data) {
    return SvcService.post(data);
  }
};

/* eslint no-param-reassign: ["error", { "props": false }] */
const mutations = {
  [ADDSVC](state, name, host, port, desc, protocol, timeout) {
    state.serviceName = name;
    state.serviceHost = host;
    state.servicePort = port;
    state.serviceDesc = desc;
    state.serviceProtocol = protocol;
    state.serviceTimeout = timeout;
  }
};

export default {
  state,
  getters,
  actions,
  mutations
};
