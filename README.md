My personal website.

[mattbutterfield.com](http://mattbutterfield.com)

# Development

## Local

The main dependencies are: Postgres (version 13 to match CloudSQL), Go, Docker, and Node. Install as usual with homebrew, etc...

To run locally on OSX, you'll need some extras installed:

    brew install lame
    brew install pkg-config
    brew install vips

Plus some environment variables set:

    export CGO_LDFLAGS="-L/opt/homebrew/opt/lame/lib" 
    export CGO_CFLAGS="-I/opt/homebrew/opt/lame/include"

Then try running:

    make test

### Getting a prod db dump:

[cloud-sql-proxy](https://cloud.google.com/sql/docs/mysql/sql-proxy) must be installed first. Then run:

    sudo mkdir /var/cloudsql && sudo chown $USER /var/cloudsql
    cloud-sql-proxy --unix-socket /var/cloudsql <CLOUD_SQL_INSTANCE_CONNECTION_NAME>

From another terminal, `pg_dump` can now be run (this will require the db password):

    pg_dump -h /var/cloudsql/mattbutterfield:us-central1:mattbutterfield -U mattbutterfield mattbutterfield > dump.sql

Now the local database can be set up and the data loaded:

    make reset-db

### Running the server

Init some env variables in case we want to use the pubsub emulator:

    $(gcloud beta emulators pubsub env-init)

Now the server can be started and accessed at [http://localhost:8000/](http://localhost:8000/):

    make run-server

The pubsub emulator can optionally be started for testing video chat locally:

    gcloud beta emulators pubsub start --project=mattbutterfield

### Building the Docker Images

To build the docker images:

    make docker-build

To run the server image locally:

    docker run -e DB_SOCKET=host=host.docker.internal\ dbname=mattbutterfield\ user=matthewbutterfield \
           -e PUBSUB_EMULATOR_HOST=host.docker.internal:8085 \
           -e AUTH_TOKEN=1234 \
           -dp 8000:8000 gcr.io/mattbutterfield/mattbutterfield.com

## Deployment

The website is deployed on GCP using Terraform (see the `infra` directory).

Running `make deploy` will build the docker images, push them to GCR and deploy them as Cloud Run services.

Connecting to the Cloud SQL instance will get you into a `psql` shell where you can edit the schema and data as needed:

    gcloud beta sql connect mattbutterfield --user=mattbutterfield --database=mattbutterfield --quiet

