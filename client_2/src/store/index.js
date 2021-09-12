import { createStore } from 'vuex'
import createPersistedState from "vuex-persistedstate";
import products from './modules/products.js'
import window from './modules/window.js'
import user from './modules/user.js'
export default createStore({
  modules: {
    products, window, user
  },
  plugins: [createPersistedState()],

})
