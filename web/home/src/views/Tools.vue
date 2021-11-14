<template>
    <div class="my_content">

        <div style="display: flex">
            <el-input
                    @change="refilldata"
                    placeholder="请输入内容"
                    v-model="input"
                    clearable>
            </el-input>
            <el-button @click="search" type="success" icon="el-icon-search" circle
                       style="margin-left: 5px;"></el-button>
            <!--            <el-button size="mini" type="success" icon="el-icon-search" style="margin-left: 5px; ">搜索</el-button>-->
        </div>

        <div style="text-align: left; margin-top: 12px" v-if="showMore">
            <el-button @click="addDrawer=true" size="mini" type="info" round>增加</el-button>
            <el-button @click="sortDrawer=true" size="mini" type="info" round>排序</el-button>
        </div>
        <el-drawer
                title="增加"
                align="center"
                :visible.sync="addDrawer"
                direction="ttb"
                :before-close="handleAddCan">
            <div>不可增加</div>

            <el-button style="margin-left: 12px" @click="handleAdd" type="success" icon="el-icon-check"
                       circle></el-button>
        </el-drawer>

        <el-drawer
                title="请选择"
                align="center"
                :visible.sync="sortDrawer"
                direction="ttb"
                :before-close="handleSort">

            <el-select v-model="sortValue"
                       placeholder="按字母排序">
                <el-option
                        v-for="item in sortOptions"
                        :key="item.value"
                        :label="item.value"
                        :value="item.value">
                </el-option>
            </el-select>
            <el-button style="margin-left: 12px" @click="handleSort" type="success" icon="el-icon-check"
                       circle></el-button>

        </el-drawer>


        <!--        分界线-->
        <el-divider content-position="right" class="el-divider--horizontal">
            <i @click="openMore" :class="dividerIcon"></i>
        </el-divider>

        <!--主题内容-->
        <el-collapse v-model="activeNames" @change="handleChange">

            <div :key='index' v-for="(val, key, index) in data">

                <el-collapse-item :name=val.index>
                    <template slot="title">
                        <div style="margin-left: 5%">{{val.title}}</div>
                    </template>
                    <div style="text-align: right">
                        <el-button @click="mydel(key)" size="mini" type="danger" icon="el-icon-delete"
                                   circle></el-button>
                        <el-button @click="mycopy(key); " size="mini" type="primary" icon="el-icon-copy-document"
                                   circle></el-button>
                        <el-button @click="mygo(key)" size="mini" type="warning" icon="el-icon-right" circle
                                   style="margin-right: 5%"></el-button>
                    </div>
                    <el-input
                            :autosize="{ minRows: 1, maxRows: 8}"
                            onfocus="this.select()"
                            style="margin-top: 5px"
                            type="textarea"
                            placeholder="被删除了 -_- "
                            v-model="val.url"
                    >
                    </el-input>
                </el-collapse-item>
            </div>
        </el-collapse>
    </div>
</template>

<script>
    export default {
        name: "Tools",
        data() {
            return {
                sortOptions: [{
                    value: "按字母排序"
                }, {
                    value: "按时间排序"
                }, {
                    value: "按使用频率"
                }, {
                    value: "按最近使用"
                }],
                sortValue: "按字母排序",
                addDrawer: false,
                sortDrawer: false,
                showMore: false,
                dividerIcon: "el-icon-caret-bottom",
                input: "",
                textarea: '',
                activeNames: [],
                data: [],

                allData: [{
                    index: "0",
                    url: "https://akarin.dev/WechatMomentScreenshot/",
                    title: "朋友圈转发截图生成工具"
                }, {
                    index: "1",
                    url: "https://www.google.com",
                    title: "google"
                }, {
                    index: "2",
                    url: "https://www.google.com/dsafasdfasdfasdfffffffffffffffff",
                    title: "test2"
                }, {
                    index: '3',
                    url: "https://www.baidu.com",
                    title: '百度'
                }]


            }
        },
        mounted: function () {
            this.refilldata();
        },
        methods: {
            handleAdd(){
                this.addDrawer = false
                this.$message({
                    type: 'success',
                    message: '增加成功!',
                    center: true,
                    showClose: true,
                });
            },
            handleAddCan() {
                this.$confirm("确认取消",{
                    message: '确认取消？' ,
                    center: true,
                    showClose: true,
                }).then(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消!',
                        center: true,
                        showClose: true,
                    });
                    this.addDrawer = false;
                }).catch(() => {
                });
            },
            handleSort() {
                this.sortDrawer = false
                this.$message({
                    type: 'success',
                    message: '已选 ' + this.sortValue,
                    center: true,
                    showClose: true,
                });
            },

            openMore() {
                // let tmp = this.dividerIcon;
                // tmp = (tmp === "el-icon-caret-top") ? "el-icon-caret-bottom" : "el-icon-caret-top";
                // this.dividerIcon = tmp;
                this.showMore ?
                    (this.dividerIcon = "el-icon-caret-bottom", this.showMore = false)
                    : (this.dividerIcon = "el-icon-caret-top", this.showMore = true)

            },
            refilldata() {
                if (this.input.trim(' ').length === 0) {
                    this.data = this.allData;
                    this.activeNames = []
                }
            },
            search() {
                this.data = this.allData;
                let value = this.input;
                let tmpData = [];
                for (let i = 0; i < this.data.length; i++) {
                    if (this.data[i].title.search(value) !== -1) {
                        tmpData.push(this.data[i])
                    } else if (this.data[i].url.search(value) !== -1) {
                        tmpData.push(this.data[i])
                        this.activeNames.push(this.data[i].index)
                    }
                }
                this.data = tmpData;
            },
            handleChange(val) {
                console.log(val);
            },
            mygo(index) {
                window.open(this.data[index].url);
            }
            , mycopy(index) {
                let _this = this;
                this.$copyText(this.data[index].url).then(function () {
                    _this.$message({
                        message: '复制成功',
                        type: 'success',
                        center: true,
                        showClose: true,
                    });
                }, function () {
                    _this.$message.error('复制失败');
                });

            }
            ,
            mydel(index) {

                this.$confirm('此操作将删除' + this.data[index].title + ', 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning',
                    center: true
                }).then(() => {
                    this.$message({
                        type: 'success',
                        message: '删除成功!',
                        center: true,
                        showClose: true,
                    });
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消删除',
                        center: true,
                        showClose: true,
                    });
                });
            }
        }
    }
</script>

<style scoped>
    .el-divider--horizontal {
        margin: 16px 0;
    }

</style>