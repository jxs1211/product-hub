https://aws.amazon.com/blogs/containers/introducing-efs-csi-dynamic-provisioning/

#### Deploy order
- sc
- pvc
- deployment
- svc



##### error:
```sh
docker-compose git:(master) ✗ k logs -f memos-79c847885b-52s8c                      
failed to get profile, error: 1 error(s) decoding:

* cannot parse 'Port' as int: strconv.ParseInt: parsing "tcp://10.100.108.81:5230": invalid syntax
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x50 pc=0x138d6f2]

goroutine 1 [running]:
github.com/usememos/memos/store/db.NewDBDriver(0x192cff0?)
        /backend-build/store/db/db.go:18 +0x12
main.glob..func1(0xc0001b0400?, {0x16a7fa1?, 0x4?, 0x16a7fa5?})
        /backend-build/bin/memos/main.go:51 +0x57
github.com/spf13/cobra.(*Command).execute(0x23a7420, {0xc00012c060, 0x0, 0x0})
        /go/pkg/mod/github.com/spf13/cobra@v1.8.0/command.go:987 +0xaa3
github.com/spf13/cobra.(*Command).ExecuteC(0x23a7420)
        /go/pkg/mod/github.com/spf13/cobra@v1.8.0/command.go:1115 +0x3ff
github.com/spf13/cobra.(*Command).Execute(...)
        /go/pkg/mod/github.com/spf13/cobra@v1.8.0/command.go:1039
main.Execute()
        /backend-build/bin/memos/main.go:108 +0x3d
main.main()
```
