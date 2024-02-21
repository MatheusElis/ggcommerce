package api

import (
	"github.com/MatheusElis/ggcommerce/types"
	"github.com/anthdm/weavebox"
)

type ProductHandler struct {

}

func (h *ProductHandler) HandleGetProduct(c *weavebox.Context) error{
  return c.JSON(&types.Product{})
}
