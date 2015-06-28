"""
Shared helper code for Posts

"""
from flask.ext.restful import abort, fields, marshal

from app.post import api as post_api


POST_FIELDS = {
    'id': fields.String,
    'text': fields.String,
    'created_at': fields.DateTime,
    'image_url': fields.String,
    'image_width': fields.Integer,
    'image_height': fields.Integer,
    'next_post_id': fields.String,
    'previous_post_id': fields.String,
}


def serialize_post(post):
    """
    Serialize a Post object

    Args:
        post (Post): The Post object to serialize

    """
    return marshal(post, POST_FIELDS)


def get_post_or_404(post_id):
    """
    Attempt to fetch a Post from the database. Abort with a 404 if not found.

    Args:
        post_id (str): The id of the Post

    """
    post = post_api.get(post_id)
    if not post:
        abort(404, message="Post {} not found".format(post_id))
    return post
