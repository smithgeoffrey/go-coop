package api

type Door struct {
	UpSensor  bool `json:"upsensor"`
	DownSensor  bool `json:"downsensor"`
	Status string `json:"status"`
}

type Temp struct {
	InsideSensor float32 `json:"inside"`
	OutsideSensor float32 `json:"outside"`
}

type Video struct {
	Location string `json:"location"`
	Url  string `json:"url"`
}
