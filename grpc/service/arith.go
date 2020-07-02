package service

import (
	"context"
)

//Arith defines the struct of arith.
type Arith struct{}

//Multiply operation
func (p *Arith) Multiply(ctx context.Context, req *ArithRequest) (*ArithResponse, error) {
	res := &ArithResponse{}
	res.Pro = req.GetA() * req.GetB()
	return res, nil
}
