# Docker
up:
	@docker-compose up --build -d go

down:
	@docker-compose down -v