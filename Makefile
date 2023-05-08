m_up:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5438/postgres?sslmode=disable' up
m_down:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5438/postgres?sslmode=disable' down
dc_up_build:
	docker compose up --build -d todo
dc_up:
	docker compose up -d todo