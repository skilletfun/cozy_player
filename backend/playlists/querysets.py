from django.db import models

from .models import Playlist


def base_queryset() -> models.QuerySet:
    return Playlist.objects.all()
