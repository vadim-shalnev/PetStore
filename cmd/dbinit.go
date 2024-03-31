package main

import (
	"database/sql"
	"log"
)

func CreateTable(db *sql.DB) {

	createCategoryTable := `
CREATE TABLE IF NOT EXISTS categories (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL
    );
    `
	createPetTable := `
CREATE TABLE IF NOT EXISTS pets (
        id SERIAL PRIMARY KEY,
    	FOREIGN KEY (category_id) REFERENCES categories(id),
    	FOREIGN KEY (owner_id) REFERENCES users(id),
    	name VARCHAR(255) NOT NULL,
    	status VARCHAR(255) NOT NULL,
    	created_at TIMESTAMP NOT NULL,
    	deleted_at TIMESTAMP DEFAULT NULL
    );
    `
	createUserTable := `
CREATE TABLE IF NOT EXISTS users (
    	id SERIAL PRIMARY KEY,
    	username VARCHAR(255) NOT NULL UNIQUE,
    	first_name VARCHAR(255),
    	last_name VARCHAR(255),
    	email VARCHAR(255) NOT NULL UNIQUE,
    	password VARCHAR(255) NOT NULL,
    	phone VARCHAR(255),
    	user_status INT NOT NULL,
    	created_at TIMESTAMP NOT NULL,
    	deleted_at TIMESTAMP DEFAULT NULL
    );
    `
	createOrderTable := `
CREATE TABLE IF NOT EXISTS orders (
        id SERIAL PRIMARY KEY,
    	FOREIGN KEY (pet_id) REFERENCES pets(id),
    	FOREIGN KEY (seller_id) REFERENCES users(id),
    	FOREIGN KEY (buyer_id) REFERENCES users(id),
    	quantity INT NOT NULL,
    	status VARCHAR(255) NOT NULL,
    	complete BOOLEAN NOT NULL,
    	created_at TIMESTAMP NOT NULL,
    	deleted_at TIMESTAMP DEFAULT NULL
    );
    `
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Failed to start transaction:", err)
	}

	_, err = tx.Exec(createCategoryTable)
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to create category table:", err)
	}

	_, err = tx.Exec(createPetTable)
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to create pet table:", err)
	}

	_, err = tx.Exec(createUserTable)
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to create users table:", err)
	}

	_, err = tx.Exec(createOrderTable)
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to create order table:", err)
	}
	err = CreateCategories(db)
	if err != nil {
		log.Fatal("Failed to create categories:", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Failed to commit transaction:", err)
	}

	log.Println("Tables created successfully")
}

// CreateCategories создаем категории питомцев
func CreateCategories(db *sql.DB) error {
	// ConnectionDB Подключаемся к бд
	query := `
	INSERT INTO categories (name) VALUES
	('Собаки'),
	('Кошки'),
	('Птицы'),
	('Рыбки'),
	('Хомяки'),
	('Грызуны'),
	('Рептилии'),
	('Аквариумные рыбы'),
	('Павлины'),
	('Собаки-поводыри');
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to insert categories: %v", err)
		return err
	}
	log.Println("Categories inserted successfully")

	return nil
}
