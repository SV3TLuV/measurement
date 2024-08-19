package user

import (
	"context"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/pkg/errors"
	"measurements-api/internal/model"
	"measurements-api/internal/repository"
	model2 "measurements-api/internal/repository/user/model"
	def "measurements-api/internal/service"
	"measurements-api/internal/service/object/utils"
	"slices"
)

var _ def.UserService = (*service)(nil)

type service struct {
	repo            repository.UserRepository
	passwordService def.PasswordService
	trManager       *manager.Manager
}

func NewService(
	userRepo repository.UserRepository,
	passwordService def.PasswordService,
	trManager *manager.Manager) *service {
	return &service{
		repo:            userRepo,
		passwordService: passwordService,
		trManager:       trManager,
	}
}

func (u *service) GetUsers(ctx context.Context,
	options *model2.GetUsersQueryParams) (*model.PagedList[model.User], error) {
	users, count, err := u.repo.Get(ctx, options)
	if err != nil {
		return nil, errors.Wrap(err, "get users")
	}

	pagedList := &model.PagedList[model.User]{
		Page:     options.Page,
		PageSize: options.PageSize,
		Total:    *count,
		Items:    users,
	}

	return pagedList, nil
}

func (u *service) GetUser(ctx context.Context, id uint64) (*model.User, error) {
	return u.repo.GetById(ctx, id)
}

func (u *service) GetUserObjects(ctx context.Context, userID uint64) ([]*model.Object, error) {
	objects, err := u.repo.GetUserObjects(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "get user objects")
	}

	slices.SortFunc(objects, utils.SortObjectByAscID)
	for i := 0; i < len(objects); i++ {
		if objects[i].Children == nil {
			continue
		}

		slices.SortFunc(objects[i].Children, utils.SortObjectByAscID)
		for j := 0; j < len(objects[i].Children); j++ {
			if objects[i].Children[j].Children == nil {
				continue
			}

			slices.SortFunc(objects[i].Children[j].Children, utils.SortObjectByAscID)
		}
	}

	return objects, nil
}

func (u *service) GetUserColumns(ctx context.Context, userID uint64) ([]*model.Column, error) {
	return u.repo.GetUserColumns(ctx, userID)
}

func (u *service) GetUserPermissions(ctx context.Context, userID uint64) ([]*model.Permission, error) {
	return u.repo.GetUserPermissions(ctx, userID)
}

func (u *service) Create(ctx context.Context, user *model.User) error {
	userDb, err := u.repo.GetByLogin(ctx, user.Login)
	if err != nil && !errors.Is(err, model.NotFound) {
		return errors.Wrap(err, "get user by login")
	}
	if userDb != nil {
		return errors.New("login is busy")
	}

	passwordHash, err := u.passwordService.HashPassword(user.Password)
	if err != nil {
		return errors.Wrap(err, "hash password")
	}

	user.Password = *passwordHash

	err = u.trManager.Do(ctx, func(ctx context.Context) error {
		err = u.repo.Save(ctx, user)
		if err != nil {
			return errors.Wrap(err, "save user")
		}

		savedUser, err := u.repo.GetByLogin(ctx, user.Login)
		if err != nil {
			return errors.Wrap(err, "get saved user")
		}

		err = u.repo.UpdateUserPermissions(ctx, savedUser.ID, user.PermissionIds)
		if err != nil {
			return errors.Wrap(err, "update permissions")
		}

		err = u.repo.UpdateUserColumns(ctx, savedUser.ID, user.ColumnIds)
		if err != nil {
			return errors.Wrap(err, "update columns")
		}

		err = u.repo.UpdateUserPosts(ctx, savedUser.ID, user.PostIds)
		if err != nil {
			return errors.Wrap(err, "update posts")
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "transaction failed")
	}

	return nil
}

func (u *service) Update(ctx context.Context, user *model.User) error {
	userDb, err := u.repo.GetById(ctx, user.ID)
	if err != nil {
		return errors.Wrap(err, "get user")
	}

	userDb.Login = user.Login
	userDb.Role.ID = user.Role.ID

	err = u.trManager.Do(ctx, func(ctx context.Context) error {
		err = u.repo.Save(ctx, userDb)
		if err != nil {
			return errors.Wrap(err, "save user")
		}

		err = u.repo.UpdateUserPermissions(ctx, userDb.ID, user.PermissionIds)
		if err != nil {
			return errors.Wrap(err, "update permissions")
		}

		err = u.repo.UpdateUserColumns(ctx, userDb.ID, user.ColumnIds)
		if err != nil {
			return errors.Wrap(err, "update columns")
		}

		err = u.repo.UpdateUserPosts(ctx, userDb.ID, user.PostIds)
		if err != nil {
			return errors.Wrap(err, "update posts")
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "transaction failed")
	}

	return nil
}

func (u *service) ChangePassword(ctx context.Context, userID uint64, password string) error {
	userDb, err := u.repo.GetById(ctx, userID)
	if err != nil {
		return err
	}

	passwordHash, err := u.passwordService.HashPassword(password)
	if err != nil {
		return err
	}

	userDb.Password = *passwordHash
	err = u.repo.Save(ctx, userDb)
	return err
}

func (u *service) Ban(ctx context.Context, userID uint64) error {
	user, err := u.repo.GetById(ctx, userID)
	if err != nil {
		return err
	}

	user.IsBlocked = true

	err = u.repo.Save(ctx, user)
	return err
}

func (u *service) Unban(ctx context.Context, userID uint64) error {
	user, err := u.repo.GetById(ctx, userID)
	if err != nil {
		return err
	}

	user.IsBlocked = false

	err = u.repo.Save(ctx, user)
	return err
}

func (u *service) Delete(ctx context.Context, userID uint64) error {
	return u.repo.Delete(ctx, userID)
}
