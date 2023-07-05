package services

import (
	"context"
	"time"
)

type Coffee struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Roast string `json:"roast"`
    Image string `json:"image"`
    Region string `json:"region"`
    Price float32 `json:"price"`
    GrindUnit int16 `json:"grind_unit"` 
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (c *Coffee) GetAllCoffees() ([]*Coffee, error) {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel() 

    query := `SELECT id, name, image, roast, region, price, grind_unit, created_at, updated_at FROM coffees`
    rows, err := db.QueryContext(ctx, query) 

    if err != nil {
        return nil, err
    }

    var coffees []*Coffee
    for rows.Next() {
        var coffee Coffee
        err := rows.Scan(
            &coffee.ID,
            &coffee.Name,
            &coffee.Image,
            &coffee.Roast,
            &coffee.Region,
            &coffee.Price,
            &coffee.GrindUnit,
            &coffee.CreatedAt,
            &coffee.UpdatedAt,
        )

        if err != nil {
            return nil, err
        }
        coffees = append(coffees, &coffee)
    }

    return coffees, nil
}  

func (c *Coffee) GetCoffeeById(id string) (*Coffee, error) {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel()
    query := `
        SELECT id, name, image, roast, region, price, grind_unit, created_at, updated_at 
        FROM coffees
        WHERE id = $1
    `
    var coffee Coffee

    row := db.QueryRowContext(ctx, query, id)
    err := row.Scan(
        &coffee.ID,
        &coffee.Name,
        &coffee.Image,
        &coffee.Roast,
        &coffee.Region,
        &coffee.Price,
        &coffee.GrindUnit,
        &coffee.CreatedAt,
        &coffee.UpdatedAt,
    )

    if err != nil {
        return nil, err
    }
    return &coffee, nil
}

func (c *Coffee) CreateCoffee(coffee Coffee) (*Coffee, error) {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel()

    query := `
        INSERT INTO coffees (name, image, region, roast, price, grind_unit, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8) returning *
    `
    
    _, err := db.ExecContext(
        ctx,
        query,
        coffee.Name,
        coffee.Image,
        coffee.Region,
        coffee.Roast,
        coffee.Price,
        coffee.GrindUnit,
        time.Now(),
        time.Now(),
    )

    if err != nil {
        return nil, err
    }

    return &coffee, nil
}

func (c *Coffee) UpdateCoffee(id string, body Coffee) (*Coffee, error) {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel()

    query := `
        UPDATE coffees
        SET
            name = $1,
            image = $2,
            roast = $3,
            price = $4,
            grind_unit = $5,
            region = $6,
            updated_at = $7
        WHERE id = $8
        returning *
    `

    _, err := db.ExecContext(
        ctx,
        query,
        body.Name,
        body.Image,
        body.Roast,
        body.Price,
        body.GrindUnit,
        body.Region,
        time.Now(), 
        id,
    )
    if err != nil {
        return nil, err
    }
    return &body, nil
}

func (c *Coffee) DeleteCoffee(id string) error {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel()

    query := `DELETE FROM coffees WHERE id = $1`
    _, err := db.ExecContext(ctx, query, id)
    if err != nil {
        return err
    }
    return nil
}
