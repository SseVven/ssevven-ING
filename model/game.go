package model

var (
	GameOne = &Game{
		GameName: "数字华容道",
		// GameURL:  Cfg.Domain + "/game/game1",
		GameURL: "http://ssevven.free.idcfengye.com" + "/game/game1",
	}
	GameTwo = &Game{
		GameName: "数独",
		// GameURL:  Cfg.Domain + "/game/game1",
		GameURL: "http://ssevven.free.idcfengye.com" + "/game/game2",
	}
)

type Game struct {
	GameName string
	GameURL  string
}
