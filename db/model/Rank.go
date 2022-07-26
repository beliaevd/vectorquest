package model

import "Vectorquest/db"

type Rank struct {
	ID   int
	Name string
}

func (rank *Rank) GetFirstRank() *Rank {
	db.GetDB().Where("id = ?", 1).First(&rank)
	return rank
}
