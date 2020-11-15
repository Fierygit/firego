import Vue from 'vue'
import App from './App.vue'
import router from './router'
import  './plugins/element.js'

Vue.config.productionTip = false

import VueClipboard from 'vue-clipboard2'
Vue.use(VueClipboard)

// router.js
router.beforeEach((to, from, next) => {
  document.title = "Firego";
  next()
})

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
