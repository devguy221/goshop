package routers

import (
	"goshop/repositories"
	"goshop/services"
)

var roleService services.RoleService
var userService services.UserService
var categoryService services.CategorySerivce
var productService services.ProductService
var warehouseService services.WarehouseSerivce

func init() {
	roleRepo := repositories.NewRoleRepository()
	userRepo := repositories.NewUserRepository()
	categoryRepo := repositories.NewCategoryRepository()
	productRepo := repositories.NewProductRepository()
	warehouseRepo := repositories.NewWarehouseRepository()

	roleService = services.NewService(roleRepo)
	userService = services.NewUserService(userRepo)
	categoryService = services.NewCategoryService(categoryRepo)
	productService = services.NewProductService(productRepo)
	warehouseService = services.NewWarehouseService(warehouseRepo)
}
