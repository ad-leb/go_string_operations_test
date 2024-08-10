.SUFFIXES: 	# fuck implicit rules  >_<
.PHONY: all not_all build


PROJECT::= string_calculator
MAIN::= main.go



all: 
	go run $(PROJECT)
not_all:
	go run $(MAIN)
build:
	go build
