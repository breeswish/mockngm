# mockngm

A just-for-fun project that pretend to be a [ng-monitoring](https://github.com/pingcap/ng-monitoring). It will not record anything at all, and doesn't need a PD instance to work. 

## Usage

```shell
make
bin/mockngm
```

To specify different targets:

```shell
bin/mockngm --targets tidb://127.0.0.1:10080,tikv://127.0.0.1:20180
```
