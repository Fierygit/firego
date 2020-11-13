<template>
    <div class="home">
        <el-container :top="20">
            <el-header style="height: 20px;text-align: center; font-size: 12px">

                <div @click="drawer = true"  style="margin-left: 16px;">
                    增加
                </div>
                <el-drawer
                        title="我是标题"
                        :visible.sync="drawer"
                        :with-header="false">
                    <span>我来啦!</span>
                </el-drawer>

            </el-header>
            <el-main>
                <el-collapse v-model="activeName" accordion>
                    <el-collapse-item title="一致性 Consistency" name="1">
                        <div>与现实生活一致：与现实生活的流程、逻辑保持一致，遵循用户习惯的语言和概念；</div>
                        <div>在界面中一致：所有的元素和结构需保持一致，比如：设计样式、图标和文本、元素的位置等。</div>
                    </el-collapse-item>
                    <el-collapse-item title="反馈 Feedback" name="2">
                        <div>控制反馈：通过界面样式和交互动效让用户可以清晰的感知自己的操作；</div>
                        <div>页面反馈：操作后，通过页面元素的变化清晰地展现当前状态。</div>
                    </el-collapse-item>
                    <el-collapse-item title="效率 Efficiency" name="3">
                        <div>简化流程：设计简洁直观的操作流程；</div>
                        <div>清晰明确：语言表达清晰且表意明确，让用户快速理解进而作出决策；</div>
                        <div>帮助用户识别：界面简单直白，让用户快速识别而非回忆，减少用户记忆负担。</div>
                    </el-collapse-item>
                    <el-collapse-item title="可控 Controllability" name="4">
                        <div>用户决策：根据场景可给予用户操作建议或安全提示，但不能代替用户进行决策；</div>
                        <div>结果可控：用户可以自由的进行操作，包括撤销、回退和终止当前操作等。</div>
                    </el-collapse-item>
                </el-collapse>
            </el-main>
            <el-footer>Footer</el-footer>
        </el-container>




        <el-button type="text" @click="table = true">打开 Drawer</el-button>
        <el-drawer
                title="我嵌套了表格!"
                :visible.sync="table"
                direction="rtl"
                size="50%">
            <el-table :data="gridData">
                <el-table-column property="date" label="日期" width="150"></el-table-column>
                <el-table-column property="name" label="姓名" width="200"></el-table-column>
                <el-table-column property="address" label="地址"></el-table-column>
            </el-table>
        </el-drawer>

    </div>
</template>

<script>
    // @ is an alias to /src


    export default {
        name: 'Home',
        data() {
            return {
                drawer: false,
                activeName: '1',
                table: false,
                gridData: [{
                    date: '2016-05-02',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1518 弄'
                }, {
                    date: '2016-05-04',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1518 弄'
                }, {
                    date: '2016-05-01',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1518 弄'
                }, {
                    date: '2016-05-03',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1518 弄'
                }],
            };
        },
        methods: {}

        ,
        components: {}
    }
</script>

<style scoped>
    .home {
        margin-left: 10%;
        margin-right: 10%;
        margin-top: 15px;
    }
</style>