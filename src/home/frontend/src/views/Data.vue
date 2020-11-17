<template>
    <div class="my_content">
        <el-table
                :data="tableData"
                stripe
                style="width: 100%"
                height="500">
            <el-table-column
                    prop="post_time"
                    label="日期"
                    width="150">
            </el-table-column>
            <el-table-column
                    prop="comments"
                    label="评论"
                    width="50">
            </el-table-column>
            <el-table-column
                    prop="good_nums"
                    label="点赞"
                    width="50">
            </el-table-column>
            <el-table-column
                    prop="resend_nums"
                    label="转发"
                    width="50">
            </el-table-column>
            <el-table-column
                    prop="send_tool"
                    label="发布工具"
            width="100">
            </el-table-column>
            <el-table-column
                    prop="text"
                    label="文章">
            </el-table-column>
        </el-table>


    </div>
</template>

<script>
    import axios from 'axios'
    export default {
        name: "Data.vue",
        data() {
            return {
                tableData: [{
                    post_time: '2016-05-02 23:23',
                    comments: "1",
                    good_nums: "2",
                    resend_nums: "3",
                    send_tool: "iphone",
                    text: '收到好的时候生的话阿斯顿发的说法斯蒂芬上海市生的话阿斯顿发的说法斯蒂芬上海市都是发生的话阿斯顿发的说法斯蒂芬上海市普陀区金沙江路 1518 弄',
                }]
            }
        },
        mounted () {
            axios
                .get('http://firego.cn/beibei/api/data')
                .then(response => (console.log(response)))
        }
    }
</script>

<style scoped>

</style>