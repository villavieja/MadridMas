package main

import (
	"fmt"
	"log"
	"net"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"MadridMas/server/incident"
	pb "MadridMas/server/proto"
)

const (
	useSSL = false
	// TODO(sara): Replace the username and password for environment variables.
	// Temporary allow anyone see this.
	dbCONNECTIONNAME = "madridmas-172613:europe-west1:madridmassql"
	// TODO(sara): Create a different user as root should never be used.
	// Changed the root password as it has been visible to others in github.
	dbUSER           = "root"
	dbPASSWORD       = "Madriz.1996.Nuria"
)

type MadridMasServer struct{}

var myServer MadridMasServer

func (s *MadridMasServer) CreateIncident(ctx context.Context, r *pb.CreateIncidentRequest) (*pb.CreateIncidentResponse, error) {

	grpclog.Printf("Received incident %+v", r)
	var i incident.Incident
	resp := &pb.CreateIncidentResponse{}

	//TODO(sara): check data is correct. Be careful of spammers.
	// We will reject incidents from unregistered users? anonymous?
	// We still don't know how to check the picture veracity.
	// i.Title = *r.Incident.Title
	i.Latitude = *r.Incident.Latitude
	i.Longitude = *r.Incident.Longitude
	i.Description = *r.Incident.Description
	// Fix the unknown error later.
	Title :=""

	db, err := mysql.DialPassword(dbCONNECTIONNAME, dbUSER, dbPASSWORD)
	if err != nil {
		resp.Error = proto.String(fmt.Sprintf("Could not open db: %v", err))
		return resp, err
	}
	defer db.Close()
	// TODO(sara): Move this to a transaction.

	query := fmt.Sprintf("INSERT INTO incident (fk_user,latitude,longitude,title,description,creation_date,status) VALUES(0, %f,%f,'%s',NOW(),0)", i.Latitude, i.Longitude, Title, i.Description)
	if _, err = db.Exec(query); err != nil {
		resp.Error = proto.String(fmt.Sprintf("Error on insert: %v", err))
	}
	return resp, err
}

func (s *MadridMasServer) ListIncidents(ctx context.Context, r *pb.ListIncidentsRequest) (*pb.ListIncidentsResponse, error) {
	res := &pb.ListIncidentsResponse{}
	res.Incident = make([]*pb.Incident,0,1)
		res.Incident = append (res.Incident,
				&pb.Incident {
					Description: proto.String("Unknown"),
				})
	return res, nil
}

func main() {
	var err error
	var lis net.Listener
	var grpcServer *grpc.Server
	if !useSSL {
		lis, err = net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer = grpc.NewServer()
	} else {
		certFile := "ssl.crt"
		keyFile := "ssl.key"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		lis, err = net.Listen("tcp", ":443")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer = grpc.NewServer(grpc.Creds(creds))
	}
	pb.RegisterMadridMasServer(grpcServer, &myServer)
	grpcServer.Serve(lis)
}
