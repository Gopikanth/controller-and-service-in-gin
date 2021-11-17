package controller
import (
	"github.com/gin-gonic/gin"
	"myapp/service"
   "myapp/golang-gin-poc/entity"
)
type Customer_service struct{
	Save(ctx *gin.Context) entity.Customer
	View() []entitty.Customer
}
type controller struct {
	s s.Customer_service
}
func  NEW(s s.Customer_service)  Customer_service {
	return &controller{
		s: s,
	}

	
}
func (c *controller) View() []entity.Video {
	return c.service.View()
}

func (c *controller) Save(ctx *gin.Context) entity.Customer {
	var customer entity.Customer
	ctx.BindJSON(&customer)
	c.service.Save(customer)
	return customer
}
