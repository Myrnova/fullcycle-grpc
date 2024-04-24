package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
		id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) List() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []Category{}

	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}

	return categories, nil
}

func (c *Category) FindOne(id string) (Category, error) {
	var category Category
	err := c.db.QueryRow("SELECT id, name, description FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name, &category.Description)

	if err != nil {
		return Category{}, err
	}

	return category, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
	var id, name, description string
	err := c.db.QueryRow(`SELECT c.id, c.name, c.description 
	FROM categories c
	INNER JOIN courses co ON co.category_id = c.id
	where co.id = $1`, courseID).Scan(&id, &name, &description)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}
