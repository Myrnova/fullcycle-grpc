package service

import (
	"context"
	"io"

	"github.com/devfullcycle/14-gRPC/internal/database"
	"github.com/devfullcycle/14-gRPC/internal/pb"
)

type CourseService struct {
	pb.UnimplementedCourseServiceServer
	CourseDB database.Course
}

func NewCourseService(CourseDB database.Course) *CourseService {
	return &CourseService{
		CourseDB: CourseDB,
	}
}

func (c *CourseService) CreateCourse(ctx context.Context, in *pb.CreateCourseRequest) (*pb.CourseResponse, error) {
	course, err := c.CourseDB.Create(in.Name, in.Description, in.CategoryId)
	if err != nil {
		return nil, err
	}
	courseResponse := &pb.Course{
		Id:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		CategoryId:  course.CategoryID,
	}

	return &pb.CourseResponse{
		Course: courseResponse,
	}, nil
}

func (c *CourseService) ListCourses(ctx context.Context, in *pb.Blank) (*pb.CourseList, error) {
	courses, err := c.CourseDB.List()
	if err != nil {
		return nil, err
	}
	var coursesResponse []*pb.Course
	for _, course := range courses {
		coursesResponse = append(coursesResponse, &pb.Course{
			Id:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CategoryId:  course.CategoryID,
		})
	}
	return &pb.CourseList{
		Courses: coursesResponse,
	}, nil
}

func (c *CourseService) GetCourse(ctx context.Context, in *pb.CourseGetRequest) (*pb.CourseResponse, error) {
	course, err := c.CourseDB.FindOne(in.Id)
	if err != nil {
		return nil, err
	}

	courseResponse := &pb.Course{
		Id:          course.ID,
		Name:        course.Name,
		Description: course.Description,
	}

	return &pb.CourseResponse{
		Course: courseResponse,
	}, nil
}

func (c *CourseService) CreateCourseStream(stream pb.CourseService_CreateCourseStreamServer) error {
	courses := &pb.CourseList{}

	for {
		course, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(courses)
		}
		if err != nil {
			return err
		}
		courseResult, err := c.CourseDB.Create(course.Name, course.Description, course.CategoryId)
		if err != nil {
			return err
		}

		courses.Courses = append(courses.Courses, &pb.Course{
			Id:          courseResult.ID,
			Name:        courseResult.Name,
			Description: courseResult.Description,
			CategoryId:  course.CategoryId,
		})
	}
}

func (c *CourseService) CreateCourseStreamBidirecional(stream pb.CourseService_CreateCourseStreamBidirecionalServer) error {
	for {
		course, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		courseResult, err := c.CourseDB.Create(course.Name, course.Description, course.CategoryId)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.Course{
			Id:          courseResult.ID,
			Name:        courseResult.Name,
			Description: courseResult.Description,
			CategoryId:  courseResult.CategoryID,
		})
		if err != nil {
			return err
		}
	}
}
