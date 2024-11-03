module project
// go get -u github.com/gorilla/mux // using this we can extract path variables
//  go get -u github.com/joho/godotenv (to load env file)
//  go get -u gorm.io/gorm(orm)
// go get -u gorm.io/driver/postgres  (to load the postgres driver)

// In Go, GORM serves as the Object-Relational Mapping (ORM) tool, similar to how JPA (Java Persistence API) works with Hibernate in Spring Boot. Like Hibernate, GORM handles database interactions, mapping Go structs to database tables and providing methods for CRUD operations. This lets you interact with the database using Go code rather than writing raw SQL queries.
go 1.23.2

require (
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.1 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	gorm.io/driver/postgres v1.5.9 // indirect
	gorm.io/gorm v1.25.12 // indirect
)
