package usecase

import (
	"context"
	"github.com/matheusvidal21/unit-of-work/internal/entity"
	"github.com/matheusvidal21/unit-of-work/internal/repository"
	"github.com/matheusvidal21/unit-of-work/pkg"
)

type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCaseUow struct {
	Uow pkg.UowInterface
}

func NewAddCourseUseCaseUow(uow pkg.UowInterface) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{
		Uow: uow,
	}
}

func (a *AddCourseUseCaseUow) Execute(ctx context.Context, input InputUseCaseUow) error {
	return a.Uow.Do(ctx, func(uow pkg.UowInterface) error {
		category := entity.Category{
			Name: input.CategoryName,
		}

		err := a.getCategoryRepository(ctx).Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}
		err = a.getCourseRepository(ctx).Insert(ctx, course)
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *AddCourseUseCaseUow) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CategoryRepositoryInterface)
}

func (a *AddCourseUseCaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CourseRepositoryInterface)
}
