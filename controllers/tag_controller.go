package controllers

import (
	"fmt"
	// "net/http"

	// "github.com/2gbeh/drygin-api/pipes"

	"github.com/2gbeh/drygin-api/models"
	"github.com/2gbeh/drygin-api/repositories"
	"github.com/2gbeh/drygin-api/utils"

	// "strconv"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

type TagController struct {
	tagRepository repositories.ITagRepository
}

func UseTagController(repository repositories.ITagRepository) *TagController {
	return &TagController{tagRepository: repository}
}

func (controller *TagController) Migrate(ctx *gin.Context) {	
	res, err := controller.tagRepository.Migrate()
	fmt.Printf("res-%v\n err-%v", res, err)
	
	ctx.Header("Content-Type", "application/json")
	if err != nil {
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, err)
	}	
}

func (controller *TagController) MigrateFresh(ctx *gin.Context) {	
	res, err := controller.tagRepository.DropTable()
	fmt.Printf("res-%v\n err-%v", res, err)

	ctx.Header("Content-Type", "application/json")
	if err != nil {
		controller.Migrate(ctx)
	} else {
		ctx.JSON(400, err)
	}	
}

func (controller *TagController) MigrateFreshSeed(ctx *gin.Context) {
	controller.MigrateFresh(ctx)

	var t models.Tags
	obj, err := utils.LoadJson[models.Tags]("./data/customers.json", t)
	fmt.Printf("obj-%v\n err-%v", obj, err)
	
	ctx.Header("Content-Type", "application/json")
	if err != nil {
		res, err := controller.tagRepository.InsertMany(obj.Data)
		if err != nil {
			ctx.JSON(200, res)
		} else {
			ctx.JSON(400, err)
		}	
	} else {
		ctx.JSON(500, err)
	}	
}

func (controller *TagController) All(ctx *gin.Context) {
	res, err := controller.tagRepository.SelectAll()
	fmt.Printf("res-%v\n err-%v", res, err)
	
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, res)

}
