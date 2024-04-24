package main

import (
	"database/sql"
	"fmt"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/devfullcycle/14-gRPC/internal/database"
	"github.com/devfullcycle/14-gRPC/internal/pb"
	"github.com/devfullcycle/14-gRPC/internal/service"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	courseDB := database.NewCourse(db)
	categoryService := service.NewCategoryService(*categoryDB)
	courseService := service.NewCourseService(*courseDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	pb.RegisterCourseServiceServer(grpcServer, courseService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server running at port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
