# opje
Golang Service Locator Thingy

## Usage

```go
package main

type AuthService interface {
    Login(username, password string) (bool, error)
    Register(username, password string) (bool, error)
}

type authentication struct {}

func newAuthService() AuthService {
    return &authentication{}
}

func (a *authentication) Login(username, password string) (bool, error) {
    return true, nil
}

func (a *authentication) Register(username, password string) (bool, error) {
    return true, nil
}

func init(){
    locator.Register(newAuthService())
}

func main(){
    auth, ok := locator.Resolve.[AuthService]()

    if !ok {
        panic("authService not registered")
    }

    if(auth.Login("admin", "admin")){ //😎
        fmt.Println("Logged in")
    }
}
```