package db

import (
	"gopkg.in/testfixtures.v2"
	"os"
	"testing"
)

const testDBDSN = "root@/bh_hotel_test?parseTime=true"

func TestMain(m *testing.M) {

	if err := NewDB(testDBDSN); err != nil {
		panic(err)
	}

	code := m.Run()

	os.Exit(code)
}

func PrepareFixture(path string) {
	fixture, err := testfixtures.NewFolder(db.DB.DB, &testfixtures.MySQL{}, path)
	if err != nil {
		panic(err)
	}
	if err := fixture.Load(); err != nil {
		panic(err)
	}
}
