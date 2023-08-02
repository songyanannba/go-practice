package slotHandle

import (
	"errors"
	"go.uber.org/zap"
	"math/rand"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/service/slot"
	"slot-server/service/slot/base"
	"slot-server/service/slot/component"
	"slot-server/utils/helper"
	"time"
)

type MergeSpin struct {
	UserId uint
	First  *component.Spin
	Spin   *component.Spin // 当前spin主体（非freeSpin）

	UserSpin *business.SlotUserSpin // 用户spin信息

	FinishNum         uint // 完成次数
	FreeSpinFinishNum uint // 免费完成次数
	ReSpinFinishNum   uint // respin完成次数
	Spins             []*component.Spin

	CurrentID int // 当前id

	FreeDebug bool // 是否需要免费转debug
	ReDebug   bool // 是否需要respin debug

	FreeRank int // 免费进度
	ReRank   int // respin进度
}

func (m *MergeSpin) newSpinByParent(p *component.Spin, isFreeSpin bool) *component.Spin {
	s := component.Spin{
		Id: m.CurrentID,
		Options: &component.Options{
			IsFree:   isFreeSpin,
			IsReSpin: !isFreeSpin,
			Demo:     p.Demo,
			IsTest:   p.IsTest,
			Raise:    p.Raise,
		},
		Config:   p.Config,
		ParentId: p.Id,
		PSpin:    p,
		Bet:      p.Bet,
	}

	m.CurrentID++

	if isFreeSpin {
		if m.FreeDebug {
			s.SetDebugInitData(m.UserId)
			m.FreeDebug = false
		}
		s.PlayNum = m.UserSpin.FreeNum + m.FreeSpinFinishNum + 1
		s.Rank = m.FreeRank
	} else {
		if m.ReDebug {
			s.SetDebugInitData(m.UserId)
			m.ReDebug = false
		}
		// 算上普通转
		s.PlayNum = m.UserSpin.FreeNum + m.FreeSpinFinishNum + 2
		s.Rank = m.ReRank
	}
	return &s
}

