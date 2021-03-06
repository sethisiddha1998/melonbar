package main

import (
	"encoding/hex"
	"strings"

	"github.com/BurntSushi/xgbutil/xgraphics"
)

func hexToBGRA(h string) xgraphics.BGRA {
	h = strings.Replace(h, "#", "", 1)
	d, _ := hex.DecodeString(h)

	return xgraphics.BGRA{B: d[2], G: d[1], R: d[0], A: 0xFF}
}

// TODO: Instead of doing this using rune-count, do this using pixel-count.
func trim(txt string, l int) string {
	if len(txt) > l {
		return txt[0:l] + "..."
	}
	return txt
}
