package inject

import (
	"github.com/Poul-george/go-api/api/core/domain/repository/user"
	userImpl "github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/repository/user"
)

func (i *Injector) userRepository() user.Repository {
	return userImpl.NewRepository(i.gormHandler())
}
