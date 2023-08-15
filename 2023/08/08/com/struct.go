package com

type Spin struct {
	Id int
}

type Machine interface {
	Exec() error
	GetInitData()
	GetResData()
}
