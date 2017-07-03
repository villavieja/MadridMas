package incident

type Incident struct {
	Id          int
	Latitude    float64
	Longitude   float64
	Creation    string
	Description string
	Status      int
}
