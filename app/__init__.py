from flask import Flask
from flask.ext.sqlalchemy import SQLAlchemy
from flask.ext.restful import Api

app = Flask(__name__)
app.config.from_object('config')

api = Api(app)

db = SQLAlchemy(app)

from app.post.resources import PostResource
from app.views import index

api.add_resource(PostResource, '/api/post/<post_id>')
