package config

const (
	VIDEO_URL = "http://172.16.1.128/video.mjpg"
	TEMP_URL = "http://localhost:8081/api/v1/sensor/temp"
	DOOR_URL = "http://localhost:8081/api/v1/sensor/door"

	MOCK_TEMP_INSIDE_SENSOR = "35.1"
	MOCK_TEMP_OUTSIDE_SENSOR = "28.4"

	MOCK_DOOR_DOWN_SENSOR = false
	MOCK_DOOR_UP_SENSOR = true
)