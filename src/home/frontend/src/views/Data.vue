<template>
    <div class="my_content">
        <el-divider>Search Data</el-divider>
        <div style="display: flex;">
            <el-input
                    placeholder="搜索"
                    v-model="searchData"
                    clearable>
            </el-input>
            <el-button @click="delSearchData" style="margin-left: 20px" type="success">搜索
            </el-button>

        </div>
        <el-table

                :data="searchRet"
                :default-sort="{prop: 'grade', order: 'descending'}"
                stripe
                style="width: 100%"
                :height="searchHeight">
            <el-table-column
                    prop="grade"
                    label="匹配度"
                    width="200">
            </el-table-column>
            <el-table-column
                    prop="post_time"
                    label="日期"
                    width="140">
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
                    width="80">
            </el-table-column>
            <el-table-column
                    prop="text"
                    label="文章">
            </el-table-column>
        </el-table>

        <br><br> <br> <br>
        <el-divider>All Data
        </el-divider>

        <div style="width: 100%; text-align: right">
            <el-button @click="delAllData" style="display: inline-block" type="success">加载</el-button>
        </div>

        <el-table
                :data="allData"
                :default-sort="{prop: 'post_time', order: 'descending'}"
                stripe
                style="width: 100%"
                :height="allHeight">
            <el-table-column
                    prop="post_time"
                    label="日期"
                    width="140">
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
                    width="80">
            </el-table-column>
            <el-table-column
                    prop="text"
                    label="文章">
            </el-table-column>
        </el-table>

        <br> <br> <br> <br>
        <el-divider>WordSet</el-divider>
        <div style="width: 100%; text-align: right">
            <el-button @click="delWordSet" style="display: inline-block" type="success">加载</el-button>
        </div>
        <div>
            <el-table
                    :data="wordSet"
                    style="width: 100%"
                    :height="wordSetHeight"
                    :row-class-name="tableRowClassName"
                    :default-sort="{prop: 'num', order: 'descending'}"
            >
                <el-table-column
                        prop="word"
                        label="word"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="num"
                        label="num"
                        width="180">
                </el-table-column>
            </el-table>
            <el-divider>Over</el-divider>
        </div>


    </div>
</template>

<script>
    import axios from 'axios'

    export default {
        name: "Data.vue",
        watch: {
            searchData(val, oldValue) {
                console.log(oldValue);
                if (val === "") {
                    this.searchRet = [],
                        this.searchHeight = "20vh"
                }
            }

        },
        methods: {
            delAllData() {
                axios
                    .get('http://firego.cn/beibei/api/data')
                    .then(response => {
                        let obj = response.data.data
                        let arr = [];
                        Object.keys(obj).forEach(v => {
                            // console.log(obj[v])
                            obj[v]["post_time"] = obj[v]["post_time"].substr(5);
                            obj[v]["comments"] = obj[v]["comments"].substr(4);
                            obj[v]["good_nums"] = obj[v]["good_nums"].substr(4);
                            obj[v]["resend_nums"] = obj[v]["resend_nums"].substr(4);
                            obj[v]["send_tool"] = obj[v]["send_tool"].substr(5);

                            arr.push(obj[v])
                        });

                        this.allData = arr
                        this.allHeight = "80vh"
                    });

            },
            delWordSet() {

                axios
                    .get('http://firego.cn/beibei/api/wordset')
                    .then(response => {
                        let obj = response.data.data.data;
                        let arr = [];
                        Object.keys(obj).forEach(v => {
                            // console.log(obj[v])
                            let o = {}
                            o["word"] = v;
                            o["num"] = obj[v]
                            arr.push(o)
                        });
                        this.wordSet = arr
                        this.wordSetHeight = "80vh"
                    })
            },

            delSearchData() {
                if (this.searchData === "") {
                    alert("请输入数据！")
                    return
                }
                let param = "data=" + this.searchData;
                axios
                    .get(`http://firego.cn/beibei/api/search?` + param)
                    .then(response => {
                        let obj = response.data.data
                        let arr = [];
                        Object.keys(obj).forEach(v => {
                            // console.log(obj[v])
                            obj[v]["post_time"] = obj[v]["post_time"].substr(5);
                            obj[v]["comments"] = obj[v]["comments"].substr(4);
                            obj[v]["good_nums"] = obj[v]["good_nums"].substr(4);
                            obj[v]["resend_nums"] = obj[v]["resend_nums"].substr(4);
                            obj[v]["send_tool"] = obj[v]["send_tool"].substr(5);
                            // obj[v]["grade"] = obj[v]["grade"].substr(0, 8);
                            arr.push(obj[v])
                        });

                        this.searchRet = arr;
                        if (arr.length === 0) {
                            this.searchRet.push(
                                {
                                    post_time: '0',
                                    comments: "0",
                                    good_nums: "0",
                                    resend_nums: "0",
                                    send_tool: "无",
                                    text: '一条数据都没搜索到',
                                    grade: "0"
                                }
                            );
                            this.searchHeight = "20vh"
                        } else {
                            this.searchHeight = "80vh"
                        }
                    });

            },


            tableRowClassName({row, rowIndex}) {
                if (row) {
                    row += ""
                }
                if (rowIndex === 1) {
                    return 'warning-row';
                } else if (rowIndex === 3) {
                    return 'success-row';
                }
                return '';
            }
        },
        data() {
            return {
                searchHeight: "20vh",
                allHeight: "20vh",
                wordSetHeight: "20vh",
                searchRet: [],
                searchData: '',
                wordSet: [
                    {
                        word: "请点击加载",
                        num: "0"
                    }
                ],

                allData: [
                    {
                        post_time: '0',
                        comments: "0",
                        good_nums: "0",
                        resend_nums: "0",
                        send_tool: "无",
                        text: '请点击加载',
                    }
                ]
            }
        },
        mounted() {


        }
    }
</script>

<style scoped>

</style>