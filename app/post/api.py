"""
Api for using Post objects

"""
from sqlalchemy.exc import IntegrityError

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


def get_or_create(post_id, image_uri, created_at, text=None):
    """
    Create a new Post

    Args:
        id (str): Id of the new post
        image_uri (str): Image uri for this post
        created_at (Datetime): The date this post was created

    Kwargs:
        text (str): The text of the post

    Returns:
        tuple(Post, bool): The newly created Post object, whether it was
            created or not

    """
    post = Post(id=post_id,
                image_uri=image_uri,
                created_at=created_at,
                text=text)

    db.session.add(post)

    try:
        db.session.commit()
    except IntegrityError:
        db.session.rollback()
        return get(post_id), False

    return post, True
