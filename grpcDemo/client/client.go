package main

import (
	"io"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "demo/grpcDemo/customer"
)

const (
	address  = "localhost:9000"
	parallel = 2  //连接并行度
	times    = 10 //每连接请求次数
)

var (
	wg sync.WaitGroup
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	currTime := time.Now()

	for i := 0; i < int(parallel); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			connexe()
		}()
	}
	wg.Wait()

	log.Printf("time taken: %.2f ", time.Now().Sub(currTime).Seconds())
}

func connexe() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerClient(conn)

	for i := 0; i < int(times); i++ {
		customer := &pb.CustomerRequest{
			Id:    int32(i),
			Name:  "golang-" + strconv.Itoa(i),
			Email: strconv.Itoa(i) + "@wxy.com",
			Phone: "0592-632-2924",
			Addresses: []*pb.CustomerRequest_Address{
				&pb.CustomerRequest_Address{
					Street:            strconv.Itoa(i) + "Street",
					City:              "shanghai",
					State:             "CA",
					Zip:               "94105",
					IsShippingAddress: true,
				},
			},
		}

		createCustomer(client, customer)
	}

	filter := &pb.CustomerFilter{Keyword: ""}
	getCustomers(client, filter)
}

func createCustomer(client pb.CustomerClient, customer *pb.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d", resp.Id)
	}
}

func getCustomers(client pb.CustomerClient, filter *pb.CustomerFilter) {

	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v", err)
	}
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}
		log.Printf("Customer: %v", customer.Name+" : "+customer.Addresses[0].Street)
	}
}
