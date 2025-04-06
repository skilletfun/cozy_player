from django.conf import settings


def is_image_file(filename: str) -> bool:
    return filename.split(".")[-1] in ("jpg", "jpeg", "png")


def is_track_file(filename: str) -> bool:
    return filename.split(".")[-1] in settings.TRACKS_EXTENSIONS
