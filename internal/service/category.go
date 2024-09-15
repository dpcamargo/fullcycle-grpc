package service

import (
	"context"
	"errors"
	"io"

	"github.com/dpcamargo/fullcycle-grpc/internal/database"
	"github.com/dpcamargo/fullcycle-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB *database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: *categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.GetName(), in.GetDescription())
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{Id: category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}

func (c *CategoryService) ListCategories(_ context.Context, _ *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categoryList []*pb.Category
	for _, category := range categories {
		category := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
		categoryList = append(categoryList, category)
	}

	return &pb.CategoryList{
		Categories: categoryList,
	}, nil
}

func (c *CategoryService) GetCategory(_ context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindByID(in.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}

	for {
		category, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}
		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
	}
}
