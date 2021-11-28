My personal website.

[mattbutterfield.com](http://mattbutterfield.com)

# Development

## Local

The main dependencies are: Postgres, Go, Docker, and Node. The website is deployed on GCP.

To get started, run:

    make test

### Getting a prod db dump:

[cloud_sql_proxy](https://cloud.google.com/sql/docs/mysql/sql-proxy) must be installed first. Then run:

    ./cloud_sql_proxy -dir=/var/cloudsql

From another terminal, `pg_dump` can now be run (this will require the db password), followed by `psql` to load the data locally:

    pg_dump -h /var/cloudsql/mattbutterfield:us-central1:mattbutterfield -U mattbutterfield mattbutterfield > dump.sql
    psql -f dump.sql mattbutterfield

### Running the server

Init some env variables in case we want to use the pubsub emulator:

    $(gcloud beta emulators pubsub env-init)

Now the server can be started and accessed at [http://localhost:8000/](http://localhost:8000/):

    make run

The pubsub emulator can optionally be started for testing video chat locally:

    gcloud beta emulators pubsub start --project=mattbutterfield

### Building the Docker Container

To build the docker container:

    docker build -t gcr.io/mattbutterfield/mattbutterfield.com .

To run the container locally:

    docker run -e DB_SOCKET=host=host.docker.internal\ dbname=mattbutterfield\ user=matthewbutterfield \
           -e PUBSUB_EMULATOR_HOST=host.docker.internal:8085 \
           -dp 8000:8000 gcr.io/mattbutterfield/mattbutterfield.com

## Deployment

The deploy script `bin/deploy` will build the docker container, push it to GCR and deploy it to Cloud Run.

Connecting to the Cloud SQL instance will get you into a `psql` shell where you can edit the schema and data as needed:

    gcloud sql connect mattbutterfield --user=mattbutterfield --database=mattbutterfield --quiet

