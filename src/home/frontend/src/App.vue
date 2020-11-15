<template>
    <div id="app">

        <div id="nav">
            <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect"
                     style="display: flex;">
                <el-menu-item index="1">
                    Home
                </el-menu-item>
                <el-menu-item index="2">
                    Marks
                </el-menu-item>
                <el-menu-item index="3">
                    About
                </el-menu-item>

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
                moveStart: true,
                activeIndex: '1',
                startX: 0,
                startY: 0
            }
        },
        methods: {
            handleSelect(key, keyPath) {
                console.log(key, keyPath);
                this.move2page(key, 0)
            },
            move2page(index, direction) {
                if(!this.moveStart) return;
                this.moveStart = false;
                switch (index) {
                    case "1":
                        if (direction === 0 && this.activeIndex !== index) {
                            this.$router.push("/");
                            this.activeIndex = '1';
                        } else if (direction === 1) {
                            this.$router.push("/marks");
                            this.activeIndex = '2';
                        }
                        break;
                    case "2":
                        switch (direction) {
                            case -1:
                                this.$router.push("/");
                                this.activeIndex = '1';
                                break;
                            case 0:
                                if (this.activeIndex !== index)
                                    this.$router.push("/marks");
                                this.activeIndex = '2';
                                break;
                            case 1:
                                this.$router.push("/about");
                                this.activeIndex = '3';
                                break;
                        }
                        break;
                    case '3':
                        if (direction === 0 && this.activeIndex !== index) {
                            this.$router.push("/about");
                            this.activeIndex = '3';
                        } else if (direction === -1) {
                            this.$router.push("/marks");
                            this.activeIndex = '2';
                        }
                        break;
                }
            }

        },
        mounted() {
            let that = this;
            document.getElementById('app').addEventListener('touchstart', function (e) {
                console.log('touchstart:', e)
                this.startX = e.changedTouches[0].pageX;
                this.startY = e.changedTouches[0].pageY
            })
            document.getElementById('app').addEventListener('touchend', function (e) {
                console.log('touchend:', e)
                that.moveStart = true
            })

            document.getElementById('app').addEventListener('touchmove', function (e) {
                console.log('touchmove:', e)
                if (e.changedTouches.length) {
                    let moveEndX = e.changedTouches[0].pageX
                    let moveEndY = e.changedTouches[0].pageY
                    let X = moveEndX - this.startX
                    let Y = moveEndY - this.startY
                    if (Math.abs(X) > Math.abs(Y) && X > 0) {
                        console.log("left 2 right", X, Y);
                        if (X > 30) {
                            that.move2page(that.activeIndex, -1)
                        }
                    } else if (Math.abs(X) > Math.abs(Y) && X < 0) {
                        console.log("right 2 left", X, Y);
                        if (X < -31) {
                            that.move2page(that.activeIndex, 1);
                        }
                    } else if (Math.abs(Y) > Math.abs(X) && Y > 0) {
                        console.log("top 2 bottom", X, Y);
                    } else if (Math.abs(Y) > Math.abs(X) && Y < 0) {
                        console.log("bottom 2 top", X, Y);
                    } else {
                        console.log("just touch", X, Y);
                        that.moveStart = true;
                    }
                }

            })
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
        height: 100%;
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
