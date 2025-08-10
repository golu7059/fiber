package main


import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type car struct {
	ID int `gorm:"primaryKey"` 
	Name string
	Brand string
	Price int
	Year int
}

var db *gorm.DB

func ConnectDB() {
	var err error
	dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	err = db.AutoMigrate(&car{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}
}

func CreateCar(car *car) error {
	result := db.Create(car)
	return Result.Error	
}

func GetCar(id int) (*car, error) {
	result := db.First(&car{}, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &car{}, nil
}

func UpdatePrice(id, newPrice int) error {
	var tempCar car
	result := db.First(&tempCar, id)
	if result.Error != nil {
		return result.Error
	}
	tempCar.Price = newPrice
	result = db.Save(&tempCar)
	return result.Error
}

func DeleteCar(id int) error {
	var tempCar car
	result := db.First(&tempCar, id)
	if result.Error != nil {
		return result.Error
	}
	result = db.Delete(&tempCar)
	return result.Error
}

 func main() {
 }