import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
// import About from/* webpackChunkName: "about" */ '../views/About.vue'
Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/home',
        redirect: "/"
    },

    {
        path: "/marks",
        name: 'marks',
        component: () => import(/* webpackChunkName: "about" */ '../views/Tools.vue')

    },

    {
        path: '/about',
        name: 'about',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
    },
    {
        path: '/data',
        name: 'data',
        component: ()=>import('../views/Data.vue')
    },
    {
        path: '*',
        redirect: "/about"

    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router
