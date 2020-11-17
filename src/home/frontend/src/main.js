import Vue from 'vue'
import App from './App.vue'
import router from './router'
import  './plugins/element.js'
import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.config.productionTip = false

import VueClipboard from 'vue-clipboard2'
Vue.use(VueClipboard)

Vue.prototype.axios = axios;
Vue.use(VueAxios, axios);

//配置全局的axios默认值（可选）

axios.defaults.baseURL = 'https://firego.cn';
axios.defaults.headers.common['Authorization'] = "AUTH_TOKEN";
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';

// router.js
router.beforeEach((to, from, next) => {
  document.title = "Firego";
  next()
})

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
