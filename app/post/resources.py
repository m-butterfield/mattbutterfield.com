"""
Resources for Posts

"""
from flask.ext.restful import abort, fields, marshal_with, Resource

from app import db
from app.post.models import Post


POST_FIELDS = {
    'id': fields.String,
    'text': fields.String,
    'created_at': fields.DateTime,
    'image_url': fields.String,
    'next_post_id': fields.String(attribute='next_post.id'),
    'previous_post_id': fields.String(attribute='previous_post.id'),
}


class PostResource(Resource):

    @marshal_with(POST_FIELDS)
    def get(self, post_id):
        post = db.session.query(Post).get(post_id)
        if not post:
            abort(404, message="Post {} not found".format(post_id))
        return post