func (m *MergeSpin) runFree(p *component.Spin) error {
	for i := 0; i < p.FreeSpinParams.FreeNum; i++ {
		spin := m.newSpinByParent(p, true)
		_, _ = slot.RunSpin(spin)
		m.Spins = append(m.Spins, spin)

		m.FreeRank = spin.NextRank
		m.FreeSpinFinishNum++
		m.FinishNum++
		if err := m.isExceed(); err != nil {
			return err
		}
		// 递归
		err := m.runFree(spin)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MergeSpin) runRe(p *component.Spin) error {
	for i := 0; i < p.FreeSpinParams.ReNum; i++ {
		spin := m.newSpinByParent(p, false)
		_, _ = slot.RunSpin(spin)
		m.Spins = append(m.Spins, spin)

		m.ReRank = spin.NextRank
		m.ReSpinFinishNum++
		m.FinishNum++
		if err := m.isExceed(); err != nil {
			return err
		}
		// 递归
		err := m.Run(spin)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MergeSpin) Run(p *component.Spin) error {
	if p.FreeSpinParams.FreeNum > 0 {
		err := m.runFree(p)
		if err != nil {
			return err
		}
		// 该轮freeSpin结束 重置进度 reRank由spin内部控制
		m.FreeRank = 0
	}

	if p.FreeSpinParams.ReNum > 0 {
		err := m.runRe(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func RunMergeSpin(userId uint, spin *component.Spin, userSpin *business.SlotUserSpin) (mer *MergeSpin, err error) {
	m := &MergeSpin{
		UserId:    userId,
		First:     spin,
		Spin:      spin,
		UserSpin:  userSpin,
		CurrentID: 2,
		FreeDebug: true,
		ReDebug:   true,
		ReRank:    spin.NextRank,
	}

	switch spin.Config.SlotId {
	case 4:
		err = m.Unit4AllSpin(spin)
	default:
		err = m.Run(spin)
	}
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MergeSpin) isExceed() error {
	if m.FinishNum >= 10000 {
		if !m.First.IsTest {
			global.GVA_DB.Create(&business.SlotRecord{
				Date:   time.Now().Format("20060102"),
				UserId: m.UserId,
				Bet:    m.First.Bet,
				Raise:  m.First.Raise,
				SlotId: m.First.Config.SlotId,
				IsBk:   enum.No,
				//PlayType: enum.SlotSpinType2Fs,
				Status: enum.CommonStatusError,
			})
		}
		global.GVA_LOG.Error("reSpin迭代次数超过100")
		return errors.New("more than 100 reSpin iterations")
	}
	return nil
}

type count struct {
	index   int
	freeNum int
	reNum   int
}

func (m *MergeSpin) Unit4AllSpin(p *component.Spin) error {
	defer func() {
		if err := recover(); err != nil {
			global.GVA_LOG.Info("err", zap.Any("err", err))
			return
		}
	}()
	count := count{
		index: 1,
	}
	err := m.Unit4ReSpin(p, &count)
	if err != nil {
		return err
	}
	return m.Unit4FreeSpin(p, &count)
}

func (m *MergeSpin) Unit4ReSpin(p *component.Spin, count *count) error {
	var lockTags = GetLockTags(p.InitDataList)
	if m.Spin.FreeSpinParams.ReNum == 0 || len(lockTags) == 15 {
		return nil
	}
	superiorId := count.index
	Xs, num, idWinOdds := m.GetOdds(lockTags)

	var opts = []component.Option{component.WithReSpin()}
	if m.First.IsTest {
		opts = append(opts, component.WithTest())
		opts = append(opts)
	}

	for i := 0; i < num; i++ {
		count.index++

		lockTags, Xs = m.SetOdds(lockTags, idWinOdds, i, Xs)
		newOpts := append(opts,
			component.WithPlayNum(m.UserSpin.FreeNum+m.FreeSpinFinishNum+1),
			component.WithTest(),
			component.WithTagsLock(lockTags),
			component.SetFreeNum(m.Spin.FreeNum),
			component.SetResNum(m.Spin.ResNum),
		)

		if i == 0 && count.reNum == 0 {
			newOpts = append(newOpts, component.WithNeedSpecify(true))
			count.reNum++
		}

		machine, err := slot.Play(m.Spin.Config.SlotId, m.Spin.Bet,
			newOpts...,
		)
		if err != nil {
			return err
		}
		spin := machine.GetSpin()
		num += spin.FreeSpinParams.ReNum
		m.Spin = spin
		spin.Id = count.index
		spin.ParentId = superiorId
		spin.PSpin = p
		m.Spins = append(m.Spins, spin)
		m.ReSpinFinishNum++
		m.FinishNum++
		if err = m.isExceed(); err != nil {
			return err
		}
		if m.Spin.Jackpot != nil {
			return nil
		}
		lockTags = base.GetSpecialTags(spin.InitDataList, "link_coin", "link_collect")
		Xs = GetNeedFill(lockTags)
		if len(lockTags) == 15 {
			break
		}
	}
	return nil
}

func (m *MergeSpin) Unit4FreeSpin(p *component.Spin, count *count) error {
	fNum := p.FreeSpinParams.FreeNum
	var opts = []component.Option{component.WithFreeSpin()}
	if m.First.IsTest {
		opts = append(opts, component.WithTest())
	}
	for i := 0; i < fNum; i++ {
		count.index++

		newOpts := append(opts,
			component.WithPlayNum(m.UserSpin.FreeNum+m.FreeSpinFinishNum+1),
			component.SetFreeNum(m.Spin.FreeNum),
			component.SetResNum(m.Spin.ResNum),
		)
		if i == 0 && count.freeNum == 0 {
			newOpts = append(newOpts, component.WithNeedSpecify(true))
			count.freeNum++
		}
		machine, err := slot.Play(m.Spin.Config.SlotId, m.Spin.Bet,
			newOpts...,
		)
		if err != nil {
			return nil
		}
		spin := machine.GetSpin()
		spin.Id = count.index
		spin.ParentId = p.Id
		spin.PSpin = p
		// 增加免费次数

		// 完成次数增加
		m.FinishNum++
		m.FreeSpinFinishNum++
		m.Spins = append(m.Spins, spin)
		m.Spin = spin
		if err = m.isExceed(); err != nil {
			return err
		}
		if spin.FreeSpinParams.ReNum > 0 {
			err := m.Unit4ReSpin(spin, count)
			if err != nil {
				return err
			}
		}

		if spin.FreeSpinParams.FreeNum > 0 {
			err := m.Unit4FreeSpin(spin, count)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetLockTags(tagList [][]*base.Tag) []*base.Tag {
	var lockTags = base.GetSpecialTags(tagList, "link_coin", "link_collect")
	Tags := make([]*base.Tag, 0)
	for _, plusSpin := range lockTags {
		var fp = *plusSpin
		fp.X = plusSpin.Y*5 + plusSpin.X
		fp.Y = 0
		Tags = append(Tags, &fp)
	}

	return Tags
}

type WinOdds struct {
	id   int
	link int
	coll int
	plus int
}

// SetOdds 获取标签数量
func (m *MergeSpin) SetOdds(lockTags []*base.Tag, idWinOdds map[int]*WinOdds, i int, Xs []int) ([]*base.Tag, []int) {
	linkNum := idWinOdds[i].link
	CollNum := idWinOdds[i].coll
	PlusNum := idWinOdds[i].plus
	for _, tag := range lockTags {
		tag.ISLock = true
	}
	for j := 0; j < linkNum && len(Xs) > 0; j++ {
		var tag = *m.First.Config.GetTag("link_coin")
		tagxId := rand.Intn(len(Xs))
		tag.X = Xs[tagxId]
		for i, x := range Xs {
			if x == tag.X {
				Xs = append(Xs[:i], Xs[i+1:]...)
				break
			}
		}
		tag.Y = 0
		tag.ISLock = false
		lockTags = append(lockTags, &tag)
	}

	for j := 0; j < CollNum && len(Xs) > 0; j++ {
		var tag = *m.First.Config.GetTag("link_collect")
		tagxId := rand.Intn(len(Xs))
		tag.X = Xs[tagxId]
		for i, x := range Xs {
			if x == tag.X {
				Xs = append(Xs[:i], Xs[i+1:]...)
			}
		}
		tag.Y = 0
		tag.ISLock = false
		lockTags = append(lockTags, &tag)
	}

	for j := 0; j < PlusNum && len(Xs) > 0; j++ {
		var tag = *m.Spin.Config.GetTag("plus_spin")
		tagxId := rand.Intn(len(Xs))
		tag.X = Xs[tagxId]
		for i, x := range Xs {
			if x == tag.X {
				Xs = append(Xs[:i], Xs[i+1:]...)
			}
		}
		tag.Y = 0
		tag.ISLock = false
		lockTags = append(lockTags, &tag)
	}
	return lockTags, Xs
}

// GetOdds 设置标签
func (m *MergeSpin) GetOdds(lockTags []*base.Tag) ([]int, int, map[int]*WinOdds) {
	weight := (len(lockTags)-6)*3 + 2

	defer func() {
		if err := recover(); err != nil {
			global.GVA_LOG.Info("weight", zap.Any("weight", weight))
			global.GVA_LOG.Info("err", zap.Any("err", err))
			return
		}
	}()
	Xs := GetNeedFill(lockTags)

	sumLinkNum := m.First.Config.Event.M[weight].(*base.ChangeTableEvent).Fetch() //本局再出现的link数量
	collNum := m.First.Config.Event.M[weight+1].(*base.ChangeTableEvent).Fetch()  //本局再出现的collect数量
	plusNum := m.First.Config.Event.M[weight+2].(*base.ChangeTableEvent).Fetch()  //本局再出现的plus数量
	LNum := helper.If(sumLinkNum-collNum > 0, sumLinkNum-collNum, 0)              //本局再出现的link数量
	collNum = helper.If(sumLinkNum-LNum > 0, sumLinkNum-LNum, 0)                  //本局再出现的collect数量
	num := m.Spin.FreeSpinParams.ReNum
	idlist := make([]int, 0)
	idWinOdds := map[int]*WinOdds{}
	for i := 0; i < num+plusNum; i++ {
		idlist = append(idlist, i)
		idWinOdds[i] = &WinOdds{}
	}
	plusNumv := 0
	collNumv := 0
	LNumv := 0
	for plusNumv < plusNum {
		plusRand := helper.RandInt(num + plusNumv)
		if idWinOdds[plusRand].plus < 2 {
			idWinOdds[plusRand].plus++
			plusNumv++
			continue
		}
	}
	for collNumv < collNum {
		collRand := helper.RandInt(num + plusNum)
		if idWinOdds[collRand].coll < 2 {
			idWinOdds[collRand].coll++
			collNumv++
			continue
		}
	}
	for LNumv < LNum {
		LRand := helper.RandInt(num + plusNum)
		if idWinOdds[LRand].link < 2 {
			idWinOdds[LRand].link++
			LNumv++
			continue
		}
	}
	return Xs, num, idWinOdds
}

// GetNeedFill 获取可填充位置
func GetNeedFill(lockTags []*base.Tag) []int {
	Xs := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	for _, tag := range lockTags {
		for i, x := range Xs {
			if x == tag.X {
				Xs = append(Xs[:i], Xs[i+1:]...)
			}
		}
	}
	return Xs
}
