# JSON SCHEMA GENERATOR

## Description

CLI utility for generating JSON Schema from json file

- you can select fixture right in CLI or pass relative path with argument `-fixture fixtures/any/path.json`;
- schema can be saved with other name just for comparison or some testing
- can be executed with VS Code Cypress extension explorer menu item `Cypress: Generate JSON Schema`

## How to use as binary

1. Go to [releases](https://github.com/Sheex/json-schema-gen/releases) section
2. Download `json-schema-gen_%Version%_%OS_Specific%`
3. Unzip to `cypress/scripts` and rename to `qa-json-schema-generator` as VS Code cypress extension command `generate json schema` relies on path
4. Execute `cypress/scripts/qa-json-schema-generator` from command line OR enable `Cypress: Generate JSON Schema` in case VS Code Cypress extension is used.

## How to build locally

Prerequisite: [Golang](https://golang.org/)

1.  `git clone https://github.com/Shelex/json-schema-gen.git`
2.  `cd json-schema-gen`
3.  `make build`
4.  `mv json-schema-gen %%Path_to_testing_repo%%/cypress/scripts/qa-json-schema-generator`
