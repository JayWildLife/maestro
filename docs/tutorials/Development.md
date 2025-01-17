## Development

-----

### Setting up the environment 

#### Grpc gateway
In order to run make generate with success, you need to have grpc-gateway dependencies installed with the following command:
```shell
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

#### Golang version
The project requires golang version 1.18 or higher.

#### Building and running
1. Run `make setup` to get all required modules
2. Run `make generate` to generate mocks, protos and wire (dependency injection)
3. Run `make deps/up` to startup service dependencies
4. Run `make migrate` to migrate database with the most updated schema

#### Running tests

1. Run `make run/unit-tests` to run all unit tests
2. Run `make run/integration-tests` to run all integration tests
3. Run `make run/e2e-tests` to run all E2E tests. NOTE: Currently it is not
   possible to run it with the development environment set. This command will
   stop the dev dependencies before running.
4. Run `make lint` to run all registered linters

---

### Running locally

To help you get along with Maestro, by the end of this section you should have a scheduler up and running.

#### Prerequisites
- Golang v1.18+
- Linux/MacOS environment
- Docker

#### Clone Repository
Clone the [repository](https://github.com/topfreegames/maestro) to your favorite folder.

#### Getting Maestro up and running
> For this step, you need docker running on your machine.

> **WARNING: Ensure using cgroupv1**
>
> K3s needs to use the deprecated `cgroupv1`, to successfully run the project in your machine ensure that your current docker use this version.


In the folder where the project was cloned, simply run:

```shell
make maestro/start
```

This will build and start all containers needed by Maestro, such as databases and maestro-modules. This will also start
all maestro components, including rooms api, management api, runtime watcher, and execution worker.

Because of that, be aware that it might take some time to finish.

#### Find rooms-api address
To simulate a game room, it's important to find the address of running **rooms-api** on the local network.

To do that, with Maestro containers running, simply use:

```shell
docker inspect -f '{{range.NetworkSettings.Networks}}{{.Gateway}}{{end}}' {{ROOMS_API_CONTAINER_NAME}}
```

This command should give you an IP address.
This IP is important because the game rooms will use it to communicate their status.

#### Create a scheduler
If everything is working as expected now, each Maestro-module is up and running.
Use the command below to create a new scheduler:

> Be aware to change the {{ROOMS_API_ADDRESS}} for the one found above.
```shell
curl --request POST \
  --url http://localhost:8080/schedulers \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "scheduler-run-local",
	"game": "game-test",
	"state": "creating",
	"portRange": {
		"start": 1,
		"end": 1000
	},
	"maxSurge": "10%",
	"spec": {
		"terminationGracePeriod": "100",
		"containers": [
			{
				"name": "alpine",
				"image": "alpine",
				"imagePullPolicy": "IfNotPresent",
				"command": [
					"sh",
					"-c",
					"apk add curl && while true; do curl --request PUT {{ROOMS_API_ADDRESS}}:8070/scheduler/$MAESTRO_SCHEDULER_NAME/rooms/$MAESTRO_ROOM_ID/ping --data-raw '\''{\"status\": \"ready\",\"timestamp\": \"12312312313\"}'\'' && sleep 5; done"
				],
				"environment": [],
				"requests": {
					"memory": "100Mi",
					"cpu": "100m"
				},
				"limits": {
					"memory": "200Mi",
					"cpu": "200m"
				},
				"ports": [
					{
						"name": "port-name",
						"protocol": "tcp",
						"port": 12345
					}
				]
			}
		],
		"toleration": "",
		"affinity": ""
	},
	"forwarders": []
}'
```

#### Congratulations
If you followed the steps above you have Maestro running in your local machine, and with a [scheduler](../reference/Scheduler.md) to try different [operations](../reference/Operations.md) on it.
Feel free to explore the available endpoints in the [API](../reference/OpenAPI.md) hitting directly the management-API.

If you have any doubts or feedbacks regarding this process, feel free to reach out in [Maestro's GitHub repository](https://github.com/topfreegames/maestro) and open an issue/question.

