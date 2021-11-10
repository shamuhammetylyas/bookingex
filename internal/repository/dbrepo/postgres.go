package dbrepo

import (
	"context"
	"time"

	"github.com/ShamuhammetYlyas/bookings/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	// Reservation insert ishlemi 3 sekuntdan kop wagt alsa onda insert ishlemi cancel bolar yaly context doredyaris
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int
	stmt := `Insert into reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
			values($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	//yokarda doreden kontextimizi Exec funksiyasy bilen bile ishletmek ucin Exec dalde ExecContext yazyarys.
	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

func (m *postgresDBRepo) InsertRoomRestriction(res models.RoomRestriction) error {
	// InsertRoomRestriction insert ishlemi 3 sekuntdan kop wagt alsa onda insert ishlemi cancel bolar yaly context doredyaris
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `Insert into room_restrictions(start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at)
			values($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, stmt,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		res.ReservationID,
		res.RestrictionID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}
