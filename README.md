# text-env

## Install

```
go install github.com/ToshihitoKon/text-env
```

## Usage

template.tmpl

```
Dear {{must_env "NAME"}},

It was a pleasure to see you at the wedding.
Thank you for the lovely {{env "GIFT"}}.

Best wishes,
Josie
```

```
$ NAME="Aunt Mildred" GIFT="bone china tea set" text-env template.tmpl

Dear Aunt Mildred,

It was a pleasure to see you at the wedding.
Thank you for the lovely bone china tea set.

Best wishes,
Josie


$ go run ./... template.tmpl

Dear execution: template: template:1:7: executing "template" at <must_env "NAME">: error calling must_env: error: NAME is must_env, but this env is empty
```
