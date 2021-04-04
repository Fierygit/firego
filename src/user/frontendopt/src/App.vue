<!--
 * @Author: Firefly
 * @Date: 2021-03-31 21:20:53
 * @Descripttion: 
 * @LastEditTime: 2021-04-04 22:05:25
-->
<template>
  <div id="nav" v-if="!$store.state.isLoginPage">
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link>
  </div>

  <router-view />
</template>

<script lang="ts">
import { defineComponent, watch } from "vue";
import { useStore } from "vuex";
import { useRoute } from "vue-router";

export default defineComponent({
  name: "APP",
  setup(ctx) {
    console.log(ctx);
    watchPage(useStore(), useRoute());
    return {};
  },
});

function watchPage(store: any, route: any): void {
  watch(
    () => route.path,
    (newPath) => {
      console.log("监听到变化");
      newPath == "/login"
        ? store.commit("setIsLogin", true)
        : store.commit("setIsLogin", false);
      console.log(newPath, store.state.isLoginPage);
    }
  );
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
