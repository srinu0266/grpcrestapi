package v1

import (
	"database/sql"

	"github.com/golang/protobuf/ptypes"
	"github.com/srinu0266/grpcrestapi/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
const (
	apiVersion = "v1"
)

type toDoServiceServer struct {
	db *sql.DB
}

func NewToDoServiceServer(db *sql.DB) v1.toDoServiceServer {
	return &toDoServiceServer{db: db}
}


func (s *toDoServiceServer)checkAPI(api string) error {
	if len(api)>0{
		if apiVersion!=api{
			return status.Errorf(codes.Unimplemented,"unsupported API version: service implements API version '%s', but asked for '%s'",apiVersion, api)
		}
	}
	return nil
}

func (s *toDoServiceServer)connect(ctx context.Context) (*sql.Conn,error) {
	c,err:=s.db.Conn(ctx)

	if err!=nil{
		return nil,status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}

	return c,nil
}


func (s *toDoServiceServer)Create(ctx context.Context,req *v1.CreateRequest) (*v1.CreateResponse, error ) {
	if err:=s.checkAPI(req.Api);err!=nil{
		return nil,sql.Err
	}

	c,err:=s.connect(ctx)
	if err!=nil{
		return nil,err
	}

	defer c.Close()


	reminder,err:=ptypes.Timestamp(req.ToDo.Reminder)
	if err!=nil{
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}

	res,err:=c.ExecContext(ctx, "INSERT INTO ToDo(`Title`, `Description`, `Reminder`) VALUES(?, ?, ?)",
	req.ToDo.Title, req.ToDo.Description, reminder)

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into ToDo-> "+err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ToDo-> "+err.Error())
	}


	return &v1.CreateResponse{
		Api:apiVersion,
		Id:id
	}

}





