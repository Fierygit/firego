<!--
 * @Author: Firefly
 * @Date: 2021-04-02 16:41:21
 * @Descripttion: 
 * @LastEditTime: 2021-04-04 22:45:42
-->

<template>
  <div style="margin-top:10%">请输入你的动态密码</div>
  <div style="margin-top:10px">
    <el-row>
      <el-input placeholder="请输入内容" v-model="input" clearable> </el-input>
      <el-button type="primary" @click="upload()">提交</el-button>
    </el-row>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import axios from "axios";

export default defineComponent({
  name: "login",
  setup() {
    let input = ref("");
    function upload(): void {
      console.log(input.value);
      input.value = "";
      axios.get("https://www.google.com").then((response) => {
        console.log(response.data);
      });
    }
    return {
      input,
      upload,
    };
  },
});
</script>

<style></style>
