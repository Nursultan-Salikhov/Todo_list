m_up:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up
m_down:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' down