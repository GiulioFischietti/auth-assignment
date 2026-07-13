# API Reference

## Authentication Service

Base URL

```text
http://localhost:8080
```

---

### Register User

Creates a new user account.

| Property       | Value        |
| -------------- | ------------ |
| Method         | `POST`       |
| Endpoint       | `/register`  |
| Authentication | Not required |

#### Request

```json
{
    "username": "john",
    "password": "password123"
}
```

#### cURL Example

```bash
curl --location 'http://localhost:8080/register' \
--header 'Content-Type: application/json' \
--data '{
    "username": "john",
    "password": "password123"
}'
```

#### Successful Response

**201 Created**

```json
{
    "message": "user registered successfully"
}
```

#### Possible Errors

| Status                      | Description                                                                                |
| --------------------------- | ------------------------------------------------------------------------------------------ |
| `400 Bad Request`           | Invalid request payload.                                                                   |
| `500 Internal Server Error` | Registration failed (e.g. username already exists or an unexpected server error occurred). |

---

### Login

Authenticates an existing user and creates a new session.

| Property       | Value        |
| -------------- | ------------ |
| Method         | `POST`       |
| Endpoint       | `/login`     |
| Authentication | Not required |

#### Request

```json
{
    "username": "john",
    "password": "password123"
}
```

#### cURL Example

```bash
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "john",
    "password": "password123"
}'
```

#### Successful Response

**200 OK**

The generated session token is returned in the `Authorization` response header.

Example:

```http
Authorization: <session-token>
```

#### Possible Errors

| Status                      | Description                   |
| --------------------------- | ----------------------------- |
| `400 Bad Request`           | Invalid request payload.      |
| `401 Unauthorized`          | Invalid username or password. |
| `500 Internal Server Error` | Unexpected server error.      |

---

### Generate Access Token

Generates a short-lived JWT access token for a specific protected service.

| Property       | Value         |
| -------------- | ------------- |
| Method         | `POST`        |
| Endpoint       | `/token`      |
| Authentication | Session Token |

#### Request Header

```http
Authorization: <session-token>
```

#### Request Body

```json
{
    "service": "orders-service"
}
```

#### cURL Example

```bash
curl --location 'http://localhost:8080/token' \
--header 'Authorization: <session-token>' \
--header 'Content-Type: application/json' \
--data '{
    "service": "orders-service"
}'
```

#### Successful Response

**200 OK**

The generated JWT access token is returned in the `Authorization` response header.

Example:

```http
Authorization: <jwt-access-token>
```

#### Possible Errors

| Status                      | Description                                      |
| --------------------------- | ------------------------------------------------ |
| `400 Bad Request`           | Invalid request payload.                         |
| `401 Unauthorized`          | Invalid or expired session.                      |
| `403 Forbidden`             | Requested service is not registered or inactive. |
| `500 Internal Server Error` | Unexpected server error.                         |

---

# Protected Service

Base URL

```text
http://localhost:8081
```

---

### Get Orders

Returns all orders belonging to the authenticated user.

| Property       | Value            |
| -------------- | ---------------- |
| Method         | `GET`            |
| Endpoint       | `/orders`        |
| Authentication | JWT Access Token |

#### Request Header

```http
Authorization: <jwt-access-token>
```

#### cURL Example

```bash
curl --location 'http://localhost:8081/orders' \
--header 'Authorization: <jwt-access-token>'
```

#### Successful Response

**200 OK**

```json
[
    {
        "id": "...",
        "customer": {
            "id": 1,
            "username": "john"
        },
        "paymentStatus": "PAID",
        "orderStatus": "SHIPPED",
        "items": [
            {
                "name": "Mechanical Keyboard",
                "quantity": 1,
                "price": 129.99
            }
        ]
    }
]
```

#### Possible Errors

| Status                      | Description                                       |
| --------------------------- | ------------------------------------------------- |
| `401 Unauthorized`          | Missing, invalid or expired JWT access token.     |
| `403 Forbidden`             | Token audience does not match the target service. |
| `500 Internal Server Error` | Unexpected server error.                          |
