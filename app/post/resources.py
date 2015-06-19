"""
Resources for Posts

"""
from flask.ext.restful import fields, marshal, Resource

from app.post.lib import get_post_or_404, MethodField


def _next_post_url(post):
    next_post = post.next_post
    if next_post:
        return '/post/{}'.format(next_post.id)


def _previous_post_url(post):
    previous_post = post.previous_post
    if previous_post:
        return '/post/{}'.format(post.previous_post.id)


POST_FIELDS = {
    'id': fields.String,
    'text': fields.String,
    'created_at': fields.DateTime,
    'image_url': fields.String,
    'next_post_url': MethodField(_next_post_url),
    'previous_post_url': MethodField(_previous_post_url),
}


def serialize_post(post):
    """
    Serialize a Post object

    """
    return marshal(post, POST_FIELDS)


class PostResource(Resource):
    """
    Api Resource for posts

    """
    def get(self, post_id):
        return serialize_post(get_post_or_404(post_id))
