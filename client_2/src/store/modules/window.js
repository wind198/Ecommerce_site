
const state = {
    scrollY: [0]
}
const getters = {
    getScrollY: (state) => { return state.scrollY }
}
const mutations = {
    UPDATE_SCROLL_Y: (state, newValue) => {
        console.log("old scrollY", state.scrollY[state.scrollY.length-1]);
        state.scrollY.shift();
        state.scrollY.push(newValue);
        console.log("new scrollY", state.scrollY[state.scrollY.length-1]);

    }
}
const actions = {
    act_updateScrollY({ commit }, newValue) {
        console.log("hello");
        commit('UPDATE_SCROLL_Y', newValue)
    }
}
export default {
    state,
    getters,
    actions,
    mutations
};