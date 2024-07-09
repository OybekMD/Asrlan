package models

import "asrlan-monolight/storage/repo"

type Profile struct {
	User          repo.ProfileUser            `json:"user"`
	StatisticYear map[string][]repo.Statistic `json:"statistic"`
	StatisticWMY  []*repo.Statistic           `json:"statisticwmy"`
	Badge         []*repo.Badge          `json:"badge"`
	Certificate   []*repo.Certificate         `json:"certificate"`
}