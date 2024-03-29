package transport

// import (
// 	"context"

// 	wallet "github.com/atefeh-syf/yumigo/api/v1/pb/wallet"
// 	"github.com/atefeh-syf/yumigo/internal"
// 	"github.com/atefeh-syf/yumigo/pkg/wallet/endpoints"
// 	grpctransport "github.com/go-kit/kit/transport/grpc"
// )

// type grpcServer struct {
// 	get           grpctransport.Handler
// }
// func NewGRPCServer(ep endpoints.Set) wallet.WalletServer {
// 	return &grpcServer{
// 		get: grpctransport.NewServer(
// 			ep.GetEndpoint,
// 			decodeGRPCGetRequest,
// 			decodeGRPCGetResponse,
// 		),
// 	}
// }

// func (g *grpcServer) Get(ctx context.Context, r *wallet.GetRequest) (*wallet.GetReply, error) {
// 	_, rep, err := g.get.ServeGRPC(ctx, r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return rep.(*wallet.GetReply), nil
// }

// func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
// 	req := grpcReq.(*wallet.GetRequest)
// 	var filters []internal.Filter
// 	for _, f := range req.Filters {
// 		filters = append(filters, internal.Filter{Key: f.Key, Value: f.Value})
// 	}
// 	return endpoints.GetRequest{Filters: filters}, nil
// }

// func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
// 	reply := grpcReply.(*wallet.GetReply)
// 	//var docs []internal.Wallet
// 	var doc internal.Wallet
// 	for _, d := range reply.Wallets {
// 		doc = internal.Wallet{
// 			Name: d.Name,
// 			Type: d.Type,
// 			Balance: d.Balance,
// 			Slug: d.Slug,
// 			Description: d.Description,
// 			UserId: d.UserId,
// 		}
// 		//docs = append(docs, doc)
// 	}
// 	return endpoints.GetResponse{Wallet: doc, Err: reply.Err}, nil
// }


