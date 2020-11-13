<template>
    <div id="app">

        <div id="nav">
            <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
                <el-menu-item index="1">
                    Home
                </el-menu-item>
                <el-menu-item index="2">
                    Tools
                </el-menu-item>
                <el-menu-item index="3">DDL</el-menu-item>
                <el-submenu index="4">
                    <template slot="title">Test</template>
                    <el-menu-item index="2-1">选项1</el-menu-item>
                    <el-menu-item index="2-2">选项2</el-menu-item>
                    <el-menu-item index="2-3">选项3</el-menu-item>
                    <el-submenu index="2-4">
                        <template slot="title">选项4</template>
                        <el-menu-item index="2-4-1">选项1</el-menu-item>
                        <el-menu-item index="2-4-2">选项2</el-menu-item>
                        <el-menu-item index="2-4-3">选项3</el-menu-item>
                    </el-submenu>
                </el-submenu>
            </el-menu>
        </div>

        <router-view/>

        <!--        // 不能 :right = 10% 带单位？-->
        <el-backtop style="right: 10%">
            <div
                    style="{        height: 100%;        width: 100%;        background-color: #f2f5f6;
        box-shadow: 0 0 6px rgba(0,0,0, .12);        text-align: center;        line-height: 40px;        color: #1989fa;      }"
            >
                UP
            </div>
        </el-backtop>


    </div>
</template>

<script>

    export default {
        data() {
            return {
                activeIndex: '1'
            }
        },
        methods: {
            handleSelect(key, keyPath) {
                console.log(key, keyPath);
                switch (key) {
                    case "1":
                        this.$router.push("/home");
                        break;
                    case  "2":
                        this.$router.push("/about");
                        break;

                }
            }
        },
        mounted() {
            console.log(this);

        },
        components: {}
    }

</script>

<style lang="scss">
    #app {
        font-family: Avenir, Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
    }

    #nav {

        padding-left: 10%;
        padding-right: 10%;

        a {
            font-weight: bold;
            color: #2c3e50;

            &.router-link-exact-active {
                color: #42b983;
            }
        }
    }
</style>
