"""
Helpers for post models/resources

"""
from urlparse import urlunsplit

from flask_restful import fields

from app import app


class ImageUrlField(fields.Raw):
    """
    Marshal field to build image urls based on the post id

    """

    def output(self, key, obj):
        return _build_image_url(obj.id)


def _build_image_url(image_id):
    return urlunsplit((
        app.config['PREFERRED_URL_SCHEME'],
        app.config['S3_IMAGE_BUCKET'],
        "{}/{}.jpg".format(app.config['S3_IMAGE_FOLDER'], image_id),
        None,
        None))
