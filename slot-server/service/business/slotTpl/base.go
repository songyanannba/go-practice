package slotTpl

import (
	"errors"
	"golang.org/x/exp/maps"
	"slot-server/model/business"
	"slot-server/service/slot/component"
	"slot-server/utils/helper"
	"strings"
)

type TplHandle struct {
	Model  *business.SlotGenTpl
	Params *TplParams
	Config *component.Config
}

type TplParams struct {
	Height    int
	Width     int
	TagNum    int   `json:"tag_num"`
	TagWeight []int `json:"tag_weight"`
}

func NewTplHandle(model *business.SlotGenTpl) *TplHandle {
	return &TplHandle{Model: model}
}

func (h *TplHandle) ParseParams() (err error) {
	if h.Model.SlotId < 5 {
		return errors.New("slot " + helper.Itoa(h.Model.SlotId) + " not a puzzle game")
	}
	if h.Model.Num < 1 {
		return errors.New("num must be greater than 0")
	}
	h.Model.Params, h.Params, err = FormatTplParams(h.Model.Params)
	if h.Model.Size != "" {
		size := helper.SplitInt[int](h.Model.Size, "*")
		if len(size) != 2 {
			err = errors.New("size format error")
			return
		}
		h.Params.Width = size[0]
		h.Params.Height = size[1]
	} else {
		h.Params.Width = 7
		h.Params.Height = 7
	}
	return
}

func ParseTplParams(m map[string]string, tplParams *TplParams) (err error) {
	must := []string{"tag_num", "tag_num"}
	diff := helper.ArrDifference(must, maps.Keys(m))
	if len(diff) > 0 {
		return errors.New("params " + strings.Join(diff, ",") + " is required")
	}

	if tplParams.TagNum = helper.Atoi(m["tag_num"]); tplParams.TagNum <= 0 {
		return errors.New("tag_num must be greater than 0")
	}

	if tplParams.TagWeight = helper.SplitInt[int](m["tag_weight"], ","); len(tplParams.TagWeight) != tplParams.TagNum+1 {
		return errors.New("tag_weight length must be equal to tag_num + 1")
	}
	return
}

func FormatTplParams(s string) (ss string, tplParams *TplParams, err error) {
	tplParams = &TplParams{}
	s = strings.Trim(s, "\n\r ")
	params := strings.Split(s, "\n")
	paramsMap := make(map[string]string)
	for _, param := range params {
		name, value, ok := strings.Cut(param, ":")
		name = strings.Trim(name, " ")
		value = strings.Trim(value, " ")
		if !ok || value == "" || name == "" {
			err = errors.New("param [" + name + "] format error")
			return
		}
		paramsMap[name] = value
		ss += name + ":" + value + "\n"
	}
	err = ParseTplParams(paramsMap, tplParams)
	return
}
