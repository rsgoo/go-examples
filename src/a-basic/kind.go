package main

import (
	"sort"
	"fmt"
)

type HeroKind int

const (
	None     HeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

type Heroes []*Hero

func (h Heroes) Len() int {
	return len(h)
}

func (h Heroes) Less(i, j int) bool {

	//如果英雄的分类不一致时，优先对分类进行排序
	if h[i].Kind != h[j].Kind {
		return h[i].Kind < h[j].Kind
	}

	//默认按英雄名称字符升序排序
	return h[i].Name < h[j].Name
}

func (h Heroes) Swap(i, j int) {

	h[i], h[j] = h[j], h[i]
}

func main() {
	heroes := Heroes{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽 ", Tank},
		&Hero{"诸葛亮", Mage},
	}

	sort.Slice(heroes, func(i, j int) bool {
		//如果英雄的分类不一致时，优先对分类进行排序
		if heroes[i].Kind != heroes[j].Kind {
			return heroes[i].Kind < heroes[j].Kind
		}

		//默认按英雄名称字符升序排序
		return heroes[i].Name < heroes[j].Name
	})

	for key, value := range heroes {
		fmt.Println(key)
		fmt.Println(*value)
	}
}
