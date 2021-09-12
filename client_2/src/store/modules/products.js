import axios from 'axios'

const state = {
  allProduct: []
}
const getters = {
  allProducts: (state) => { console.log("getter");return state.allProduct }
}
const mutations = {
  LOAD_ALL_PRODUCTS: (state, allProduct) => {
    state.allProduct = allProduct
  }
}
const actions = {
  async act_loadAllProducts({ commit }) {
    console.log("hello");
    const response = await axios.get("/home")
    console.log(response.data);
    commit('LOAD_ALL_PRODUCTS', await response.data)
  }
}
export default {
  state,
  getters,
  actions,
  mutations
};