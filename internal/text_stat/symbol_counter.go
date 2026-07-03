package textstat

import (
	"github.com/rivo/uniseg"
	"unicode/utf8"
)

// Count возвращает три метрики длины строки s, которые в Unicode не совпадают:
//
//   - bytes     — число байт в UTF-8 (len). ASCII = 1 байт, кириллица = 2,
//     эмодзи = 4 и более.
//   - runes     — число рун (Unicode code points). Один видимый символ может
//     состоять из нескольких рун.
//   - graphemes — число видимых символов (кластеров графем, UAX #29): столько
//     «символов» видит человек.
//
// Например, семейное эмодзи 👨‍👩‍👧‍👦 — это 25 байт, 7 рун, но 1 графема.
//
// Стандартная библиотека считает байты и руны, но не графемы — для них
// используется github.com/rivo/uniseg.
func Count(s string) (bytes, runes, graphemes int) {
	bytes = len(s)
	runes = utf8.RuneCountInString(s)
	graphemes = uniseg.GraphemeClusterCount(s)

	return bytes, runes, graphemes
}
