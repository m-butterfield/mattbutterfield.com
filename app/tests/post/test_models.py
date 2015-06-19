"""
Tests for post/models.py

"""
from datetime import timedelta

from app.post import api as post_api
from app.tests.post.lib import PostTestBase


class PostModelTestCase(PostTestBase):

    def test_image_url(self):
        expected_url = (
            'http://images.mattbutterfield.com/post_images/' +
            self.post.image_uri)
        self.assertEqual(
            expected_url, self.post.image_url)

    def test_next_post(self):
        next_post, _ = post_api.get_or_create(
            'post_id2', 'image_uri2', self.post.created_at + timedelta(days=1))
        self.assertEqual(self.post.next_post_id, next_post.id)
        self.assertIsNone(next_post.next_post_id)

    def test_previous_post(self):
        previous_post, _ = post_api.get_or_create(
            'post_id3', 'image_uri3', self.post.created_at - timedelta(days=1))
        self.assertEqual(self.post.previous_post_id, previous_post.id)
        self.assertIsNone(previous_post.previous_post_id)
