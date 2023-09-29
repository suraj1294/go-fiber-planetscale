package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/suraj1294/go-fiber-planetscale/repository"
)

type ProductHandler struct {
	productRepository *repository.ProductRepository
}

func (ph *ProductHandler) GetProducts(c *fiber.Ctx) error {

	products, err := ph.productRepository.GetAll()

	if err != nil {
		c.JSON(map[string]string{"message": "error"})
	}

	c.JSON(products)

	return nil

}

func (ph *ProductHandler) GetProduct(c *fiber.Ctx) error {

	id := c.Params("id")

	s, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest)
		c.JSON("bad request")
		return nil
	}

	product, selectedErr := ph.productRepository.GetById(s)

	if selectedErr != nil {
		c.Status(http.StatusNotFound)
		c.JSON("product not found")
		return nil
	}

	c.Status(http.StatusOK)

	c.JSON(product)

	return nil

}

func (product *ProductHandler) AddProduct(c *fiber.Ctx) error {

	newProduct := new(repository.Product)

	c.BodyParser(newProduct)

	addedProduct, selectErr := product.productRepository.Add(newProduct)

	if selectErr != nil {
		c.Status(http.StatusNotFound)
		c.JSON("product not found")
	}

	c.JSON(addedProduct)

	return nil

}

func (product *ProductHandler) UpdateProduct(c *fiber.Ctx) error {

	id := c.Params("id")

	productId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest)
		c.JSON("bad request")
		return nil
	}

	updateProduct := new(repository.Product)

	c.BodyParser(updateProduct)

	updatedProduct, updateErr := product.productRepository.Update(updateProduct, productId)

	if updateErr != nil {

		c.Status(http.StatusInternalServerError)
		c.JSON("failed to update product")

		return nil

	}

	c.JSON(updatedProduct)

	return nil

}

func (product *ProductHandler) DeleteProduct(c *fiber.Ctx) error {

	id := c.Params("id")

	productId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest)
		c.JSON("bad request")
		return nil
	}

	product.productRepository.Delete(productId)

	c.JSON("ok")

	return nil
}

func GetProductHandler() *ProductHandler {
	return &ProductHandler{productRepository: repository.GetProductRepository()}
}
