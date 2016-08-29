package player

import "time"

type Player struct {
	ID                 int       `db:"id"`
	FirstName          string    `db:"firstName"`
	LastName           string    `db:"lastName"`
	BirthDate          time.Time `db:"birthDate"`
	Email              string    `db:"email"`
	Password           string    `db:"password"`
	Roles              string    `db:"roles"`
	Enabled            bool      `db:"enabled"`
	Created            time.Time `db:"birthDate"`
	Rating             float32   `db:"rating"`
	RatingMean         float32   `db:"ratingMean"`
	RatingStdDeviation float32   `db:"ratingStdDeviation"`
}
