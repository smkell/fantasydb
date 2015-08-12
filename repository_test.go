package main

import (
	"testing"
)

// TestUpsertPlayer tests the upsert player method of the mongodb repository
func TestUpsertPlayer(t *testing.T) {
	repo := NewMongoDBRepository("localhost/fantasydb_test")

	// Need to clean out the test database
	err := repo.RemoveAllPlayers()
	if err != nil {
		t.Fatal("An error occured initializing the test case", err)
	}

	players, err := repo.GetAllPlayers()
	if err != nil {
		t.Fatal("An error occured initializing the test case.", err)
	}

	if len(players) > 0 {
		t.Fatalf("An error occured initializing the test case. Expected there to be no players in database but there were %d", len(players))
	}

	input := Player{
		FirstName: "Andrew",
		LastName:  "Luck",
		Position:  "QB",
	}

	player, err := repo.UpsertPlayer(input)

	if err != nil {
		t.Error("An unexpected error occured", err)
	}

	if input.ID == player.ID {
		t.Errorf("Expected the repository to set the player's ID but did not. (input.ID = %v, actual.ID = %v)", input.ID, player.ID)
	}

	updated, err := repo.UpsertPlayer(player)
	if err != nil {
		t.Error("An unexpected error occured.", err)
	}

	if player.ID != updated.ID {
		t.Errorf("Expected the result of the upser to be idempotent when the input is unchanged. (input.ID = %v, actual.ID =%v)", player.ID, updated.ID)
	}

	updated, err = repo.UpsertPlayer(input)
	if err != nil {
		t.Error("An unexpected error occured.", err)
	}

	if player.ID != updated.ID {
		t.Errorf("Expected the result of the upsert to be idempotent when the input is unchanged. (input.ID = %v, actual.ID =%v)", player.ID, updated.ID)
	}
}
