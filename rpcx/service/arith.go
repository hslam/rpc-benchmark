package service

import (
	"context"
)

//Arith defines the struct of arith.
type Arith struct{}

//Multiply operation
func (a *Arith) Multiply(ctx context.Context, req *ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}
