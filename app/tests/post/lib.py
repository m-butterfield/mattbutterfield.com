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
        cls.post = post_api.create(
            cls.post_id, "image_id", datetime.now(), "stuff")
