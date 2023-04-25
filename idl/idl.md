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

> **TODO**: add authority in v2 API. The plan of RBAC authority system, is an independent module. For each request in API-gateway, it would ask the authority module primary. If there is a success content of response, then gateway will send it to the actual module like user or device.

#### Some Introduction

- Default Page Limit is Page 1, Limit 15.

- In Delete Operation, it is actually lazy delete. If a user is deleted, the data will write the timestamp into the field `delete_at`. In each query, there will be a condition of `delete_at is null`. 
- In Create Operation, if someone is renewed in the database who was deleted from database. User Module Service will check there are same phone number, it will remove the content of field `delete_at` and become the variable user data.

---
> The following temp Content is a template generate by GPT.

### Device Module

#### API List

##### Create Device

Create a new device.

```
rpc CreateDevice(CreateDeviceReq) returns (CreateDeviceRsp) {}
```

> Request

| Field             | Type   | Description             |
|-------------------|--------|-------------------------|
| `device_name`     | string | Name of the device.     |
| `device_type`     | string | Type of the device.     |
| `device_location` | string | Location of the device. |

> Response

| Field       | Type   | Description                     |
|-------------|--------|---------------------------------|
| `device_id` | string | ID of the newly created device. |
| `status`    | string | Status message.                 |

---

###### Remove Device

Remove an existing device.

```
rpc RemoveDevice(RemoveDeviceReq) returns (RemoveDeviceRsp) {}
```

> Request

| Field       | Type   | Description                     |
|-------------|--------|---------------------------------|
| `device_id` | string | ID of the device to be removed. |

> Response

| Field       | Type   | Description               |
|-------------|--------|---------------------------|
| `device_id` | string | ID of the removed device. |
| `status`    | string | Status message.           |

---

##### Update Device

Update an existing device.

```
rpc UpdateDevice(UpdateDeviceReq) returns (UpdateDeviceRsp) {}
```

> Request

| Field             | Type   | Description                     |
|-------------------|--------|---------------------------------|
| `device_id`       | string | ID of the device to be updated. |
| `device_name`     | string | Updated name of the device.     |
| `device_type`     | string | Updated type of the device.     |
| `device_location` | string | Updated location of the device. |

> Response

| Field       | Type   | Description               |
|-------------|--------|---------------------------|
| `device_id` | string | ID of the updated device. |
| `status`    | string | Status message.           |

---

##### Get Device

Get details of a device.

```
rpc GetDevice(GetDeviceReq) returns (GetDeviceRsp) {}
```

> Request

| Field       | Type   | Description                         |
|-------------|--------|-------------------------------------|
| `device_id` | string | ID of the device to get details of. |

> Response

| Field             | Type   | Description             |
|-------------------|--------|-------------------------|
| `device_id`       | string | ID of the device.       |
| `device_name`     | string | Name of the device.     |
| `device_type`     | string | Type of the device.     |
| `device_location` | string | Location of the device. |

---

##### List Devices

Get a list of all devices.

```
rpc ListDevice(UpdateDeviceReq) returns (ListDeviceRsp) {}
```

> Response

| Field     | Type            | Description          |
|-----------|-----------------|----------------------|
| `devices` | list of objects | List of all devices. |
| `status`  | string          | Status message.      |

> Devices List Object

| Field             | Type   | Description             |
|-------------------|--------|-------------------------|
| `device_id`       | string | ID of the device.       |
| `device_name`     | string | Name of the device.     |
| `device_type`     | string | Type of the device.     |
| `device_location` | string | Location of the device. |

---

That's all! Thank you for using OpeniotDeviceService.w