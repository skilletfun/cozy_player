from django.db import models

from artists.models import Artist
from tracks.models import Track


class Playlist(models.Model):
    name = models.CharField()
    artists = models.ManyToManyField(Artist, related_name="playlists")
    tracks = models.ManyToManyField(Track, related_name="playlists")
