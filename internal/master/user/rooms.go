package user

type Rooms struct {
}

func NewRooms() *Rooms {
	return &Rooms{}
}

func (rooms *Rooms) CreateRoom(name string) *Room {
	return NewRoom("123", name)
}
