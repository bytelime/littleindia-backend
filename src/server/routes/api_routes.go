package routes

import (
	"github.com/kataras/iris"
)

func Call(ctx iris.Context) {

	content := ctx.PostValue("content")

	json(ctx, map[string]interface{}{
		"mensaje": content,
	})

}

func CreateCategory(ctx iris.Context){



}

func UpdateCategory(ctx iris.Context){

	

}

func DeleteCategory(ctx iris.Context){

	

}

func CreateSubCategory(ctx iris.Context){



}

func UpdateSubCategory(ctx iris.Context){

	

}

func DeleteSubcategory(ctx iris.Context){

	

}

func CreateProduct(ctx iris.Context){

	

}

//paginacion, por categoria y subcategoria
func GetProducts(ctx iris.Context){

	

}

func UpdateProduct(ctx iris.Context){

	

}

func DeleteProduct(ctx iris.Context){

	

}

func GetArchivedOrders(ctx iris.Context){

	

}

func ArchiveOrder(ctx iris.Context){

	

}
