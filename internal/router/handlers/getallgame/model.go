package getallgame

type Request struct {
}

type Response struct {
	GameIds []string `json:"game_ids"`
}
