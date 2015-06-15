"""
Tests for post/models.py

"""
from app.tests.post.lib import PostTestBase


class PostModelTestCase(PostTestBase):

    def test_image_url(self):
        expected_url = (
            'http://images.mattbutterfield.com/post_images/' +
            self.post.image_uri)
        self.assertEqual(
            expected_url, self.post.image_url)
