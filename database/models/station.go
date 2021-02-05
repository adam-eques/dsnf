packages models 

type Station struct {
	ID uuid.UUID `json:"id"`
	StationName string `json:"station_name"`
	LocationId  Location `json:"location_id"`
}