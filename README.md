# Websnake - a small multiplayer web game

## Requirements

- Go 1.8+, with [Glide](https://glide.sh/)
- Node 8.1.3+ (for building the UI)

## Building the Back-end

1. Check out the code at `$GOPATH/src/github.com/fsufitch/websnake`
1. `cd $GOPATH/src/github.com/fsufitch/websnake`
1. `glide intall`; this installs dependencies under `vendor/`
1. `go build`
1. Enjoy your static built executable

## Building the Front-end

1. Check out the code anywhere and enter the directory
1. `npm install`; this installs dependencies under `node_modules/`
1. `npm run webpack`
1. Enjoy the static built artifacts under `dist/`

Alternatively, for auto-building, use `npm run webpack:server`. Your perpetually up-to-date artifacts are still under `dist/` but also available at `http://localhost:8888`.

### Front-end build environment variables

The front-end build process is configured using environment variables. They have sane defaults, but feel free to modify them.

| Variable  | Default (if no value) | Description |
| ------------- | ------------- | ------ |
| `WEBPACK_DEV_SERVER_PORT`  | `8888` | the port at which `npm run webpack:server` listens |
| `ENV`  | `dev` | `prod` or `dev`; affects the level of minification and source mapping used |

## Running a game server

The game server can proxy the front-end resources (HTML, JS, CSS, etc), but it can only find them if they are already hosted on a webserver themselves. This is because in a "production" deployment, they are stored in a private AWS S3 bucket. Locally, the Webpack dev server can fulfill this purpose just fine. That means you must have this command running in the background:

    npm run webpack:server

Once that is running (or you have otherwise configured your server to point to valid front-end resources), run:

    go run main.go

Or, if you have manually built an executable, `./websnake` or `websnake.exe`, depending on your platform.

### Runtime server environment variables

The server process is configured using environment variables. For sane out-of-the-box defaults simply run `source scripts/dev_environment.sh`.

| Variable  | Dev Default | Description |
| ------------- | ------------- | ------- |
| `PORT`  | `8080` | the port at which the server process listens |
| `API_HOST`  | `localhost:8080` | the host to be injected into the UI resources (should be `<hostname>:<PORT>`, probably matching the one from above) |
| `UI_RES_URL` | `http://localhost:8888` | the base URL that the backend should proxy UI resources from |
| `PROXY_TTL` | `1` | time (in seconds) that the server should hold on to UI resources before re-querying |
| `DEBUG` | `true` | `true` or `false`; toggle for debug logging |
