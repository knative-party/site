# www

Party on the web!

## Sources of Truth

For now we dup the data, but we would like to move this into the repos at some point.

- [Serving on-call](https://github.com/knative/serving/blob/main/support/COMMUNITY_CONTACTS.md)
- [Eventing on-call](https://github.com/knative/eventing/blob/main/support/COMMUNITY_CONTACTS.md)
- [Security on-call](https://github.com/knative/)
- [ToC meeting](https://docs.google.com/document/d/1LzOUbTMkMEsCRfwjYm5TKZUWfyXpO589-r9K2rXlHfk/edit#heading=h.jlesqjgc1ij3)
  - [WG Member List](https://github.com/knative/community/blob/main/working-groups/WORKING-GROUPS.md)

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
