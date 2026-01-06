module github.com/Snansidansi/hue-api-go

go 1.23.2

retract (
	v1.0.1 // Contains faulty import paths
	v1.0.2 // Checksum error or faulty indexing
)

require github.com/joho/godotenv v1.5.1
