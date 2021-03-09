# www

Party on the web!

## Release

```
pushd partyui; yarn build; popd
ko apply -f ./config
```

## Local dev

You'll need [`yarn`](https://yarnpkg.com/), and to run `yarn install` once in
the `partyui` directory.

```
pushd partyui; yarn build; popd
KO_DATA_PATH=$PWD/kodata go run .
```
