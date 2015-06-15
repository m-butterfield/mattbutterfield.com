"""
Tests for post/lib.py

"""
from datetime import datetime

from app.post import api as post_api
from app.post.lib import ImageUrlField
from app.tests.lib import BaseTest


class ImageUrlFieldTestCase(BaseTest):

    def test_output(self):
        image_id = "image_id"
        post = post_api.create("123", image_id, datetime.now())
        image_url = ImageUrlField().output(None, post)
        self.assertEqual(
            'http://images.mattbutterfield.com/post_images/123.jpg', image_url)
