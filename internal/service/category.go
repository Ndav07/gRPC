package service

import (
	"context"

	"github.com/Ndav07/gRPC/internal/database"
	"github.com/Ndav07/gRPC/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, body *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(body.Name, body.Description); if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}
	
	return categoryResponse, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, null *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll(); if err != nil {
		return nil, err
	}

	var categoriesReponse []*pb.Category

	for _, category := range categories {
		categoryReponse := &pb.Category{
			Id: category.ID,
			Name: category.Name,
			Description: category.Description,
		}
		categoriesReponse = append(categoriesReponse, categoryReponse)
	}

	return &pb.CategoryList{Category: categoriesReponse}, nil
}

func (c *CategoryService) FindCategory(ctx context.Context, body *pb.CategoryFindRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Find(body.Id); if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}
	
	return categoryResponse, nil
}