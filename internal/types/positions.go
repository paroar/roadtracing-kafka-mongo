package types

// Position type
type Position struct {
	Timestamp        int     `json:"timestamp"`
	Altitude         float32 `json:"altitude"`
	AltitudeAccuracy float32 `json:"altitudeAccuracy"`
	Heading          float32 `json:"heading"`
	Latitude         float32 `json:"latitude"`
	Longitude        float32 `json:"longitude"`
	Accuracy         float32 `json:"accuracy"`
	Speed            float32 `json:"speed"`
	X                float32 `json:"z"`
	Y                float32 `json:"y"`
	Z                float32 `json:"x"`
	Deviceid         string  `json:"deviceid"`
	Groupid          string  `json:"groupid"`
}
