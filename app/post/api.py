"""
Api for using Post objects

"""
from app import db
from app.post.models import Post


def get(post_id):
    """
    Get an existing post by id

    Args:
        post_id (str): The post's id

    Returns:
        Post: The post object or None

    """
    return db.session.query(Post).get(post_id)


def create(post_id, image_id, created_at, text=None):
    """
    Create a new Post

    Args:
        id (str): Id of the new post
        image_id (str): Image id for this post
        created_at (Datetime): The date this post was created

    Kwargs:
        text (str): The text of the post

    Returns:
        Post: The newly created Post object

    """
    post = Post(id=post_id,
                image_id=image_id,
                created_at=created_at,
                text=text)
    db.session.add(post)
    db.session.commit()
    return post
