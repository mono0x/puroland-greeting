package normalizing

import "golang.org/x/text/width"

var placeRules = map[string]string{
	"ビレッジ(1F)": "ピューロビレッジ(1F)",
	"エンターティメントホール入口(1F)":       "エンターテイメントホール入口(1F)",
	"エンターティメントホール(1F)":         "エンターテイメントホール(1F)",
	"3Fきゃらぐりスポット（レインボーホール 3F）": "3F(3F)",
	"キャンディファクトリー":              "キャンディファクトリー(1F)",
	"メイメロードドライブ付近":             "マイメロードドライブ横(1F)",
	"ポムポムプリンフォトスポット":           "PNフォト(4F)",
	"マイメロディショップきゃらぐり":          "4Fマイメロディショップ(4F)",
	"縁日会場（きゃらグリ）":              "キャラクター縁日会場(1F)",
}

func NormalizePlace(source string) string {
	folded := width.Fold.String(source)
	if replaced, ok := placeRules[folded]; ok {
		return replaced
	}
	return folded
}
