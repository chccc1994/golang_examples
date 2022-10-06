package service

import "context"

var ProductService = &productService{}

type productService struct {
}

func (p *productService) GetProductStock(context context.Context, request *ProductRequest) (*ProductResponse, error) {
	stock := p.GetStockById(request.ProdId)
	return &ProductResponse{ProdStock: stock}, nil
}

func (p *productService) GetStockById(id int32) int32 {
	return 100
}
