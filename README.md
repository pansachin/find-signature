# find-signature
Forge github event payload signature

## Run server

```
git clone git@github.com:pansachin/find-signature.git
cd find-signature
go run main.go
```

## How to test
- Copy curl request payload from `payload.txt` or make your own.
- Run it in terminal.
- See `Actual Signature` and `Expected Signature` in server log.
- Actual Signature
  + Signature coming from payload.
- Expected Signature
  + Signature we need to successfully test againts `webhook-listener`.

## Example logs
```
2022/09/02 17:03:12 Actual signature: sha1=89945d34ad0073f7fb92d28cf154da642a324a0e
2022/09/02 17:03:12 Expected signature: sha1=89945d34ad0073f7fb92d28cf154da642a324a0e
2022/09/02 17:03:12 👏🎉Congratulations!!!!Signature matched
```

Test
