package inject

import (
	"github.com/Poul-george/go-api/api/core/domain/repository/auth/token"
	"github.com/Poul-george/go-api/api/core/domain/repository/user"
	tokenImpl "github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/repository/auth/token"
	userImpl "github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/repository/user"
)

func (i *Injector) userRepository() user.Repository {
	return userImpl.NewRepository(i.gormHandler())
}

func (i *Injector) tokenRepository() token.Repository {
	return tokenImpl.NewRepository(i.cacheHandler())
}
