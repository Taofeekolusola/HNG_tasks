package main

import (
    "context"
    "log"

    "ent/ent"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    client, err := ent.Open("mysql", "<user>:<pass>@tcp(<host>:<port>)/<database>?parseTime=True")
    if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    }
    defer client.Close()
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
    u, err := client.User.
        Create().
        SetAge(30).
        SetName("a8m").
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed creating user: %w", err)
    }
    log.Println("user was created: ", u)
    return u, nil
}
func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
    u, err := client.User.
        Query().
        Where(user.Name("a8m")).
        // `Only` fails if no user found,
        // or more than 1 user returned.
        Only(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed querying user: %w", err)
    }
    log.Println("user returned: ", u)
    return u, nil
}
func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
    // Create a new car with model "Tesla".
    tesla, err := client.Car.
        Create().
        SetModel("Tesla").
        SetRegisteredAt(time.Now()).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed creating car: %w", err)
    }
    log.Println("car was created: ", tesla)

    // Create a new car with model "Ford".
    ford, err := client.Car.
        Create().
        SetModel("Ford").
        SetRegisteredAt(time.Now()).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed creating car: %w", err)
    }
    log.Println("car was created: ", ford)

    // Create a new user, and add it the 2 cars.
    a8m, err := client.User.
        Create().
        SetAge(30).
        SetName("a8m").
        AddCars(tesla, ford).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed creating user: %w", err)
    }
    log.Println("user was created: ", a8m)
    return a8m, nil
}
func QueryCars(ctx context.Context, a8m *ent.User) error {
    cars, err := a8m.QueryCars().All(ctx)
    if err != nil {
        return fmt.Errorf("failed querying user cars: %w", err)
    }
    log.Println("returned cars:", cars)

    // What about filtering specific cars.
    ford, err := a8m.QueryCars().
        Where(car.Model("Ford")).
        Only(ctx)
    if err != nil {
        return fmt.Errorf("failed querying user cars: %w", err)
    }
    log.Println(ford)
    return nil
}
func QueryCarUsers(ctx context.Context, a8m *ent.User) error {
    cars, err := a8m.QueryCars().All(ctx)
    if err != nil {
        return fmt.Errorf("failed querying user cars: %w", err)
    }
    // Query the inverse edge.
    for _, c := range cars {
        owner, err := c.QueryOwner().Only(ctx)
        if err != nil {
            return fmt.Errorf("failed querying car %q owner: %w", c.Model, err)
        }
        log.Printf("car %q owner: %q\n", c.Model, owner.Name)
    }
    return nil
}