package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// A Repository is an object used to store and query players.
type Repository interface {
	GetAllPlayers() ([]Player, error)
	GetPlayerByName(fistName, lastName string) (Player, error)
	UpsertPlayer(player Player) (Player, error)
	RemoveAllPlayers() error
}

type mongoDBRepository struct {
	connectionString string
}

// NewMongoDBRepository constructs a new repository which uses MongoDB as the
// backend.
func NewMongoDBRepository(connectionString string) Repository {
	return mongoDBRepository{connectionString}
}

func (repo mongoDBRepository) GetAllPlayers() ([]Player, error) {
	session, err := mgo.Dial(repo.connectionString)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	db := session.DB("")

	var players []Player
	err = db.C("players").Find(nil).All(&players)
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (repo mongoDBRepository) GetPlayerByName(firstName, lastName string) (Player, error) {
	session, err := mgo.Dial(repo.connectionString)
	if err != nil {
		return Player{}, err
	}
	defer session.Close()

	db := session.DB("")

	var player Player
	err = db.C("players").Find(bson.M{"first_name": firstName, "last_name": lastName}).One(&player)
	if err != nil {
		return Player{}, err
	}

	return player, nil
}

func (repo mongoDBRepository) UpsertPlayer(player Player) (Player, error) {
	session, err := mgo.Dial(repo.connectionString)
	if err != nil {
		return Player{}, err
	}
	defer session.Close()

	db := session.DB("")

	exist, err := repo.GetPlayerByName(player.FirstName, player.LastName)
	if err != nil {
		player.ID = bson.NewObjectId().Hex()
	} else {
		player.ID = exist.ID
	}

	_, err = db.C("players").UpsertId(player.ID, player)
	if err != nil {
		return Player{}, err
	}

	return player, nil
}

func (repo mongoDBRepository) RemoveAllPlayers() error {
	session, err := mgo.Dial(repo.connectionString)
	if err != nil {
		return err
	}
	defer session.Close()

	db := session.DB("")

	_, err = db.C("players").RemoveAll(nil)
	if err != nil {
		return err
	}
	return nil
}
