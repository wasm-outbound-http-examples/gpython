# Use Go-Python to send HTTP(s) requests from inside WASM

## Instructions for this devcontainer

Tested with Go 1.21.4, Bun 1.0.13, Deno 1.38.2, gpython [e9cde5fc](https://github.com/go-python/gpython/commit/e9cde5fcf8e89407c50c12534c1b14bb1f840ffa).

### Preparation

1. Open this repo in devcontainer, e.g. using Github Codespaces.
   Type or copy/paste following commands to devcontainer's terminal.

### Building

1. `cd` into the folder of this example:

```sh
cd browser-and-deno
```

Since current (early Nov 2023) version of gpython has some WASM incompatibilities yet,
need to use source distribution of gpython and apply a patch to it.

2. Clone gpython repo:
```sh
git clone --depth=1 https://github.com/go-python/gpython
```

3. `cd` into the folder of gpython sources and apply a provided patch:

```sh
cd gpython
git apply ../gpython.patch
```

4. Make directory for WASM example, `cd` into it and copy provided code there:

```sh
mkdir main
cd main
cp ../../main.go ./
cp ../../*.js ./
cp ../../index.html ./
```

5. Compile the example:

```sh
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

6. Copy the glue JS from Golang distribution to example's folder:

```sh
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

### Test with browser

1. Run simple HTTP server to temporarily publish project to Web:

```sh
python3 -m http.server
```

Codespace will show you "Open in Browser" button. Just click that button or
obtain web address from "Forwarded Ports" tab.

2. As `index.html` and a **14.5M**-sized `main.wasm` are loaded into browser, refer to browser developer console
   to see the results.

### Test with Node.js

Impossible yet due to https://github.com/golang/go/issues/59605.

### Test with Bun

1. Install Bun:

```sh
curl -fsSL https://bun.sh/install | bash
```

2. Run with Bun:

```sh
~/.bun/bin/bun bun.js
```

### Test with Deno

1. Install Deno:

```sh
curl -fsSL https://deno.land/x/install/install.sh | sh
```

2. Run with Deno:

```sh
~/.deno/bin/deno run --allow-read --allow-net deno.js
```

### Finish

Perform your own experiments if desired.
