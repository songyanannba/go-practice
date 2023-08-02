package template

import (
	"fmt"
	"go.uber.org/zap"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/utils/helper"
)

func CreateTemplate(tem *business.SlotTemplateGen) error {
	template, err := NewGenTemplate(tem)
	if err != nil {
		return err
	}
	var (
		str   string
		count int
		ratio *LogicResult
	)
	tem.Schedule = ""
	tem.State = enum.CommonStatusBegin
	global.GVA_DB.Save(tem)
	for {
		count++
		ratio = GoRun(template)
		i := template.Range(ratio)
		switch i {
		case MaxRange:
			err := template.AdjLarge()
			if err != nil {
				return err
			}
			str = fmt.Sprintf("第%d次:当前返还比:%g,当前Sca触发率:%g 进行大调整\n", count, ratio.GainRatio, ratio.ScatterRatio)
		case DownTrim:
			err := template.AdjTrimDown()
			if err != nil {
				return err
			}
			str = fmt.Sprintf("第%d次:当前返还比:%g,当前Sca触发率:%g 向下微调\n", count, ratio.GainRatio, ratio.ScatterRatio)
		case UpTrim:
			err := template.AdjTrimUp()
			if err != nil {
				return err
			}
			str = fmt.Sprintf("第%d次:当前返还比:%g,当前Sca触发率:%g 向上微调\n", count, ratio.GainRatio, ratio.ScatterRatio)
		case UpScatter:
			err := template.AdjScatterUp()
			if err != nil {
				return err
			}
			str = fmt.Sprintf("第%d次:当前返还比:%g,当前Sca触发率:%g 增加Scatter\n", count, ratio.GainRatio, ratio.ScatterRatio)
		case DownScatter:
			err := template.AdjScatterDown()
			if err != nil {
				return err
			}
			str = fmt.Sprintf("第%d次:当前返还比:%g,当前Sca触发率:%g 减少Scatter\n", count, ratio.GainRatio, ratio.ScatterRatio)
		case Ok:
			goto end
		}
		tem.Schedule += str
		global.GVA_DB.Save(tem)
		if count > 1 {
			return nil
		}
	}
end:
	tem.State = enum.CommonStatusFinish
	str = fmt.Sprintf("第%d次:当前返还比:%g,完成\n", count, ratio)
	tem.Schedule += str

	temStr := ""
	for i, tags := range template.Template {
		temStr += fmt.Sprintf("%d:", i)
		for _, tag := range tags {
			temStr += fmt.Sprintf("%s,", tag.Name)
		}
		temStr += "\n"
	}

	tem.Template = temStr
	finalWeight := ""
	for i, weight := range template.InitialWeight {
		sum := 0
		weiStr := fmt.Sprintf("%d:scatter&multiplier&high_1&high_2&high_3&high_4&low_1&low_2&low_3&low_4&low_5@0", i)
		sum += weight["scatter"]
		weiStr += fmt.Sprintf("&%d", sum)
		sum += weight["multiplier"]
		weiStr += fmt.Sprintf("&%d", sum)
		sum += weight["high_1"]
		weiStr += fmt.Sprintf("&%d", sum)
		sum += weight["high_2"]
		weiStr += fmt.Sprintf("&%d", sum)
		sum += weight["high_3"]
		weiStr += fmt.Sprintf("&%d", sum)
		sum += weight["high_4"]
		weiStr += fmt.Sprintf("&%d", sum)
		sum += weight["low_1"]
		weiStr += fmt.Sprintf("&%d", sum)
		sum += weight["low_2"]
		weiStr += fmt.Sprintf("&%d", sum)
		sum += weight["low_3"]
		weiStr += fmt.Sprintf("&%d", sum)
		sum += weight["low_4"]
		weiStr += fmt.Sprintf("&%d", sum)
		sum += weight["low_5"]
		weiStr += fmt.Sprintf("&%d", sum)
		finalWeight += weiStr + "\n"
	}
	tem.FinalWeight = finalWeight
	global.GVA_DB.Save(tem)
	template.CreateTem(tem)
	return nil
}

func GoRun(tem *GenTemplate) *LogicResult {
	var runResults []*RunResult

	ch, errCh, _ := helper.Parallel(10, 1000, func() (result *RunResult, err error) {
		err, result = TemTest(tem)
		return
	})

	count := 0
	a := 0
	for {
		a++
		select {
		case err := <-errCh:
			global.GVA_LOG.Error("GoRun", zap.Error(err))
			//runResults = append(runResults, 0)
		case v, beforeClosed := <-ch:
			if !beforeClosed {
				goto end
			}
			runResults = append(runResults, v)
			count++
		}
	}
end:
	return GetLogicResult(runResults, int(tem.Type))
}
