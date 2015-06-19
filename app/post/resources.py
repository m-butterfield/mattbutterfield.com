"""
Resources for Posts

"""
from flask.ext.restful import Resource

from app.post.lib import get_post_or_404, serialize_post


class PostResource(Resource):
    """
    Api Resource for posts

    """
    def get(self, post_id):
        return serialize_post(get_post_or_404(post_id))
