"""
Tests for post/api.py

"""
from datetime import datetime

from app.post import api as post_api
from app.tests.post.lib import PostTestBase


class PostAPITestCase(PostTestBase):

    def test_get(self):
        self.assertEqual(post_api.get(self.post_id), self.post)

    def test_create(self):
        post_id2 = "post_id2"
        post, _ = post_api.get_or_create(post_id2, "image_uri", datetime.now())
        self.assertEqual(post_id2, post.id)
