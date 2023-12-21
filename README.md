# Blocked-list-IP

This project is a simple API that stores and retrieves blocked ip addresses.

## Usage

Is the ip address blocked?

```bash
curl -X GET http://localhost:8080/{ip}
```

example:

```bash
curl -X GET http://localhost:8080/8.8.8.8
```

Block an ip address

```bash
curl -X POST http://localhost:8080/{ip}
```

example:

```bash
curl -X POST http://localhost:8080/109.2.2.1
```

Unblock an ip address

```bash
curl -X DELETE http://localhost:8080/{ip}
```
