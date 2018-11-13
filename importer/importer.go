package importer

import (
	"bufio"
	"io"
	"log"

	"github.com/Songmu/go-ltsv"
	"github.com/mono0x/puroland-greeting/model"

	"github.com/jmoiron/sqlx"
)

type Importer interface {
	Import(r io.Reader) error
}

type importerImpl struct {
	db *sqlx.DB
}

func NewImporter(db *sqlx.DB) Importer {
	return &importerImpl{
		db: db,
	}
}

func (i *importerImpl) Import(r io.Reader) error {
	var data []*model.RawData

	s := bufio.NewScanner(r)
	for s.Scan() {
		t := s.Text()

		var d model.RawData
		if err := ltsv.Unmarshal([]byte(t), &d); err != nil {
			return err
		}
		data = append(data, &d)
	}
	if err := s.Err(); err != nil {
		return err
	}

	log.Println(data)

	return nil
}
