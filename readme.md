# Self-pickup Order with Microservices
- API Gateaway: Recieve request from the client in http
- Auth Service: Manage user authentication and authorization
- Product Service: Manage available products
- Order Service: Handler order from users

## Diagram
![Image](https://github.com/user-attachments/assets/30501277-b552-4e09-9b6b-d30bfa200854)

## To-do List
- [x] Create admin login and register
- [x] Create order service
- [x] setup Docker on every service

## Run Locally

Clone the project

```bash
  git clone https://github.com/magistraapta/e-comm-microservices
```

Go to the project directory

```bash
  cd e-comm-microservices
```

Run services

```bash
  make run-all
```

Run on Docker

```bash
  make run
```

