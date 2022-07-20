-Click on the Raw Format View
-Garden Trees
-before run make a MySQL database and enter the config in cmd/main.go and then run main.go

├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── domain
│   │   └── user.go
│   └── user
│       ├── delivery
│       │   └── httpdelivery
│       │       └── user_handler.go
│       ├── repository
│       │   └── mysqlhandler
│       │       └── mysqlhandler.go
│       └── usecase
│           └── user_usecase.go
└── README.md


User funcs
├──── SignUp
├─── SignIn
├── Account(account info)
└─ Comment(add comment to trees)

Farmer funcs
├────── SignIn
├───── ShowTrees
├──── ShowComments
├─── AddTree
├── RemoveTree
└─ AddAttend

Admin funcs
├───── SignIn
├──── ShowGarden
├─── AddGarden
├── RemoveGarden
└─ Addfarmer
