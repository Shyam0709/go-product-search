🧠 Product Search Backend API (Go)
This project is a simple Go-based backend API that simulates 100,000 product records in-memory and allows full-text search using the Bleve search engine. It uses the Chi router for handling HTTP requests and supports graceful shutdown.

🚀 Features
Simulates 100,000+ product records with random names and categories

Full-text search using Bleve

REST API endpoint for search

Graceful server shutdown using Go's context and signals

Fast response time for search queries

Easily extensible

🧪 API Usage
Search Endpoint
sql
Copy
Edit
GET /search?q=<query>
Query parameter: q — the keyword to search for in product names

Example:

bash
Copy
Edit
http://localhost:8080/search?q=Product
Response:

json
Copy
Edit
[
  {
    "id": 1,
    "name": "Product 1",
    "category": "Electronics"
  },
  ...
]
🛠️ Tech Stack
Technology	Purpose
Go	Backend language
Chi	HTTP router
Bleve	Full-text indexing and search
Postman	API testing

📦 Installation & Run Locally
Prerequisite: Go 1.21+ installed

Clone the repo

bash
Copy
Edit
git clone https://github.com/Shyam0709/go-product-search.git
cd go-product-search
Initialize Go modules

bash
Copy
Edit
go mod tidy
Run the server

bash
Copy
Edit
go run main.go
Test the API
Use Postman or a browser to access:

bash
Copy
Edit
http://localhost:8080/search?q=Product
💡 Notes
Currently simulates 100,000 products for performance reasons. This can be increased as needed.

Data is stored in-memory and not persisted.

Example categories include: Electronics, Footwear, Home & Kitchen, Fitness, Books, Clothing

📄 Decision Log https://docs.google.com/document/d/147Qe_ht3WqQK_xj79qbZWZZ2Mr3CNY8s0KcCHLcdks0/edit?usp=sharing

See Google Doc Decision Log for detailed documentation on decisions, learning path, and challenges.

👨‍💻 Author
Shyam Sundar
GitHub: @Shyam0709
