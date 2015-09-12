My personal website, written in Python, using [Flask](http://flask.pocoo.org/), [SQLAlchemy](http://www.sqlalchemy.org/), [SQLite](https://www.sqlite.org/), and [Backbone](http://backbonejs.org/).

[mattbutterfield.com](http://mattbutterfield.com)

[![Circle CI](https://circleci.com/gh/m-butterfield/mattbutterfield.com.png?circle-token=c615ced31f0190dbb0405f67aa1ccb44b8f3c9cd)](https://circleci.com/gh/m-butterfield/mattbutterfield.com)

## Setup

### Install dependencies:
Besides Python (version 2.7), you must have working installations of [Node.js](https://nodejs.org/) and [SQLite](https://www.sqlite.org/).

    $ pip install -r requirements.txt 


### Set environment variables

API keys are referenced from environment variables to avoid storing them in source code.  Run the commands below, pasting in your tokens and keys where necessary (removing the `<>` characters as well):

    $ export INSTAGRAM_CLIENT_SECRET=<paste instagram client secret here>
    $ export INSTAGRAM_ACCESS_TOKEN=<paste instagram access token here>
    $ export AWS_ACCESS_KEY_ID=<paste aws access key here>
    $ export AWS_SECRET_ACCESS_KEY=<paste aws secret access key here>
