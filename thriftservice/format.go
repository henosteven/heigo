package thriftservice

import (
	"context"
	"github.com/henosteven/heigo/heiThrift"
)

type FormatDataImpl struct{}

func (fdi FormatDataImpl) DoFormat (ctx context.Context, data *heiThrift.Data) (r *heiThrift.Data, err error) {
	return data, nil
}
