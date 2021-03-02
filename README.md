# www
Party on the web!


## Release

```
pushd partyui; yarn build; popd
ko apply -f ./config
```

## Local dev

```
pushd partyui; yarn build; popd
KO_DATA_PATH=$PWD/kodata go run .
```