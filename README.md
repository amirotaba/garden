Garden Trees,
before run make a MySQL database and enter the config in internal/utils/utils.go and then run cmd/main.go
 ```
 ├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── domain
│   │   ├── comment
│   │   │   └── comment.go
│   │   ├── domain.go
│   │   ├── garden
│   │   │   └── garden.go
│   │   ├── gardenLocation
│   │   │   └── gardenLocation.go
│   │   ├── gardenType
│   │   │   └── gardenType.go
│   │   ├── service
│   │   │   └── service.go
│   │   ├── tag
│   │   │   └── tag.go
│   │   ├── tree
│   │   │   └── tree.go
│   │   ├── treeType
│   │   │   └── treeType.go
│   │   ├── user
│   │   │   └── user.go
│   │   └── userType
│   │       └── userType.go
│   ├── features
│   │   ├── comment
│   │   │   ├── handler
│   │   │   │   └── http
│   │   │   │       └── handler.go
│   │   │   ├── repository
│   │   │   │   └── mysql
│   │   │   │       └── repository.go
│   │   │   └── usecase
│   │   │       └── usecase.go
│   │   ├── garden
│   │   │   ├── handler
│   │   │   │   └── http
│   │   │   │       └── handler.go
│   │   │   ├── repository
│   │   │   │   └── mysql
│   │   │   │       └── repository.go
│   │   │   └── usecase
│   │   │       └── usecase.go
│   │   ├── gardenLocation
│   │   │   ├── handler
│   │   │   │   └── http
│   │   │   │       └── handler.go
│   │   │   ├── repository
│   │   │   │   └── mysql
│   │   │   │       └── repository.go
│   │   │   └── usecase
│   │   │       └── usecase.go
│   │   ├── gardenType
│   │   │   ├── handler
│   │   │   │   └── http
│   │   │   │       └── handler.go
│   │   │   ├── repository
│   │   │   │   └── mysql
│   │   │   │       └── repository.go
│   │   │   └── usecase
│   │   │       └── usecase.go
│   │   ├── service
│   │   │   ├── handler
│   │   │   │   └── http
│   │   │   │       └── handler.go
│   │   │   ├── repository
│   │   │   │   └── mysql
│   │   │   │       └── repository.go
│   │   │   └── usecase
│   │   │       └── usecase.go
│   │   ├── tag
│   │   │   ├── handler
│   │   │   │   └── http
│   │   │   │       └── handler.go
│   │   │   ├── repository
│   │   │   │   └── mysql
│   │   │   │       └── repository.go
│   │   │   └── usecase
│   │   │       └── usecase.go
│   │   ├── tree
│   │   │   ├── handler
│   │   │   │   └── http
│   │   │   │       └── handler.go
│   │   │   ├── repository
│   │   │   │   └── mysql
│   │   │   │       └── repository.go
│   │   │   └── usecase
│   │   │       └── usecase.go
│   │   ├── treeType
│   │   │   ├── handler
│   │   │   │   └── http
│   │   │   │       └── handler.go
│   │   │   ├── repository
│   │   │   │   └── mysql
│   │   │   │       └── repository.go
│   │   │   └── usecase
│   │   │       └── usecase.go
│   │   ├── user
│   │   │   ├── handler
│   │   │   │   └── http
│   │   │   │       └── handler.go
│   │   │   ├── repository
│   │   │   │   └── mysql
│   │   │   │       └── repository.go
│   │   │   └── usecase
│   │   │       └── usecase.go
│   │   └── userType
│   │       ├── handler
│   │       │   └── http
│   │       │       └── handler.go
│   │       ├── repository
│   │       │   └── mysql
│   │       │       └── repository.go
│   │       └── usecase
│   │           └── usecase.go
│   ├── middleware
│   │   ├── access
│   │   │   └── access.go
│   │   └── jwt
│   │       └── jwt.go
│   └── utils
│       └── utils.go
└── README.md
```
