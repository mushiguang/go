package v1

import (
	"context"

	"github.com/rebirthmonkey/go/pkg/log"
)

// ListStudents lists the students in the storage.
func (c *controller) ListStudents(ctx context.Context, r *ListStudentsRequest) (*ListStudentsResponse, error) {
	log.L(ctx).Info("[GrpcServer] controller: ListStudents")

	students, err := c.srv.NewStudentService().List()
	if err != nil {
		return nil, err
	}

	items := make([]*StudentInfo, 0)
	for _, student := range students.Items {
		items = append(items, &StudentInfo{
			Nickname:  student.Name,
			Password:  student.Password,
			Email:     student.Email,
			Phone:     student.Phone,
			LoginedAt: student.LoginedAt.Format("2006-01-02 15:04:05"),
		})

	}

	return &ListStudentsResponse{
		TotalCount: students.TotalCount,
		Items:      items,
	}, nil
}