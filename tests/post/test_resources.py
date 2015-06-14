"""
Tests for post/resources.py

"""
import json

from app.post import api as post_api

from tests.post.lib import PostTestBase


class PostResourceTestCase(PostTestBase):

    def test_get(self):
        resp = json.loads(self.client.get('/api/post/' + self.post_id).data)
        self.assertEqual(post_api.get(self.post_id).id, resp['id'])
