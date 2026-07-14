# Table of Contents

* [1. Authentication Service](#1-authentication-service)

  * [1.1 Health Check](#11-health-check)
  * [1.2 Register User](#12-register-user)
  * [1.3 Login](#13-login)
  * [1.4 Generate Access Token](#14-generate-access-token)
  * [1.5 Log Out](#15-log-out)

* [2. Protected Service](#2-protected-service)

  * [2.1 Health Check](#21-health-check)
  * [2.1 Get Orders](#21-get-orders)


# API Reference

## [1. Authentication Service](#1-authentication-service)

Base URL

```text
http://localhost:8080
```

---

## [1.1 Health Check](#11-health-check)

Returns the current status of the service.

This endpoint does not require authentication and is intended for monitoring purposes, container health checks and service availability verification.

| Property       | Value        |
| -------------- | ------------ |
| Method         | `GET`        |
| Endpoint       | `/health`    |
| Authentication | Not required |

### cURL Example

```bash
curl --location 'http://localhost:8080/health'
```

### Successful Response

**200 OK**

```json
{
    "status": "ok"
}
```

### Possible Errors

| Status                      | Description             |
| --------------------------- | ----------------------- |
| `500 Internal Server Error` | Service is not healthy. |

```
```


### [1.2 Register User](#12-register-user)

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
    "message": "user created"
}
```

#### Possible Errors

| Status                      | Description                                                                                |
| --------------------------- | ------------------------------------------------------------------------------------------ |
| `400 Bad Request`           | Invalid request payload.                                                                   |
| `500 Internal Server Error` | Registration failed (e.g. username already exists or an unexpected server error occurred). |

---

### [1.3 Login](#13-login)

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

The generated session token is returned in the response body.

Example:

```json
{
    "session_token": "<session-token>"
}
```

#### Possible Errors

| Status                      | Description                   |
| --------------------------- | ----------------------------- |
| `400 Bad Request`           | Invalid request payload.      |
| `401 Unauthorized`          | Invalid username or password. |
| `500 Internal Server Error` | Unexpected server error.      |

---

### [1.4 Generate Access Token](#14-generate-access-token)

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
    "service_name": "orders-service"
}
```

#### cURL Example

```bash
curl --location 'http://localhost:8080/token' \
--header 'Authorization: <session-token>' \
--header 'Content-Type: application/json' \
--data '{
    "service_name": "orders-service"
}'
```

#### Successful Response

**200 OK**

The generated JWT access token is returned in the response body as `access-token`.

Example response:

```json
{
    "access_token": "<jwt-access-token>"
}
```

#### Possible Errors

| Status                      | Description                                      |
| --------------------------- | ------------------------------------------------ |
| `400 Bad Request`           | Invalid request payload.                         |
| `401 Unauthorized`          | Invalid or expired session.                      |
| `403 Forbidden`             | Requested service is not registered or inactive. |
| `500 Internal Server Error` | Unexpected server error.                         |

---
### [1.5 Log Out](15-log-out)

Invalidates the current authenticated session.

The client must provide the session token in the `Authorization` header. If the session exists and is still valid, it is revoked and can no longer be used to generate new JWT access tokens.

Already issued JWT access tokens remain valid until their expiration time.

| Property       | Value                                  |
| -------------- | -------------------------------------- |
| Method         | `POST`                                 |
| Endpoint       | `/logout`                              |
| Authentication | Session Token (`Authorization` header) |

### Request Headers

| Header          | Value             |
| --------------- | ----------------- |
| `Authorization` | `<session_token>` |

### cURL Example

```bash
curl --location --request POST 'http://localhost:8080/logout' \
--header 'Authorization: <session_token>'
```

### Successful Response

**200 OK**

```json
{
    "message": "Logout completed successfully."
}
```

### Possible Errors

| Status                      | Description                                |
| --------------------------- | ------------------------------------------ |
| `400 Bad Request`           | Missing or malformed request.              |
| `401 Unauthorized`          | Invalid, expired or revoked session token. |
| `500 Internal Server Error` | Unexpected server error.                   |


# [2. Protected Service](#2-protected-service)

Base URL

```text
http://localhost:8081
```


## [2.1 Health Check](#21-health-check)

Returns the current status of the service.

This endpoint does not require authentication and is intended for monitoring purposes, container health checks and service availability verification.

| Property       | Value        |
| -------------- | ------------ |
| Method         | `GET`        |
| Endpoint       | `/health`    |
| Authentication | Not required |

### cURL Example

```bash
curl --location 'http://localhost:8081/health'
```

### Successful Response

**200 OK**

```json
{
    "status": "ok"
}
```

### Possible Errors

| Status                      | Description             |
| --------------------------- | ----------------------- |
| `500 Internal Server Error` | Service is not healthy. |

```
```


---

### [2.1 Get Orders](#21-get-orders)

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
