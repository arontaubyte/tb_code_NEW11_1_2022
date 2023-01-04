package lib

import (
	"github.com/taubyte/go-sdk/event"
	"bitbucket.org/taubyte/go-sdk/database"
)

//export ping
func ping(e event.Event) uint32 {
	database, err := database.New("failed_database")
	if err != nil {
		return err
	}
}
