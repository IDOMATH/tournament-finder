package repository

import "github.com/IDOMATH/session/memorystore"

type Repository struct {
	TH      TournamentHandler
	UH      UserHandler
	Session *memorystore.MemoryStore
}
