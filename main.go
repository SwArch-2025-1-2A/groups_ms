package main

import (
	"context"
	"log"
	"os"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/SwArch-2025-1-2A/backend/repository"

	"github.com/google/uuid"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := repository.New(conn)

	// list all events
	events, err := queries.GetEvents(ctx)
	if err != nil {
		return err
	}
	log.Println(events)

	var id pgtype.UUID = pgtype.UUID{}

	if err := id.Scan(uuid.New().String()); err != nil {
		log.Fatalln(err)
	}

	// create a user
	insertedUser, err := queries.CreateUser(ctx, repository.CreateUserParams{
		ID:   id,
		Name: "John Doe",
		ProfilePic: pgtype.Text{
			String: "https://example.com/profile.jpg"},
	})
	if err != nil {
		return err
	}
	log.Println("inserted:", insertedUser)

	// get the user we just inserted
	fetchedUser, err := queries.GetUserById(ctx, insertedUser.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedUser, fetchedUser))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
