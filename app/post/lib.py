"""
Helpers for Posts

"""
from flask.ext.restful import abort, fields

from app.post import api as post_api


class MethodField(fields.Raw):
    """
    A field which accepts a callable object to extract and format data from
    an object in a custom way.  The callable should expect one argument, the
    object being serialized.

    """

    def __init__(self, callable, **kwargs):
        self.callable = callable
        super(MethodField, self).__init__(**kwargs)

    def output(self, key, obj):
        return self.callable(obj)


def get_post_or_404(post_id):
    post = post_api.get(post_id)
    if not post:
        abort(404, message="Post {} not found".format(post_id))
    return post



