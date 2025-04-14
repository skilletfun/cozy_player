import os
import music_tag
from django.db import models

from artists.models import Artist
from common.utils import is_image_file


class Track(models.Model):
    title = models.CharField()
    artist = models.ForeignKey(Artist, on_delete=models.CASCADE, related_name="tracks")
    path = models.CharField()

    duration = models.IntegerField()
    play_count = models.IntegerField(default=0)

    def _cover_in_folder(self) -> str | None:
        path = self.path[: self.path.rfind("/")]
        for file in os.listdir(path):
            if is_image_file(file_path := os.path.join(path, file)):
                return file_path

    @property
    def cover(self):
        if (cover := music_tag.load_file(self.path)["artwork"].first) is not None:
            return cover.raw

        path = self._cover_in_folder() or self.artist.cover
        return open(path, "rb").read()
