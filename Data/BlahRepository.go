package Data
import (
	"github.com/FanszHub/test-site/Models"
)

type BlahRepository interface {
	AllBlahs() ([]*Models.Blah, error)
	AddBlah(*Models.Blah) (error)
}

func (db *DB) AllBlahs() ([]*Models.Blah, error) {

	var blahs []*Models.Blah

	err := db.C("Blahs").Find(nil).All(&blahs)

	if err != nil {
		return nil, err
	}

	return blahs, nil
}

func (db *DB) AddBlah(blah *Models.Blah) (error) {
	return db.C("Blahs").Insert(blah)
}