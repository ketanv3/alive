# alive <img src="./docs/heart.svg" width="80" align="right" />

[![Maintainability](https://api.codeclimate.com/v1/badges/48384621113d4c380813/maintainability)](https://codeclimate.com/github/ketanv3/alive/maintainability)

Hyper-scalable health-checker written completely in Go
+ Fully configuration driven
+ Define custom health-checkers
+ Supports sync and async health-checks
+ Extremely lightweight and performant

## Usage Instructions

### Health Check Definition

Health-checks are defined using `.yaml` configuration files. Let's create a simple health-check that checks if google.com is down or not:

```yaml
name: google-com
strategy: async
interval: 30.0
checker:
  type: url
  parameters:
    url: https://google.com
    timeout: 2.0
retryPolicy:
  initialDelay: 1.0
  backoffMultiplier: 2.0
  maxRetries: 3
```

Each health-check needs a unique `name` property and the type of health `checker` that'll be used.

Health-checks can be `sync` or `async`, defined by the strategy field. The periodicity of async health-checks is defined by the `interval` field.

### Starting the Server

Start the health-check server by passing the listen-port and the list of health-check definitions:
```bash
./alive -p 8055 google-com.yaml
```

Health-check results can then be accessed at:
```
http://localhost:8055/health
http://localhost:8055/health?full=true
```

Health-check strategy is defined in the definition file, but can be overridden with an extra query parameter:
```
http://localhost:8055/health?full=true&strategy=[default|sync|async]
```

## Issues and Suggestions
If you encounter any issues or have suggestions, please [file an issue](https://github.com/ketanv3/alive/issues) along with a detailed description. Remember to apply labels for easier tracking.


## Versioning
We use [SemVer](http://semver.org/) for versioning. For the available versions, see the [tags on this repository](https://github.com/ketanv3/alive/tags)


## Authors
See the list of [contributors](https://github.com/ketanv3/alive/contributors) who participated in this project.