import os

from django.conf import settings
from django.db import models

from common.utils import is_image_file


class Artist(models.Model):
    name = models.CharField()

    def __str__(self) -> str:
        return str(self.name)

    @property
    def path(self) -> str:
        return os.path.join(settings.MUSIC_FOLDER, self.name)

    @property
    def cover(self) -> str:
        if not os.path.exists(self.path):
            return settings.DEFAULT_ARTIST_COVER

        for file in os.listdir(self.path):
            file_path = os.path.join(self.path, file)
            if is_image_file(file_path):
                return file_path

        return settings.DEFAULT_ARTIST_COVER
