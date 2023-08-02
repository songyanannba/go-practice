package eliminate

import (
	"fmt"
	"slot-server/global"
	"slot-server/utils/helper"
)

type LevelEvent struct {
	Collect   int // 收集的数量
	SkillMul  int //倍率
	WildCount int
}

func (t *Table) FillSkillTags() {
	if t.SkillTags != nil {
		return
	}
	var count int
	t.SkillTags = make(map[[2]int]int, 5)

	skillNum := helper.RandInt(5) + 1   //随机技能个数
	arrs := helper.ListToArr(t.TagList) //把表 变成一维

	skillInts := make([]int, 5)
	skillInts = []int{1, 2, 3, 4, 5}

	for len(t.SkillTags) != skillNum {
		if count > 200 {
			global.GVA_LOG.Error("FillSkillTags 循环次数太多")
			break
		}
		count++
		arrsRand := helper.RandInt(len(arrs))
		tag := arrs[arrsRand]
		if _, ok := t.SkillTags[[2]int{tag.X, tag.Y}]; !ok {
			ir := helper.RandInt(len(skillInts))
			t.SkillTags[[2]int{tag.X, tag.Y}] = skillInts[ir]
			skillInts = append(skillInts[:ir], skillInts[ir:]...)
		}
		if len(t.SkillTags) == skillNum {
			break
		}
	}
	fmt.Println("FillSkillTags")
}

func (t *Table) LeveEvent() {
	//level 1 、level 2、level 3、level 4、level 5分别对应该等级
	//赢钱倍率*2、*4、*6、*8、*10，
	//每个等级所需升级标签个数为114、116、120、125、132。
	//每达到一个等级，进度条5个技能全部重置补充，同时该档位中赢钱乘以当前档位倍率，
	collect := []int{114, 116, 120, 125, 132}
	skillMul := []int{2, 4, 6, 8, 10}
	for i := 0; i < 5; i++ {
		le := &LevelEvent{
			Collect:  collect[i],
			SkillMul: skillMul[i],
		}
		t.LevelEvents = append(t.LevelEvents, le)
	}
}

func (t *Table) Skill1() {

}

func (t *Table) Skill2() {

}

func (t *Table) Skill3() {

}

func (t *Table) Skill4() {

}

func (t *Table) Skill5() {

}
