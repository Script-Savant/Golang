/*
Seed sample products into the db
1. Check if product already exists
2. Create sample product if none exist
3. Log seeding result
*/

package db

import "log"

func SeedProducts() {
	// 1. Check if products already exist
	var count int64
	DB.Model(&models.Product{}).Count(&count)
	if count > 0 {
		log.Println("Products already seeded, skipping")
		return
	}

	// 2. Create sample products
	products := []models.Product{
		{
			Name:        "Laptop",
			Description: "High performance laptop with 16GB RAM",
			Price:       999.99,
			Stock:       10,
		},
		{
			Name:        "Smartphone",
			Description: "Latest smartphone with 128GB storage",
			Price:       699.99,
			Stock:       20,
		},
		{
			Name:        "Headphones",
			Description: "Noise cancelling wireless headphones",
			Price:       199.99,
			Stock:       30,
		},
		{
			Name:        "Smart Watch",
			Description: "Fitness tracking smart watch",
			Price:       249.99,
			Stock:       15,
		},
	}

	// 3. Insert products
	result := DB.Create(&products)
	if result.Error != nil {
		log.Printf("Failed to seed products: %v", result.Error)
		return
	}

	log.Printf("Seeded %d products successfully", len(products))
}