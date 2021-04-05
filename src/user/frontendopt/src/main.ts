/*
 * @Author: Firefly
 * @Date: 2021-03-31 21:20:53
 * @Descripttion: 
 * @LastEditTime: 2021-04-04 22:41:08
 */
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css';


createApp(App).use(store).use(router)
    .use(ElementPlus, { size: 'small', zIndex: 3000 }).mount('#app')
