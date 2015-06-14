"""
Resources for Posts

"""
from flask.ext.restful import fields, marshal_with, Resource

from app import db
from app.post.models import Post
from app.post.lib import ImageUrlField


post_fields = {
    'id': fields.String,
    'text': fields.String,
    'image_url': ImageUrlField,
}


class PostResource(Resource):

    @marshal_with(post_fields)
    def get(self, post_id):
        return db.session.query(Post).get(post_id)
