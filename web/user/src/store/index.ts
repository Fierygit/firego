import { createStore } from 'vuex'

export default createStore({
  state: {

    isLoginPage: false
  },
  mutations: {
    changeIP(state) {
      state.isLoginPage = !state.isLoginPage
    }
    ,
    setIsLogin(state, payload){
      state.isLoginPage = payload;
    }
  },
  actions: {
  },
  modules: {
  }
})
