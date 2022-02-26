package app

import (
	controller "github.com/Dryluigi/auction-system/controller/api/v1"
	"github.com/Dryluigi/auction-system/repository"
	service "github.com/Dryluigi/auction-system/service/api/v1"
)

var productRepository repository.ProductRepository
var bidSessionRepository repository.BidSessionRepository
var bidProductRepository repository.BidProductRepository

var bidSessionService service.BidSessionService
var productService service.ProductService
var biddingService service.BiddingService

var productController controller.ProductController
var bidSessionController controller.BidSessionController
