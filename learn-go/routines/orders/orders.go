package orders

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// A WaitGroup is a synchronization primitive in Go (from sync package) that lets
// you wait for a collection of goroutines to finish.

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func HandleOrders() {
	var wg sync.WaitGroup
	wg.Add(2)
	orders := generateOrders(20)

	go func() {
		defer wg.Done()
		processOrders(orders)
	}()

	go func() {
		defer wg.Done()
		updateOrderStatuses(orders)
	}()

	wg.Wait()
	reportOrderStatus(orders)
	fmt.Println("All operations completed. Exiting.")
	fmt.Print("total updates", totalUpdates)
}

func updateOrderStatuses(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		status := []string{
			"Processing", "Delivered", "Shipped",
		}[rand.Intn(3)]
		order.Status = status
		fmt.Printf("Updated order status for order ID: %d, status:%s\n", order.ID, order.Status)
		updateMutex.Lock()
		currentUpdates := totalUpdates
		time.Sleep(5 * time.Millisecond)
		totalUpdates = currentUpdates + 1
		updateMutex.Unlock()
	}
}

func processOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Processing order %d \n", order.ID)
	}
}

func reportOrderStatus(orders []*Order) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("\nOrder status report")
		for _, order := range orders {
			fmt.Printf("Order ID: %d %s \n", order.ID, order.Status)
		}
		fmt.Println("-----")
	}
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID: i + 1, Status: "pending",
		}
	}
	return orders
}
