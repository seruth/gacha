package gacha

import "errors"

type Player struct {
	tickets int // ガチャ券の枚数
	coin    int // コイン
}

// 引数にガチャ券とコインの枚数をもらい、
// それぞれのフィールドに設定したPlayer型の値を生成し、
// そのポインタを返す NewPlayer 関数
func NewPlayer(tickets int, coin int) *Player {
	return &Player{tickets: tickets, coin: coin}
}

// プレイヤーが行えるガチャの回数
func (p *Player) DrawableNum() int {
	// ガチャが行える回数を返す
	// ガチャ券は1枚で1回、コインは10枚で1回ガチャが行える
	return p.tickets + p.coin/10
}

func (p *Player) draw(n int) error {

	if p.DrawableNum() < n {
		return errors.New("ガチャ券またはコインが不足しています")
	}

	// ガチャ券で足りる場合はガチャ券だけ使う
	// ガチャ券から優先的に使う
	if p.tickets > n {
		p.tickets -= n
		return nil
	}

	p.coin -= (n - p.tickets) * 10 // 一回あたり10 枚消費する
	p.tickets = 0

	return nil
}
