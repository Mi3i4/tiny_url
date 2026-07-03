package shortener

import "time"

type (
	URL  string
	Code string
)

type ShortLink struct {
	Code      Code      // short code
	Original  URL       // original URL
	CreatedAt time.Time // creation time
	Clicks    int64     // click counter (int64 — will grow)
}

type Stats struct {
	Code        Code
	TotalClicks int64
	LastClick   time.Time
}
