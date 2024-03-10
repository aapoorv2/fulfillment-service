package main

import (
	"context"
	"encoding/base64"
	"testing"

	"fulfillment/fulfillment"
	pb "fulfillment/fulfillment"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
		errorCode codes.Code
	}{
		{
			name: "Register a user - Expect Success",
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
		{
			name: "User already exists - Expect Error",
			args: args{
				ctx: context.Background(),
				req: &pb.RegisterRequest{
					Username: "existingUser",
					Password: "password123",
					City:     "TestCity",
				},
			},
			rows: func() {
				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			want:    nil,
			wantErr: true,
			errorCode: codes.AlreadyExists,
		},
		{
			name: "Missing username - Expect Error",
			args: args{
				ctx: context.Background(),
				req: &pb.RegisterRequest{
					// Missing username
					Password: "password123",
					City:     "TestCity",
				},
			},
			rows:    func() {},
			want:    nil,
			wantErr: true,
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Missing password - Expect Error",
			args: args{
				ctx: context.Background(),
				req: &pb.RegisterRequest{
					Username: "newUser",
					// Missing password
					City:     "TestCity",
				},
			},
			rows:    func() {},
			want:    nil,
			wantErr: true,
			errorCode: codes.InvalidArgument,
		},
		{
			name: "Missing city - Expect Error",
			args: args{
				ctx: context.Background(),
				req: &pb.RegisterRequest{
					Username: "newUser",
					Password: "password123",
					// Missing city
				},
			},
			rows:    func() {},
			want:    nil,
			wantErr: true,
			errorCode: codes.InvalidArgument,
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
				assert.Equalf(t, tt.errorCode, statusErr.Code(), "Expected %v error", tt.errorCode)
			} else {
				assert.Equalf(t, tt.want, got, "RegisterDeliveryAgent(%v, %v)", tt.args.ctx, tt.args.req)
			}
		})
	}
}

func TestAssignDeliveryAgent(t *testing.T) {
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
		req *pb.AssignRequest
	}
	tests := []struct {
		name    string
		args    args
		rows    func()
		want    *pb.AssignResponse
		wantErr bool
		errorCode codes.Code
	}{
		{
			name: "Assigning an agent to an order - Expect Success",
			args: args{
				ctx: context.Background(),
				req: &pb.AssignRequest{
					OrderId: 1,
					City: "TestCity",
				},
			},
			rows: func() {
				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "city", "availability"}).
					AddRow(1, "availableUser", "password123", "TestCity", "AVAILABLE"))

				mock.ExpectBegin()
				mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()


				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{
					"id",
					"order_id",
					"city",
					"delivery_agent_id",
				}).AddRow(1, 1, "TestCity", 1)
				mock.ExpectQuery("INSERT").WillReturnRows(rows)
				mock.ExpectCommit()
				
			},
			want: &pb.AssignResponse{
				Message: "Successfully Assigned a delivery agent!",
			},
			wantErr: false,
		},
		{
			name : "No delivery agents available - Expect Error",
			args: args{
				ctx: context.Background(),
				req: &pb.AssignRequest{
					OrderId: 1,
					City: "TestCity",
				},
			},
			rows: func() {
				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{}))
			},
			want: nil,
			wantErr: true,
			errorCode: codes.Unavailable,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rows()
			server := &Server{DB: gormDb} 

			got, err := server.AssignDeliveryAgent(tt.args.ctx, tt.args.req)

			if (err != nil) != tt.wantErr {
				t.Fatalf("AssignDeliveryAgent() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				statusErr, ok := status.FromError(err)
				assert.True(t, ok, "Expected gRPC status error")
				assert.Equalf(t, tt.errorCode, statusErr.Code(), "Expected %v error", tt.errorCode)
			} else {
				assert.Equalf(t, tt.want, got, "AssignDeliveryAgent(%v, %v)", tt.args.ctx, tt.args.req)
			}
		})
	}
}

func TestFetchingAllDeliveriesForADeliveryAgent(t *testing.T) {
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
		req *pb.FetchDeliveriesRequest
	}
	tests := []struct {
		name    string
		args    args
		rows    func()
		want    *pb.FetchDeliveriesResponse
		wantErr bool
		errorCode codes.Code
		errorMessage string
	}{
		{
			name: "Fetching all deliveries for an agent - Expect Success",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.MD{
					"authorization": []string{"Basic " + base64.StdEncoding.EncodeToString([]byte("testuser:testpassword"))},
				}),
				req: &pb.FetchDeliveriesRequest{
				},
			},
			rows: func() {
				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "city", "availabilty"}).
					AddRow(1, "testuser", "testpassword", "TestCity", "AVAILABLE"))

				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "order_id", "city", "delivery_agent_id", "status"}).
					AddRow(1, 1, "TestCity", 1, "DELIVERED").
					AddRow(2, 2, "TestCity", 1, "DELIVERED"))
			},
			want: &pb.FetchDeliveriesResponse{
				Deliveries: []*fulfillment.Delivery{
					{
						Id:     1,
						OrderId: 1,
						City:    "TestCity",
						Status:  "DELIVERED",
					},
					{
						Id: 2,
						OrderId: 2,
						City: "TestCity",
						Status: "DELIVERED",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Fetching deliveries but invalid username or password - Expect Unauthorized",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.MD{
					"authorization": []string{"Basic " + base64.StdEncoding.EncodeToString([]byte("testuser:testpassword"))},
				}),
				req: &pb.FetchDeliveriesRequest{
				},
			},
			rows: func() {
				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{})) 
			},
			want: nil,
			wantErr: true,
			errorCode: codes.Unauthenticated,
			errorMessage: "invalid credentials",
		},
		{
			name: "Fetching deliveries but authorization header not provided - Expect Unauthorized",
			args: args{
				ctx: context.Background(),
				req: &pb.FetchDeliveriesRequest{
				},
			},
			rows: func() {
			},
			want: nil,
			wantErr: true,
			errorCode: codes.Unauthenticated,
			errorMessage: "authorization header not found",
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rows()
			server := &Server{DB: gormDb} 

			got, err := server.FetchAllDeliveriesForAnAgent(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FetchAllDeliveriesForAnAgent() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				statusErr, ok := status.FromError(err)
				assert.True(t, ok, "Expected gRPC status error")
				assert.Equalf(t, tt.errorCode, statusErr.Code(), "Expected %v error", tt.errorCode)
				assert.Equal(t, tt.errorMessage, statusErr.Message())
			} else {
				assert.Equalf(t, tt.want, got, "FetchAllDeliveriesForAnAgent(%v, %v)", tt.args.ctx, tt.args.req)
			}
		})
	}
}

