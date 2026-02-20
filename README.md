# Project Title
Online Test Eratani Backend Engineer Akhmad Ali Ernawan

## 🚀 Installation
### Prerequisites
- Go (>= 1.24 recommended)
- PostgreSQL (for point 3)
- golang-migrate (for point 3)
- Make

### Install Dependencies (Point 3)
```bash
cd 3
make bin-deps
```
This command installs required binary dependencies (e.g., golang-migrate).

## ▶ Running the Project
### Run Point 1, 2, or 4
```bash
cd 1   # or 2 or 4
go run .
```

### Run Point 3 (With Database)
1. Setup Environment Variables (Configure your PostgreSQL environment variables)
2. Run the Application
```bash
cd 3
make migrate-up
make run
```

## 🧪 Testing the Project
### Testing Point 1
After running:
```bash
go run .
```
Provide input based on this test case:
```bash
10
1
2
3
4
5
6
7
8
9
1000
```
Expected output:
```bash
1
2
4
5
7
8
10
11
14
1666
```

### Testing Point 2 or 4
Modify the test case inside the source code variable "testInput"

Then run:
```bash
go run .
```

### Testing Point 3 (REST API)
1. Get All Data User
```bash
curl -X GET http://localhost:8080/v1/user
```

2. Get Most Spend Country Data
```bash
curl -X GET http://localhost:8080/v1/user/a
```

3. Get Most Credit Card Type Data
```bash
curl -X GET http://localhost:8080/v1/user/b
```

4. Create User
```bash
curl -X POST http://localhost:8080/v1/user \
-H "Content-Type: application/json" \
-d '{
  "data": {
    "country": "Indonesia",
    "credit_card_type": "mastercard",
    "credit_card": "5002356896930801",
    "first_name": "Akhmad Ali",
    "last_name": "Ernawan"
  }
}'
```

## 🏗 Architecture (Point 3)
Point 3 architecture is inspired by the Clean Architecture template from
Evrone — https://github.com/evrone/go-clean-template
