run:
	go run ./examples

.PHONY: example-live
example-live:
	@cd examples/live-display && go run .
