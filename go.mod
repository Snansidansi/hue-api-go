module github.com/Snansidansi/hue-api-go

go 1.23.2

retract (
	v1.0.1 // Contains faulty import paths
	v1.0.2 // Checksum error or faulty indexing
	v1.0.0 // Changed visibility of Structs
	v1.0.3 // -||-
	v1.0.4 // Fixed invalid names of service methodes without: e165d13
	v1.0.5 // mod.go not updated with tag retracts after v1.0.4 deletion
	v1.0.6 // Eventstream events are pointers to the events
)

require github.com/joho/godotenv v1.5.1
