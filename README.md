# AUCA Numerical computation software template

The template might be used to create web applications for performing and graphing various computations. 

### Dev commands

`make run` - runs the backend in the dev mode (go run)

`make front` - runs the frontend in the dev mode (yarn dev)

`make docker` - creates the docker image (with both front and back)

`make release` - builds the front and back (for different architectures and OSs), creates distributive, commits, tags and pushes the new version, creates a GitHub release with uploaded distributive, creates docker and pushes it to the registry if possible. 

`make major` - creates a release with a new major version (1.x.x -> 2.0.0)

`make minor` - creates a release with a new major version (x.1.x -> x.2.0)

`make patch` - creates a release with a new major version (x.x.1 -> x.x.2)

`make clean` - deletes build artifacts and distros

**Note**: `make release` and other commands, which create releases require following software:

- **Golang 1.18+** - to build the binaries
- **Node.js 16+** - to build the frontend
- **Yarn** - js package manager
- **Docker** - to create containers (pushing to registry is optional)
- **GitHub CLI** - to create releases and upload ditributive (optional)
