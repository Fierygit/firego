/*
 * @Author: Firefly
 * @Date: 2020-11-16 22:50:55
 * @Descripttion:
 * @LastEditTime: 2020-11-16 23:16:21
 */
package beibei

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"

	jiebago "github.com/wangbin/jiebago"
)

var seg jiebago.Segmenter

func print(data interface{}) {
	fmt.Println(data)
}
func init() {
	seg.LoadDictionary("dict.txt")
}

func example() {
	fmt.Print("【全模式】：")
	print(seg.CutAll("我来到北京清华大学"))

	fmt.Print("【精确模式】：")
	print(seg.Cut("我来到北京清华大学", false))

	fmt.Print("【新词识别】：")
	print(seg.Cut("他来到了网易杭研大厦", true))

	fmt.Print("【搜索引擎模式】：")
	print(seg.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", true))
}

func getInitData() []map[string]interface{} {
	b, err1 := ioutil.ReadFile("beibei/weibodata/dealed.json")
	if err1 != nil {
		fmt.Println("read fail", err1)
	}
	var m interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
	}

	// fmt.Println(reflect.TypeOf(m))
	dataMap := m.(map[string]interface{})
	var dataList []map[string]interface{}
	for _, v := range dataMap {
		dataLine := v.(map[string]interface{})
		dataList = append(dataList, dataLine)
	}
	return dataList
}

var wordSet = make(map[string]int)
var wordMatrix [][]string
var word2idx = make(map[string]int)
var delWord = [...]string{"你", "我", "?", "？", ",", ".", "。", " ", "，", "!", "！", "吗", "什么"}
var vecex [][]int
var dataList []map[string]interface{}

func dealData() {
	dataList = getInitData()
	for _, v := range dataList {
		var wordList []string
		wordListCH := seg.CutForSearch(v["text"].(string), true)
		for word := range wordListCH {
			wordList = append(wordList, word)
			wordSet[word]++
			nameRune := []rune(word)
			for i := 0; i < len(nameRune); i++ {
				wordSet[string(nameRune[i:i+1])]++
			}
		}
		wordMatrix = append(wordMatrix, wordList)
	}
	for key := range wordSet {
		for i := 0; i < len(delWord); i++ {
			if key == delWord[i] {
				delete(wordSet, key)
			}
		}
	}
	// print(wordSet)
	// print(wordMatrix)
	// 向量化
	index := 0
	for key := range wordSet {
		word2idx[key] = index
		// if index == 0 {
		// 	print(key)
		// }
		index++
	}

	for _, wordList := range wordMatrix {
		vecLen := len(wordSet)
		vec := make([]int, vecLen)
		for i := 0; i < len(wordList); i++ {
			for j := 0; j < len(delWord); j++ {
				if wordList[i] == delWord[j] {
					continue
				}
			}
			vec[word2idx[wordList[i]]] = 1
		}
		vecex = append(vecex, vec)
	}

}

var initData bool = false

//SearchDataCos s
func SearchDataCos(inputData string) (map[int]float64, []map[string]interface{}) {
	if !initData {
		dealData()
		initData = true
	}

	var inputWords []string
	var inputWordSet = make(map[string]int)
	wordListCH := seg.CutForSearch(inputData, true)
	for word := range wordListCH {
		inputWords = append(inputWords, word)
		inputWordSet[word]++
		nameRune := []rune(word)
		for i := 0; i < len(nameRune); i++ {
			inputWordSet[string(nameRune[i:i+1])]++
		}
	}

	vecLen := len(wordSet)
	inVec := make([]int, vecLen)
	for key := range inputWordSet {
		if _, ok := word2idx[key]; !ok {
			continue
		}
		for i := 0; i < len(delWord); i++ {
			if key == delWord[i] {
				continue
			}
		}
		inVec[word2idx[key]] = 1
	}
	// print(inVec)
	var ans = make(map[int]float64)
	for index, v := range vecex {
		value := Cosine(v, inVec)
		if value != 0 {
			ans[index] = value
		}
	}
	return ans, dataList
}

//Cosine c
func Cosine(aa []int, bb []int) float64 {

	var a = make([]float64, len(aa))
	var b = make([]float64, len(bb))

	for i := 0; i < len(aa); i++ {
		if aa[i] == 1 {
			a[i] = 1.0
		}
		if bb[i] == 1 {
			b[i] = 1.0
		}
	}
	top := 0.0
	for i := 0; i < len(a); i++ {
		top += a[i] * b[i]
	}
	r := 0.0
	l := 0.0
	for _, i := range a {
		l += i * i
	}
	l = math.Sqrt(l)
	for _, i := range b {
		r += i * i
	}
	r = math.Sqrt(r)
	// fmt.Print(top / l * r)
	// fmt.Print(" ")
	// fmt.Print(top)
	// fmt.Print(" ")
	if top/l*r == 0 {
		return 0.0
	}
	return top / l * r
}
