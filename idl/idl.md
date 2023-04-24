# OpenIoT-server Interface Definition Language File Document

```
.
├── api                     
│   ├── api.proto           // copyright of protobuf.descriptor.proto 
│   │                       // for the api annotations
│   ├── gateway.proto       // contents rpc services need to expose to gateway
│   ├── message             // contains req and rsp with api annotaion.
│   │   ├── auth.proto      // different file means different modules.
│   │   └── node.proto
│   └── model.proto         // contains all basic module messages
│                           // like User, Device and so on.
├── base.proto              // some public proto messages like BaseRsp
├── device.proto            // device module service and req and rsp messages.
└── user.proto              // user module service and req, rsp messages.
```

Obviously, there are two versions of idl files, which are taken the api annotations in the `api/` dir, while others in the `idl/` root dir with no annotations are prepare for generate the basic stub for different microservices.

Api Online Document: `https://example.com/api/`

## API List

The following is the list of api defined in the idl files.

### User Module

#### API List

- `POST    /api/v1/user`

  Create a new user. 

- `DELETE  /api/v1/user/:id`

  Delete the user specified by id. 

- `POST    /api/v1/user/:id`

  Update user Info specified by id. 

- `GET     /api/v1/user`

  Get User List. Default `PageLimit` is 15.  , 

- `GET     /api/v1/user/:id`

> **TODO**: add authority in v2 API. The plan of RBAC authority system, is a independent module. For each request in API-gateway, it would ask the authority module primary. If there is a success content of response, then gateway will send it to the actually module like user or device.

#### Some Introduction

- Default Page Limit is Page 1, Limit 15.

- In Delete Operation, it is actually lazy delete. If a user is deleted, the data will write the timestamp into the field `delete_at`. In each query, there will be a condition of `delete_at is null`. 
- In Create Operation, if someone is renew in the database who was deleted from database. User Module Service will check there are same phone number, it will remove the content of field `delete_at` and become the variable user data.