# ods-pipeline-tkn

[![Tests](https://github.com/opendevstack/ods-pipeline-tkn/actions/workflows/main.yaml/badge.svg)](https://github.com/opendevstack/ods-pipeline-tkn/actions/workflows/main.yaml)

Tekton task for use with [ODS Pipeline](https://github.com/opendevstack/ods-pipeline) to run [`tkn`](https://github.com/tektoncd/cli) commands.

## Usage

```yaml
tasks:
- name: logs
  taskRef:
    resolver: git
    params:
    - { name: url, value: https://github.com/opendevstack/ods-pipeline-tkn.git }
    - { name: revision, value: v0.1.0 }
    - { name: pathInRepo, value: tasks/logs.yaml }
    workspaces:
    - { name: source, workspace: shared-workspace }
```

See the [documentation](https://github.com/opendevstack/ods-pipeline-tkn/blob/main/docs/logs.adoc) for detailed usage and available task parameters.

## About this repository

`docs` and `tasks` are generated directories from recipes located in `build`. See the `Makefile` target for how everything fits together.
