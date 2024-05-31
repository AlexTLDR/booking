package db

const (
	DBNAME     = "booking"
	TestDBNAME = "bookig-test"
	DBURI      = "mongodb://localhost:27017"
)

type Store struct {
	User  UserStore
	Hotel HotelStore
	Room  RoomStore
}
