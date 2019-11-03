# punch a clock to WorkCloud

## Usage
Construct a new WorkCloud client with login info, then use the client to punch in/out at curent time. For example:
```go
client := workcloud.New("loginUser", "loginPass")
client.PunchTime("in")
```
