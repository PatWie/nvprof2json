# nvprof2json

[![Build Status](https://ci.patwie.com/api/badges/PatWie/nvprof2json/status.svg)](https://ci.patwie.com/PatWie/nvprof2json)

A small self-contained binary that converts the output of nvprof into a json file for visualization in the Chrome browser.
No external dependencies, just download the [latest release](https://github.com/PatWie/nvprof2json/releases) here from GitHub and run

```bash
nvprof -o main.nvvp ./main
nvprof2json -o main.json main.nvvp
```

Point the Chrome browsers to `chrome://tracing/` and load `main.json`

<p align="center"> <img src="./screenshot.gif" /> </p>


```
./nvprof2json --help
Usage:
  nvprof2json [OPTIONS] nvvpfile

Application Options:
  -o, --output=   output file for Chrome (default: [nvvpfile].json)
  -f              override output file if exists

Help Options:
  -h, --help      Show this help message

Arguments:
  nvvpfile:       output from nvprof, e.g., 'nvprof -o main.nvvp ./main'

```

**Alternatives**
There exists [proof-of-concept](https://github.com/ezyang/nvprof2json) written in Python with less features.
