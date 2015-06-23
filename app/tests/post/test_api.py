"""
Tests for post/api.py

"""
from datetime import datetime, timedelta

from app.post import api as post_api
from app.tests.post.lib import PostTestBase


class PostAPITestCase(PostTestBase):

    def test_get(self):
        self.assertEqual(post_api.get(self.post_id), self.post)

    def test_create(self):
        new_post_id = "new_post"
        post, _ = post_api.get_or_create(
            new_post_id, "new_image_uri", datetime.now(), 640, 640)
        self.assertEqual(new_post_id, post.id)

    def test_get_most_recent_post(self):
        recent_id = "recent_id"
        recent_post, _ = post_api.get_or_create(
            recent_id,
            "recent_image_uri",
            self.post.created_at + timedelta(days=1),
            640,
            640,)
        result = post_api.get_most_recent_post()
        self.assertEqual(recent_post, result)
