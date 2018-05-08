package config

type config struct {
	Verbose     bool `json:"verbose_logger"`
	BurgerStore struct {
		Staff struct {
			TillWorkers int `json:"till_workers"`
			Milkshakes  int `json:"milkshake_maker"`
			Fryer       int `json:"fryer"`
			Grill       int `json:"grill"`
			Soda        int `json:"soda"`
		} `json:"staff"`
		Equipment struct {
			MilkshakeMixer int `json:"milkshake_mixer"`
			Fryers         int `json:"fryers"`
			Grill          int `json:"grills"`
			Soda           int `json:"soda"`
		} `json:"equipment"`
	} `json:"burger_store"`
}
