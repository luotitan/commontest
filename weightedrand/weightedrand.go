package main

import (
	"log"

	"github.com/go-ego/murmur"
	"github.com/qiuker521/weightedrand"
)

func main() {
	log.Println("随机选一个语言")
	选语言()
	log.Println("使用一致性hash作弊")
	一致性哈希()
}
func 选语言() {
	c := weightedrand.Chooser{}
	var choices = []weightedrand.Choice{}
	choices = append(choices, weightedrand.Choice{"java", 0})
	choices = append(choices, weightedrand.Choice{"rust", 0})
	choices = append(choices, weightedrand.Choice{"python", 1})
	choices = append(choices, weightedrand.Choice{"go", 9})
	c.NewChooser(choices...)
	for i, v := range choices {
		log.Println("选项", i+1, "：", v.Item)
	}
	log.Println("我应该学", c.Pick(), "语言")
}

func 一致性哈希() {
	c := weightedrand.Chooser{}
	var choices = []weightedrand.Choice{}
	choices = append(choices, weightedrand.Choice{"java", 1})
	choices = append(choices, weightedrand.Choice{"rust", 1})
	choices = append(choices, weightedrand.Choice{"python", 1})
	choices = append(choices, weightedrand.Choice{"go", 1})
	for i, v := range choices {
		log.Println("选项", i+1, "：", v.Item, "权重:", v.Weight)
	}

	c.NewChooser(choices...)
	var hash = "一致性哈希永远会选java"
	log.Println("然而事实上是：虽然不作弊，但是", hash)
	log.Println("不信的话我选择给你看：", c.PickByHash(float64(murmur.Sum32(hash))))
}