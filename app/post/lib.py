"""
Helpers for post models/resources

"""
from flask_restful import fields


IMAGE_URL_BASE = 'http://images.mattbutterfield.com/post_images/{}.jpg'


class ImageUrlField(fields.Raw):
    """
    Marshal field to build image urls based on the post id

    """

    def output(self, key, obj):
        return IMAGE_URL_BASE.format(obj.id)
