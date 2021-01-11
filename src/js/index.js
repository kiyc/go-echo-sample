import Vue from "vue";
import VueRouter from "vue-router";
import Vuetify from "vuetify";
import Vuex from "vuex";
import Filters from './filters';

// Vue Router
Vue.use(VueRouter);

import Login from "./components/Login.vue";
import Dashboard from "./components/Dashboard";
import User from "./components/User";

const routes = [
  { path: "/", name: "login", component: Login, meta: { title: "Login" } },
  { path: "/dashboard", name: "dashboard", component: Dashboard, meta: { title: "Dashboard" } },
  { path: "/users", name: "users", component: User, meta: { title: "Users" } },
];

const router = new VueRouter({
  mode: "history",
  routes,
});

router.afterEach((to, from) => {
  if (to.meta && to.meta.title) {
    document.title = to.meta.title;
  }
});

// Vuetify
Vue.use(Vuetify);
const opts = {};

const vuetify = new Vuetify(opts);

// Vuex
Vue.use(Vuex);

const store = {
  state: {
    authToken: null
  },
  mutations: {
    setAuthToken(state, authToken) {
      state.authToken = authToken;
    }
  },
  actions: {
    logout({ commit }) {
      commit("setAuthToken", null);
      document.location.href = "/";
    },
    getAuthHeaders({ state }) {
      const headers = {
        headers: {
          Authorization: `Bearer ${state.authToken}`
        }
      };
      return headers;
    },
  }
};

Vue.use(Filters);

new Vue({
  el: "#app",
  router,
  vuetify,
  store: new Vuex.Store(store),
});