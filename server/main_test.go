package main

import (
	"context"

	"testing"

	pb "fulfillment/fulfillment"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRegisterDeliveryAgent(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err, "Error creating mock db: %v", err)

	defer db.Close()

	dialect := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})

	gormDb, err := gorm.Open(dialect, &gorm.Config{})
	assert.Nil(t, err, "Error creating mock gorm db: %v", err)

	type args struct {
		ctx context.Context
		req *pb.RegisterRequest
	}
	tests := []struct {
		name    string
		args    args
		rows    func()
		want    *pb.RegisterResponse
		wantErr bool
	}{
		{
			name: "Register a user",
			args: args{
				ctx: context.Background(),
				req: &pb.RegisterRequest{
					Username: "newUser",
					Password: "password123",
					City:     "TestCity",
				},
			},
			rows: func() {
				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{})) 
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{
					"id",
					"username",
					"password",
					"city",
					"availability",
				}).AddRow(1, "newUser", "password123", "TestCity", "AVAILABLE")
				mock.ExpectQuery("INSERT").WillReturnRows(rows)
				mock.ExpectCommit()
			},
			want: &pb.RegisterResponse{
				Message: "Successfully Registered!",
			},
			wantErr: false,
		},
		
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rows()
			server := &Server{DB: gormDb} 

			got, err := server.RegisterDeliveryAgent(tt.args.ctx, tt.args.req)

			if (err != nil) != tt.wantErr {
				t.Fatalf("RegisterDeliveryAgent() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				statusErr, ok := status.FromError(err)
				assert.True(t, ok, "Expected gRPC status error")
				assert.Equal(t, codes.AlreadyExists, statusErr.Code(), "Expected AlreadyExists error")
			} else {
				assert.Equalf(t, tt.want, got, "RegisterDeliveryAgent(%v, %v)", tt.args.ctx, tt.args.req)
			}
		})
	}
}
