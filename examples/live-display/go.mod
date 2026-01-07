module github.com/Snasidansi/hue-api-go/examples/live-display

go 1.24.0

toolchain go1.24.11

replace github.com/Snansidansi/hue-api-go => ../../

require (
	github.com/Snansidansi/hue-api-go v1.0.3
	golang.org/x/term v0.38.0
)

require (
	github.com/joho/godotenv v1.5.1
	golang.org/x/sys v0.39.0 // indirect
)
