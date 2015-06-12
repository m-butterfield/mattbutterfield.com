from flask import Flask
from flask.ext.sqlalchemy import SQLAlchemy
from flask_restful import Api

app = Flask(__name__)
app.config.from_object('config')

api = Api(app)

db = SQLAlchemy(app)


from app.views import index


from app.post.resources import PostResource


api.add_resource(PostResource, '/api/post/<post_id>')
