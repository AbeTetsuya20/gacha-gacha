package gacha

import (
	"math/rand"
	"time"
)

func init() {
	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())
}

func DrawN(p *Player, n int) ([]*Card, map[Rarity]int) {
	p.draw(n)

	results := make([]*Card, n)
	summary := make(map[Rarity]int)
	for i := 0; i < n; i++ {
		results[i] = draw()
		summary[results[i].Rarity]++
	}

	// 変数resultsとsummaryの値を戻り値として返す
	return results, summary
}

func draw() *Card {
	num := rand.Intn(100)

	normal := [...]{"スライム1","スライム2"}
	rare := [...]{"オーク1","オーク2"}
	super := [...]{"ドラゴン1","ドラゴン2"}
	xrare := [...]{"イフリート1","イフリート2"}

	//numでレア度が決まる→choiceでインデックスが決まる→そのインデックスに対応したキャラ名がreturnされる
	switch {
	case num < 80:
		choice := rand.Intn(len(normal))
		return &Card{Rarity: RarityN, Name:normal[choice]}
	case num < 95:
		choice := rand.Intn(len(rare))
		return &Card{Rarity: RarityR, Name:rare[choice]}
	case num < 99:
		choice := rand.Intn(len(Normal))
		return &Card{Rarity: RaritySR, Name: super[choice]}
	default:
		choice := rand.Intn(len(Normal))
		return &Card{Rarity: RarityXR, Name: xrare[choice]}
	}
}

