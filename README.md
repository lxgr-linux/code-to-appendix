# code-to-appendix
Utility to create a latex appendix containing listings of code

## Usage
Inside of the folder containing the code
```shell
code-to-appendix
```

## Config
An optional config can be used

`config.yaml`
```yaml
outputFile: "/home/arch/Dokumente/Uni/TFL/TFL/frame/appendix.pb.tex"
excludeFiles:
  - .git/
  - .mvn/
  - target/
  - .fleet
  - .idea
  - code-to-appendix.yaml
excludeFormats:
  - md
```

To use just add the config file as an commanline argument
```shell
code-to-appendix config.yaml
```
