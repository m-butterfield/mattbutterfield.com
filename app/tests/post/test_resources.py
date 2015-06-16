"""
Tests for post/resources.py

"""
import json

from flask.ext.restful import marshal

from app.post.resources import POST_FIELDS, POST_PAGINATION_FIELDS
from app.tests.post.lib import PostTestBase


class PostResourceTestCase(PostTestBase):

    def test_get(self):
        resp = self.client.get('/api/post/' + self.post_id)
        data = marshal(
            self.post, POST_PAGINATION_FIELDS, envelope='pagination')
        data.update(marshal(self.post, POST_FIELDS, envelope='data'))
        self.assertEqual(data, json.loads(resp.data))

    def test_get_404(self):
        resp = self.client.get('/api/post/bogus_id')
        self.assertEqual(resp.status_code, 404)
