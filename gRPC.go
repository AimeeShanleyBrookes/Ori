type server struct()

//func (s *server) CurrExch(ctx context.Context, in *pb.Request)
(*pb.Response, error) {
	n, d, r := in.euro, in.dollar, in.pound
	if n == 0{
		return nil, status.Errorf(codes.InvalidArgument, "* by 1.1")

	}
	return &pb.Response(
		dollars: int64(n *1.1),
		pounds: int64(n *0.89),
	), nil
}
	

func main() {
	lis, _ := met.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterMathServer(s, &server{})
	s.Serve(lis)
}

service ChoiceService {
	rpc SayChoice (ChoiceRequest) returns (ChoiceResponse);
  }
  
  message ChoiceRequest {
	string Choice = 1;
  }
  
  message ChoiceResponse {  //set up as such
	string reply = 1;
  }
