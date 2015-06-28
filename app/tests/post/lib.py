"""
Helpers for Post tests

"""
from datetime import datetime

from app.post import api as post_api
from app.tests.lib import BaseTest


class PostTestBase(BaseTest):

    @classmethod
    def setUpClass(cls):
        super(PostTestBase, cls).setUpClass()
        cls.post_id = "post_id"
        cls.post, _ = post_api.get_or_create(
            cls.post_id, "image_uri", datetime.now(), 640, 640, "stuff")
