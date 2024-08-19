package object

import (
	"context"
	"github.com/pkg/errors"
	"measurements-api/internal/interfaces/converter"
	"measurements-api/internal/model"
	"measurements-api/internal/repository"
	model2 "measurements-api/internal/repository/object/model"
	def "measurements-api/internal/service"
	"measurements-api/pkg/asoiza"
)

var _ def.ObjectService = (*service)(nil)

type service struct {
	objectRepo   repository.ObjectRepository
	postInfoRepo repository.PostInfoRepository
	asoizaClient asoiza.Client
}

func NewService(
	objectRepo repository.ObjectRepository,
	postInfoRepo repository.PostInfoRepository,
	asoizaClient asoiza.Client) *service {
	return &service{
		objectRepo:   objectRepo,
		postInfoRepo: postInfoRepo,
		asoizaClient: asoizaClient,
	}
}

func (s *service) GetObjects(ctx context.Context, options *model2.GetObjectsQueryParams) ([]*model.Object, error) {
	return s.objectRepo.Get(ctx, options)
}

func (s *service) GetPosts(ctx context.Context) ([]*model.Object, error) {
	postKey := uint64(model.PostKey)
	return s.GetObjects(ctx, &model2.GetObjectsQueryParams{
		TypeID: &postKey,
	})
}

func (s *service) GetPost(ctx context.Context, userID, objectID uint64) (*model.Object, error) {
	object, err := s.objectRepo.GetUserPostById(ctx, userID, objectID)
	if err != nil {
		return nil, err
	}
	if object == nil {
		return nil, errors.Wrap(model.NotFound, "post")
	}

	return object, nil
}

func (s *service) GetTotalPostCount(ctx context.Context) (uint64, error) {
	typeID := uint64(model.PostKey)
	return s.objectRepo.GetCount(ctx, &model2.GetObjectCountParams{
		TypeID: &typeID,
	})
}

func (s *service) GetListenedPostCount(ctx context.Context) (uint64, error) {
	typeID := uint64(model.PostKey)
	enabled := true
	return s.objectRepo.GetCount(ctx, &model2.GetObjectCountParams{
		TypeID:     &typeID,
		IsListened: &enabled,
	})
}

func (s *service) SearchNew(ctx context.Context) ([]*model.Object, error) {
	laboratories, err := s.asoizaClient.GetNavTree(ctx, "root")
	if err != nil {
		return nil, errors.Wrap(err, "fetch laboratories")
	}

	// TODO: try optimize with gorutine
	cities := make([]*asoiza.Node, 0)
	for i := 0; i < len(laboratories); i++ {
		objects, err := s.asoizaClient.GetNavTree(ctx, laboratories[i].ID)
		if err != nil {
			return nil, errors.Wrap(err, "fetch cities")
		}
		cities = append(cities, objects...)
	}

	posts := make([]*asoiza.Node, 0)
	for i := 0; i < len(cities); i++ {
		objects, err := s.asoizaClient.GetNavTree(ctx, cities[i].ID)
		if err != nil {
			return nil, errors.Wrap(err, "fetch posts")
		}
		posts = append(posts, objects...)
	}

	objects := make([]*model.Object, 0)
	for i := 0; i < len(laboratories); i++ {
		objects = append(objects, converter.ToObjectFromAsoiza(laboratories[i]))
	}

	for i := 0; i < len(cities); i++ {
		objects = append(objects, converter.ToObjectFromAsoiza(cities[i]))
	}

	for i := 0; i < len(posts); i++ {
		objects = append(objects, converter.ToObjectFromAsoiza(posts[i]))
	}

	inserted, err := s.objectRepo.Save(ctx, objects)
	if err != nil {
		return nil, errors.Wrap(err, "save objects")
	}

	postInfos := make([]*model.PostInfo, 0)
	for i := 0; i < len(inserted); i++ {
		object := inserted[i]
		if object.Type.ID == uint64(model.PostKey) {
			postInfos = append(postInfos, &model.PostInfo{
				ObjectID:            object.ID,
				LastPollingDateTime: nil,
				IsListened:          true,
			})
		}
	}

	err = s.postInfoRepo.Save(ctx, postInfos)
	if err != nil {
		return nil, errors.Wrap(err, "save post_infos")
	}

	return inserted, nil
}

func (s *service) Enable(ctx context.Context, id uint64) error {
	postInfo, err := s.postInfoRepo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if postInfo == nil {
		return errors.Wrap(model.NotFound, "postInfo")
	}

	postInfo.IsListened = true

	return s.postInfoRepo.SaveOne(ctx, postInfo)
}

func (s *service) Disable(ctx context.Context, id uint64) error {
	postInfo, err := s.postInfoRepo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if postInfo == nil {
		return errors.Wrap(model.NotFound, "postInfo")
	}

	postInfo.IsListened = false

	return s.postInfoRepo.SaveOne(ctx, postInfo)
}
