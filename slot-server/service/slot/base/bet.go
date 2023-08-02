package base

import (
	"regexp"
	"slot-server/utils/helper"
	"strings"
)

type Bet struct {
	Currency string
	Bets     []int64
}

func NewBet(currency string, bets []int64) *Bet {
	return &Bet{
		Currency: currency,
		Bets:     bets,
	}
}

type BetMap struct {
	m map[string]*Bet
}

func NewBetMap(s string) BetMap {
	m := BetMap{m: make(map[string]*Bet)}
	if s == "" {
		return m
	}

	reg := regexp.MustCompile("\\s+")
	s = reg.ReplaceAllString(s, "")
	if s == "" {
		return m
	}

	r := regexp.MustCompile(`([A-Z]+):([0-9,]+)`)
	matches := r.FindAllStringSubmatch(s, -1)

	for _, match := range matches {
		if len(match) != 3 {
			continue
		}

		currency := match[1]
		bets := match[2]

		currency = strings.ToUpper(currency)
		m.m[currency] = NewBet(currency, helper.SplitInt[int64](bets, ","))
	}
	return m
}

func (b BetMap) Get(currency string) *Bet {
	currency = strings.ToUpper(currency)
	if bet, ok := b.m[currency]; ok {
		return bet
	}
	if bet, ok := b.m["USD"]; ok {
		return bet
	}
	return NewBet("", []int64{0})
}

func (b BetMap) Check(currency string, bet int64) bool {
	if bet == 0 {
		return false
	}
	return helper.InArr(bet, b.Get(currency).Bets)
}
