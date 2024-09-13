package main

import (
    "encoding/json"
    "fmt"
)

type Product struct {
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Quantity int     `json:"quantity"`
}
// Convert a Product instance to JSON
func ProductToJSON(p Product) (string, error) {
    jsonData, err := json.Marshal(p)     //кодировка
    if err != nil {
        return "", err
    }
    return string(jsonData), nil
}
// Convert a JSON string to a Product instance
func JSONToProduct(jsonStr string) (Product, error) {
    var p Product
    err := json.Unmarshal([]byte(jsonStr), &p)  //декодировка
    if err != nil {
        return Product{}, err
    }
    return p, nil
}
func main() {
    // Create a Product instance
    product := Product{
        Name:     "Laptop",
        Price:    1200.50,
        Quantity: 10,
    }

    // Convert Product to JSON
    jsonStr, err := ProductToJSON(product)
    if err != nil {
        fmt.Println("Error encoding to JSON:", err)
    } else {
        fmt.Println("Product as JSON:", jsonStr)
    }

    // Convert JSON back to Product
    decodedProduct, err := JSONToProduct(jsonStr)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
    } else {
        fmt.Printf("Decoded Product: %+v\n", decodedProduct)
    }
}

