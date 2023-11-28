# Splitwise
Application for managing splits with friends or family when we want to split the amount which we spend

## Run the Application 
```bash
  go run main.go 
```

```http
  POST  /api/account/
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get item

```http
  GET /api/account/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |