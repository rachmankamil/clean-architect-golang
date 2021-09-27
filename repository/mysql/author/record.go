package author

import (
	"kampus-merdeka-ca/bussiness/author"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name      string
	Publisher string
}

func toDomain(rec Author) author.Domain {
	return author.Domain{
		Name:      rec.Name,
		Publisher: rec.Publisher,
	}
}

// func fromDomain(domain author.Domain) Record {
// 	return Record{
// 		Name:      domain.Name,
// 		Publisher: domain.Publisher,
// 	}
// }
