package lib

import (
	"fmt"

	"github.com/taubyte/go-sdk/event"
	"bitbucket.org/taubyte/go-sdk/database"
)

//export ping
func ping(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	err = testDatabase(h)
	if err != nil {
		h.Write([]byte(fmt.Sprintf("Failed: %s", err.Error())))
		return 1
	}

	return 0
}

func testDatabase(h event.HttpEvent) error{
	_database, err := database.New("failed_db")
	if err != nil {
		return err
	}

	fmt.Println(_database)
	return nil
}