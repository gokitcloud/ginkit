curl -H "X-Token: 12345678" localhost:3333/test | jq
curl -H "X-Token: 12345678" localhost:3333/test/123/a/b/c/d | jq
curl -H "X-Token: 12345678" localhost:3333/org/123 | jq
curl -H "X-Token: 12345678" localhost:3333/org/123/asdf | jq
curl -H "X-Token: 12345678" localhost:3333/org/123?_fmt=json | jq
curl -H "X-Token: 12345678" localhost:3333/org/123?_fmt=yaml | yq
curl -H "X-Token: 12345678" localhost:3333/org/123?_fmt=toml
curl -H "X-Token: 12345678" localhost:3333/org/1234?_fmt=json | jq
curl -H "X-Token: 123456789" localhost:3333/org/123?_fmt=json | jq
