package main

import (
	"fmt"
	"log"
	"github.com/gocql/gocql"
)

func main() {
	// Khởi tạo cluster
	cluster := gocql.NewCluster("127.0.0.1") // Địa chỉ node Scylla
	cluster.Keyspace = "system"               // Keyspace mặc định để test
	cluster.Consistency = gocql.One

	// Tạo session
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Không thể tạo session: %v", err)
	}
	defer session.Close()

	// Truy vấn đơn giản
	var clusterName string
	if err := session.Query("SELECT cluster_name FROM system.local").Scan(&clusterName); err != nil {
		log.Fatalf("Không thể truy vấn ScyllaDB: %v", err)
	}

	fmt.Printf("Tên cụm Scylla: %s\n", clusterName)
}
