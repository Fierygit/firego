/*
 * @Author: Firefly
 * @Date: 2020-11-07 17:29:09
 * @Descripttion: 
 * @LastEditTime: 2020-11-08 14:19:03
 */
// 霸都丶傲天 2019年10月10日 https://github.com/AJLoveChina/birthday
var config = {
    // 句子的长度可以任意， 你可以写十句话， 二十句话都可以
    // 每句话尽量不要超过15个字,不然展示效果可能不太好
    texts: [
        "我想给贝贝最好的祝福",
        "无奈写不出太多华丽句子",
        "绞尽脑汁，到底如何才能让贝贝开心呢？",
        "翻遍歌单，寻找了很多动人的旋律",
        "才发现想说的祝福歌词里原来如此之多",
        "送给我",      
        "心爱的贝贝",
        "愿你的身后总有力量",
        "愿你成为自己的太阳",
        "愿你永驻时光爱上彼此的模样",
        "愿你开开心心每一天!!!",  
        "......"
        
        
    ],
    /**
     * imgs 可以不填, 但是如果要填写的话必须遵循下面的格式
     * "对应上面的文字, 要完全一样" : "图片地址, 可以把图片放在imgs文件夹中"
     * 例如
     * "心爱的小可爱": "./imgs/xiaokeai.jpg"
     *
     * 如果不要图片的话, 直接在每行开头写两个斜杠注释即可, 例如下面的 "今天是你的生日" 的图片就不会展示了:)
     * Tip: 图片最好用正方形or接近正方形, 看起来效果更好
     */
    imgs: {
        "心爱的贝贝": "./imgs/xiaokeai.png",
        // "今天是你的生日": "./imgs/birthday.jpg",
    },
    // 按钮文字描述, 以下是默认的按钮文字，英文的，您可以改成你喜欢的文字
    desc: {
        turn_on: "开灯",
        play: "music",
        bannar_coming: "go",
        balloons_flying: "好像少点东西",
        cake_fadein: "蛋糕",
        light_candle: "fire",
        wish_message: "生日快乐",
        story: "A MESSAGE FOR beibei",
    }
};
