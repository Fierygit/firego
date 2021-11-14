<template>
    <div class="my_content">
        <br/>
        <font size="49px">Firego</font>
        <p>
            this is a project used to study vue
        </p>

        <br/><br/><br/>
        <el-divider>image</el-divider>

        <div class="demo-image__preview">
            <el-image
                    style="width: 100px; height: 100px; padding: 5px"
                    :src="url"
                    :preview-src-list="srcList">
            </el-image>

        </div>


    </div>
</template>

<script>
    // @ is an alias to /src


    export default {
        name: 'Home',
        data() {
            return {
                url: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
                srcList: [
                    'https://fuss10.elemecdn.com/8/27/f01c15bb73e1ef3793e64e6b7bbccjpeg.jpeg',
                    'https://fuss10.elemecdn.com/1/8e/aeffeb4de74e2fde4bd74fc7b4486jpeg.jpeg'
                ]
            };
        },
        methods: {}

        ,
        components: {}
    }
</script>

<style>
    .my_content {
        margin-left: 10%;
        margin-right: 10%;
        margin-top: 15px;
    }

</style>