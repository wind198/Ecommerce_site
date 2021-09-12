const state = {
    userLogged: false,
    credential: {
        email: '',
        jwtToken: ''
    }
}
const getters = {
    getStatus: (state) => state.userLogged,
    getEmailCredential:(state)=>state.credential
}
const mutations = {
    LOGIN: (state, payload) => {
        console.log("payload",payload);
        state.userLogged = true;
        state.credential.email = payload.email;
        state.credential.jwtToken = payload.jwtToken;
    },
    LOGOUT: (state) => {
        state.userLogged = false;
        state.credential.email = '';
        state.credential.jwtToken = '';
    }
}
const actions = {
    act_login: ({ commit }, email, jwtToken) => {
        commit("LOGIN", email, jwtToken)
    },
    act_logout: ({ commit }) => {
        commit("LOGOUT")
    }
}
export default {
    state,
    getters,
    actions,
    mutations
};