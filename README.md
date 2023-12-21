# jubelio-chat
Jubelio-Chat is Realtime chat application, try to clone the project.

```
git clone https://github.com/rizalilhamm/jubelio-chat.git
```
## Needed Stack
1. Golang
2. Go Fiber
3. Websocket
4. PostgreSQL
5. Supabase Realtime Account

### Installation
```
go mod init
```
### Download Needed Modules
```
go mod download
```

### PostgreSQL Database
Create your database named jubelio_chat, and migrate the models by running
```
go run migrations/migrate.go
```

## Start server
```
go run main.go
```