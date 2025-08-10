package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

var db *sqlx.DB

func connectDB() {
	dsn := "user=username password=password dbname=mydb sslmode=disable"
	var err error
	dbTemp, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Error in connecting to db: %v", err)
	}
	db = dbTemp
	fmt.Println("Connected to the database successfully!")
}

type Car struct {
	ID    int     `db:"id"`
	Name  string  `db:"name"`
	Model string  `db:"model"`
	Brand string  `db:"brand"`
	Year  int     `db:"year"`
	Price float64 `db:"price"`
}

func InsertCar(car Car) {
	query := `
		INSERT INTO cars (name, model, brand, year, price) VALUES (:name, :model, :brand, :year, :price)
	`
	_, err := db.NamedExec(query, car)
	if err != nil {
		fmt.Printf("Error in inserting car: %v\n", err)
		return
	}
	fmt.Println("Car inserted successfully")
}


func GetCar(id int) {
	var car Car
	query := `SELECT * FROM cars WHERE id = $1`
	err := db.Get(&car, query, id)
	if err != nil {
		fmt.Printf("Error in getting car by Id : %v\n", err)
		return
	}
	fmt.Printf("Fetched car : %+v\n", car)
}

func UpdateCar(car Car) {
	query := `
		UPDATE cars SET name=:name, model=:model, brand=:brand, year=:year, price=:price
		WHERE id=:id
	`
	_, err := db.NamedExec(query, car)
	if err != nil {
		fmt.Printf("Unable to update car : %v\n", err)
		return
	}
	fmt.Println("Car updated successfully!")
}
func DeleteCar(id int) {
	query := `DELETE FROM cars WHERE id=$1`
	_, err := db.Exec(query, id)
	if err != nil {
		fmt.Printf("Unable to delete the car: %v\n", err)
		return
	}
	fmt.Println("Car deleted Successfully!")
}

func main() {
	connectDB()
	// Example usage:
	// InsertCar(Car{Name: "Test", Model: "2024", Brand: "BrandX", Year: 2024, Price: 10000})
	// GetCar(1)
	// UpdateCar(Car{ID: 1, Name: "Updated", Model: "2024", Brand: "BrandX", Year: 2024, Price: 12000})
	// DeleteCar(1)
}