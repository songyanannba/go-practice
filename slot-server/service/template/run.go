package template

import (
	"fmt"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/service/slot/base"
	"slot-server/service/slot/template"
	"slot-server/service/slot/template/flow"
	"slot-server/utils/helper"
)

func NewTestGameInfo(gTem *GenTemplate) *template.SpinInfo {
	tagMapById := make(map[int]*base.Tag)
	tagMapByName := make(map[string]*base.Tag)
	tags := gTem.Config.GetAllTagQuote()
	for i, tag := range tags {
		tagMapById[i] = tag
		tagMapByName[tag.Name] = tag
	}

	config := &template.Config{
		SlotId:       gTem.SlotId,
		Event:        gTem.Config.Event,
		PayTableList: gTem.Config.PayTableList,
		Template:     gTem.Template, //列 和 列标签
		Col:          gTem.Config.Index,
		Row:          gTem.Config.Row,
		GameType:     int(gTem.Type),
		TagMapById:   tagMapById,
		TagMapByName: tagMapByName,
	}
	info := template.NewGameInfo(config)
	return info
}

func TemTest(gTem *GenTemplate) (error, *RunResult) {
	switch gTem.Type {
	case enum.FreeSpin:
		return TemExecFree(gTem)
	case enum.NormalSpin:
		return TemExec(gTem)
	default:
		return fmt.Errorf("not support type %d", gTem.Type), nil
	}
}

func TemExec(gTem *GenTemplate) (error, *RunResult) {

	sumSca := 0
	si := NewTestGameInfo(gTem)
	err := TemRun(si, 0)
	if err != nil {
		return err, nil
	}

	scatters := si.FindTagsByName(enum.ScatterName)
	scatterLine := si.GetWinLine(scatters)
	si.Scatter = scatterLine

	multipliers := si.FindTagsByName(enum.MultiplierName)
	si.Multiplier = multipliers

	if len(scatters) >= 4 {
		sumSca++
	}
	gain := si.GetWin(100)
	return nil, &RunResult{
		Gain:           gain,
		ScatterTrigger: sumSca,
		SpinCount:      1,
	}
}

func TemExecFree(gTem *GenTemplate) (error, *RunResult) {
	sumGain := 0
	sumSca := 0
	spinCount := 15
	var SumMultipliers []*base.Tag
	for i := 0; i < spinCount; i++ {
		si := NewTestGameInfo(gTem)
		err := TemRun(si, 0)
		if err != nil {
			return err, nil
		}

		scatters := si.FindTagsByName(enum.ScatterName)
		scatterLine := si.GetWinLine(scatters)
		si.Scatter = scatterLine

		multipliers := si.FindTagsByName(enum.MultiplierName)
		SumMultipliers = append(SumMultipliers, multipliers...)
		si.Multiplier = helper.CopyList(SumMultipliers)

		if len(scatters) >= 3 {
			sumSca++
			spinCount += 5
		}
		sumGain += si.GetWin(100)
	}

	return nil, &RunResult{
		Gain:           sumGain,
		ScatterTrigger: sumSca,
		SpinCount:      spinCount,
	}
}

func TemRun(si *template.SpinInfo, count int) error {
	SpinFlow := flow.NewSpinFlow(count)
	SpinFlow.FlowMap += si.PrintTable("初始")
	SpinFlow.InitList = helper.CopyListArr(si.Display)
	//获取划线
	lines := si.FindCountLine(enum.LineLength)

	//获取划线赢钱
	SpinFlow.AddOmitList(si.GetWinLines(lines)...)
	//删除标签
	si.DeleteTagList(lines)
	SpinFlow.FlowMap += si.PrintTable("删除")
	//掉落标签
	si.Drop()
	SpinFlow.FlowMap += si.PrintTable("掉落")
	if len(si.GetEmptyTags()) > 0 {
		global.GVA_LOG.Error("模版缺失还有标签没有填充" + fmt.Sprintf("流程长度:%d\n", len(si.SpinFlow)))
		//return fmt.Errorf("模版缺失还有标签没有填充" + fmt.Sprintf("流程长度:%d\n", len(si.SpinFlow)))
	}
	si.SpinFlow = append(si.SpinFlow, SpinFlow)
	lines = si.FindCountLine(enum.LineLength)
	if len(lines) == 0 {
		return nil
	}
	if count > enum.SlotMaxSpinNum {
		//global.GVA_LOG.Error(enum.SlotMaxSpinStr)
		return fmt.Errorf(enum.SlotMaxSpinStr)
	}
	return TemRun(si, count+1)
}
