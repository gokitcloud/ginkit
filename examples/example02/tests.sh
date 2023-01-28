curl -H "X-Token: 12345678" localhost:3333/test | jq
curl -H "X-Token: 12345678" localhost:3333/test?_fmt=yaml | yq
curl -H "X-Token: 12345678" localhost:3333/test?_fmt=toml
