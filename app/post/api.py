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
        db.session.flush()
    except IntegrityError:
        post = get(post_id), False

    _update_post_links(post)

    db.session.commit()

    return post, True


def _update_post_links(post):
    next_post = _get_next_post(post)
    previous_post = _get_previous_post(post)
    if next_post:
        post.next_post_id = next_post.id
        next_post.previous_post_id = post.id
    if previous_post:
        post.previous_post_id = previous_post.id
        previous_post.next_post_id = post.id


def _get_next_post(post):
    return (db.session.query(Post)
            .filter(Post.created_at > post.created_at)
            .order_by(Post.created_at).first())

def _get_previous_post(post):
    return (db.session.query(Post)
            .filter(Post.created_at < post.created_at)
            .order_by(Post.created_at.desc()).first())


def get_most_recent_post():
    """
    Get the most recently created post

    Returns:
        Post: The most recently created post object or None

    """
    return db.session.query(Post).order_by(Post.created_at.desc()).first()
