```bash
├── Dockerfile
├── README.md
├── assets
│   └── customer_support.png
└── src
    ├── Makefile
    ├── cmd
    │   └── main.go
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── domain
    │   │   ├── acl.go
    │   │   ├── company.go
    │   │   ├── company_user.go
    │   │   ├── role.go
    │   │   ├── service.go
    │   │   └── user.go
    │   ├── infrastructure
    │   │   └── server
    │   │       └── route.go
    │   └── value_object
    │       ├── faq.go
    │       ├── faq_attachment.go
    │       ├── ticket.go
    │       ├── ticket_attachment.go
    │       └── ticket_reply.go
    └── web
        └── handler
            ├── hello_world.go
            └── user_handler.go
```
