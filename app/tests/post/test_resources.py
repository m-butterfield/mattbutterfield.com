"""
Tests for post/resources.py

"""
import json

from flask.ext.restful import marshal

from app.post.lib import POST_FIELDS
from app.tests.post.lib import PostTestBase


class PostResourceTestCase(PostTestBase):

    def test_get(self):
        resp = self.client.get('/api/post/' + self.post_id)
        self.assertEqual(
            marshal(self.post, POST_FIELDS), json.loads(resp.data))

    def test_get_404(self):
        resp = self.client.get('/api/post/bogus_id')
        self.assertEqual(resp.status_code, 404)
