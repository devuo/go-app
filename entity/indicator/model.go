package indicator

import "time"

type Indicator struct {
	ID        int       `db:"id"`
	Namespace string    `db:"namespace"`
	Type      string    `db:"type"`
	Variant   string    `db:"variant"`
	Value     float32   `db:"value"`
	Date      time.Time `db:"date"`
}
