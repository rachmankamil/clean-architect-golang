package drivers

import (
	categoryDomain "ca-amartha/businesses/category"
	categoryDB "ca-amartha/drivers/mysql/category"

	newsDomain "ca-amartha/businesses/news"
	newsDB "ca-amartha/drivers/mysql/news"

	userDomain "ca-amartha/businesses/users"
	userDB "ca-amartha/drivers/mysql/users"

	ipLocatorDomain "ca-amartha/businesses/iplocator"
	ipAPI "ca-amartha/drivers/thirdparties/ipapi"

	"gorm.io/gorm"
)

//NewCategoryRepository Factory with category domain
func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

//NewNewsRepository Factory with news domain
func NewNewsRepository(conn *gorm.DB) newsDomain.Repository {
	return newsDB.NewMySQLRepository(conn)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}

func NewIPLocatorRepository() ipLocatorDomain.Repository {
	return ipAPI.NewIpAPI()
}
