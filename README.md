# Splitwise
Application for managing splits with friends or family when we want to split the amount which we spend

#### Run the Application 
```bash
  go run main.go 
```

```http
GET /api/account/balance
```
Endpoints with Authentication token and gives current balance of the account 

```http
GET /api/account/spent
```

Gives the totals spends by the Account 
```http
POST /api/account/
```

endpoint that creates account 

```bash
      {
        "username":"",
        "password":"",
        "confirmPassword":"",
        "email":""
      }
```

###  