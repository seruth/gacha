package gacha

type Rarity string

const (
	RarityN  Rarity = "N"
	RarityR  Rarity = "R"
	RaritySR Rarity = "SR"
	RarityXR Rarity = "XR"
)

func (r Rarity) String() string {
	return string(r)
}

type Card struct {
	Rarity Rarity // レア度
	Name   string // 名前
}

func (c *Card) String() string {
	// レア度:名前のように文字列を作る
	// 例："SR:ドラゴン"
	return c.Rarity.String() + ":" + c.Name
}
