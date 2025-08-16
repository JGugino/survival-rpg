default:
	@echo "Building executable (this might take a minute)"
	@go build -o ./build/survival_rpg
	@echo "Starting executable"
	@./build/survival_rpg

dev:
	@echo "Starting in dev mode"
	@air run main.go
