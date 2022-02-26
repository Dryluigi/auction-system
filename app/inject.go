package app

import (
	controller "github.com/Dryluigi/auction-system/controller/api/v1"
	"github.com/Dryluigi/auction-system/database"
	"github.com/Dryluigi/auction-system/repository"
	service "github.com/Dryluigi/auction-system/service/api/v1"
)

func doInjection() {
	productRepository = &repository.ProductRepositoryImpl{DB: database.DB}
	bidSessionRepository = &repository.BidSessionRepositoryImpl{DB: database.DB}
	bidProductRepository = &repository.BidProductRepositoryImpl{DB: database.DB}

	bidSessionService = &service.BidSessionServiceImpl{BidSessionRepository: bidSessionRepository}
	productService = &service.ProductServiceImpl{ProductRepository: productRepository}
	biddingService = &service.BiddingServiceImpl{BidProductRepository: bidProductRepository}

	productController = &controller.ProductControllerImpl{ProductService: productService}
	bidSessionController = &controller.BidSessionControllerImpl{BidSessionService: bidSessionService, BiddingService: biddingService}
}
