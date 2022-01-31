# goatest
just a test of goa as a framework

1. Write design/design.go
2. Run ```goa gen github.com/xxx/yyy/design```
3. Run ```goa example github.com/xxx/yyy/design```
4. Now you have a boilerplate for a service. This service has one or more files on the top-level github.com/xxx/yyy that need to implement the business logic you want to achieve. 
5. The you compile ./cmd/yyyserver/ and (if you want) the cli ./cmd/yyyserver-cli 
6. start testing!
