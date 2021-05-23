# yaml-cli

Command line tool to generate commands from YAML file and run on terminal

## TL;DR

Save the command file as `ls.yaml`

```yaml
type: Command
spec:
  cmds:
  - cmd: ls
    opts:
    - opt: "l"
```

Run this command file.

```sh
yaml-cli ls.yaml

-rw-r--r--  1 hiroyukosaki  staff  11358  5 22 16:47 LICENSE
-rw-r--r--  1 hiroyukosaki  staff    192  5 22 20:35 README.md
drwxr-x--x  4 hiroyukosaki  staff    128  5 22 20:40 cmd
...
```
